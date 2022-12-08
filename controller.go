/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"k8s.io/client-go/rest"
	resourceclient "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
	"knative.dev/pkg/apis"
	"log"
	"time"

	samplev1alpha1 "github.com/ajklthr/proactive-pod-autoscalar/pkg/apis/autoscaling/v1alpha1"
	clientset "github.com/ajklthr/proactive-pod-autoscalar/pkg/client/clientset/versioned"
	samplescheme "github.com/ajklthr/proactive-pod-autoscalar/pkg/client/clientset/versioned/scheme"
	informers "github.com/ajklthr/proactive-pod-autoscalar/pkg/client/informers/externalversions/autoscaling/v1alpha1"
	listers "github.com/ajklthr/proactive-pod-autoscalar/pkg/client/listers/autoscaling/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

const controllerAgentName = "sample-controller"

const (
	// SuccessSynced is used as part of the Event 'reason' when a Foo is synced
	SuccessSynced = "Synced"
	// ErrResourceExists is used as part of the Event 'reason' when a Foo fails
	// to sync due to a Deployment of the same name already existing.
	ErrResourceExists = "ErrResourceExists"

	// MessageResourceExists is the message used for Events when a resource
	// fails to sync due to a Deployment already existing
	MessageResourceExists = "Resource %q already exists and is not managed by Foo"
	// MessageResourceSynced is the message used for an Event fired when a Foo
	// is synced successfully
	MessageResourceSynced = "Foo synced successfully"
)

// Controller is the controller implementation for PodAutoScalar resources
type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// sampleclientset is a clientset for our own API group
	sampleclientset clientset.Interface

	deploymentsLister appslisters.DeploymentLister
	deploymentsSynced cache.InformerSynced
	foosLister        listers.PodAutoscalerLister
	foosSynced        cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
}

// NewController returns a new sample controller
func NewController(
	kubeclientset kubernetes.Interface,
	sampleclientset clientset.Interface,
	deploymentInformer appsinformers.DeploymentInformer,
	fooInformer informers.PodAutoscalerInformer) *Controller {

	// Create event broadcaster
	// Add sample-controller types to the default Kubernetes Scheme so Events can be
	// logged for sample-controller types.
	utilruntime.Must(samplescheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:     kubeclientset,
		sampleclientset:   sampleclientset,
		deploymentsLister: deploymentInformer.Lister(),
		deploymentsSynced: deploymentInformer.Informer().HasSynced,
		foosLister:        fooInformer.Lister(),
		foosSynced:        fooInformer.Informer().HasSynced,
		workqueue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Foos"),
		recorder:          recorder,
	}

	klog.Info("Setting up event handlers")
	// Set up an event handler for when Foo resources change
	fooInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueFoo,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueFoo(new)
		},
	})
	// Set up an event handler for when Deployment resources change. This
	// handler will lookup the owner of the given Deployment, and if it is
	// owned by a Foo resource then the handler will enqueue that Foo resource for
	// processing. This way, we don't need to implement custom logic for
	// handling Deployment resources. More info on this pattern:
	// https://github.com/kubernetes/community/blob/8cafef897a22026d42f5e5bb3f104febe7e29830/contributors/devel/controllers.md
	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.handleObject,
		UpdateFunc: func(old, new interface{}) {
			newDepl := new.(*appsv1.Deployment)
			oldDepl := old.(*appsv1.Deployment)
			if newDepl.ResourceVersion == oldDepl.ResourceVersion {
				// Periodic resync will send update events for all known Deployments.
				// Two different versions of the same Deployment will always have different RVs.
				return
			}
			controller.handleObject(new)
		},
		DeleteFunc: controller.handleObject,
	})

	return controller
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) Run(workers int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	klog.Info("Starting Foo controller")

	// Wait for the caches to be synced before starting workers
	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.deploymentsSynced, c.foosSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.Info("Starting workers")
	// Launch two workers to process Foo resources
	for i := 0; i < workers; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	klog.Info("Started workers")
	<-stopCh
	klog.Info("Shutting down workers")

	return nil
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget if we
		// do not want this work item being re-queued. For example, we do
		// not call Forget if a transient error occurs, instead the item is
		// put back on the workqueue and attempted again after a back-off
		// period.
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		// We expect strings to come off the workqueue. These are of the
		// form namespace/name. We do this as the delayed nature of the
		// workqueue means the items in the informer cache may actually be
		// more up to date that when the item was initially put onto the
		// workqueue.
		if key, ok = obj.(string); !ok {
			// As the item in the workqueue is actually invalid, we call
			// Forget here else we'd go into a loop of attempting to
			// process a work item that is invalid.
			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// Run the syncHandler, passing it the namespace/name string of the
		// Foo resource to be synced.
		if err := c.syncHandler(key); err != nil {
			// Put the item back on the workqueue to handle any transient errors.
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		// Finally, if no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		c.workqueue.Forget(obj)
		klog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

// syncHandler compares the actual state with the desired, and attempts to
// converge the two. It then updates the Status block of the Foo resource
// with the current status of the resource.
func (c *Controller) syncHandler(key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	// Get the Foo resource with this namespace/name
	foo, err := c.foosLister.PodAutoscalers(namespace).Get(name)
	if err != nil {
		// The Foo resource may no longer exist, in which case we stop
		// processing.
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("foo '%s' in work queue no longer exists", key))
			return nil
		}

		return err
	}

	deploymentName := foo.Name
	if deploymentName == "" {
		// We choose to absorb the error here as the worker would requeue the
		// resource otherwise. Instead, the next time the resource is updated
		// the resource will be queued again.
		utilruntime.HandleError(fmt.Errorf("%s: deployment name must be specified", key))
		return nil
	}

	// Get the deployment with the name specified in Foo.spec
	deployment, err := c.deploymentsLister.Deployments(foo.Namespace).Get(deploymentName + "-deployment")
	// If the resource doesn't exist, we'll create it
	if errors.IsNotFound(err) {
		log.Printf("Error Finding deployment")

		//deployment, err = c.kubeclientset.AppsV1().Deployments(foo.Namespace).Create(context.TODO(), newDeployment(foo), metav1.CreateOptions{})
	}

	// If an error occurs during Get/Create, we'll requeue the item so we can
	// attempt processing again later. This could have been caused by a
	// temporary network failure, or any other transient reason.
	if err != nil {
		return err
	}

	// If the Deployment is not controlled by this Foo resource, we should log
	// a warning to the event recorder and return error msg.

	/*if !metav1.IsControlledBy(deployment, foo) {
		msg := fmt.Sprintf(MessageResourceExists, deployment.Name)
		c.recorder.Event(foo, corev1.EventTypeWarning, ErrResourceExists, msg)
		return fmt.Errorf("%s", msg)
	}*/

	if foo.Annotations["autoscaling.knative.dev/class"] == "ppa" {

		// If this number of the replicas on the Foo resource is specified, and the
		// number does not equal the current desired replicas on the Deployment, we
		// should update the Deployment resource.

		/*if foo.Spec.ContainerConcurrency != nil && *foo.Spec.Replicas != *deployment.Spec.Replicas {
			klog.V(4).Infof("Foo %s replicas: %d, deployment replicas: %d", name, *foo.Spec.Replicas, *deployment.Spec.Replicas)
			deployment, err = c.kubeclientset.AppsV1().Deployments(foo.Namespace).Update(context.TODO(), newDeployment(foo), metav1.UpdateOptions{})
		}*/

		deployment, err = c.kubeclientset.AppsV1().Deployments(foo.Namespace).Get(context.TODO(), foo.Name+"-deployment", metav1.GetOptions{})

		if err != nil {
			log.Printf("Error getting deployment #{foo.Name}")
			return nil
		}
		var x int32 = 2

		deploymentCopy := deployment.DeepCopy()
		deploymentCopy.Spec.Replicas = &x
		scale, err := c.kubeclientset.AppsV1().Deployments(foo.Namespace).GetScale(context.TODO(), foo.Name+"-deployment", metav1.GetOptions{})
		log.Printf("Current Scale '%v'", scale)
		var scalecopy = scale.DeepCopy()
		scalecopy.Spec.Replicas = x
		//_, err = c.kubeclientset.AppsV1().Deployments(foo.Namespace).UpdateScale(context.TODO(), foo.Name+"-deployment", &autoscalingv1.Scale{Spec: autoscalingv1.ScaleSpec{Replicas: 5}}, metav1.UpdateOptions{})
		_, err = c.kubeclientset.AppsV1().Deployments(foo.Namespace).UpdateScale(context.TODO(), foo.Name+"-deployment", scalecopy, metav1.UpdateOptions{})
		// If an error occurs during Update, we'll requeue the item so we can
		// attempt processing again later. This could have been caused by a
		// temporary network failure, or any other transient reason.
		if err != nil {
			return err
		}

		// Finally, we update the status block of the Foo resource to reflect the
		// current state of the world

		/*pa, err := c.sampleclientset.AutoscalingV1alpha1().PodAutoscalers(foo.Namespace).Get(context.TODO(), foo.Name, metav1.GetOptions{})
		if err != nil {
			log.Printf("Error getting PodAutoscaler %q: %v.", pa.Name, err)
			return err
		}
		log.Printf(pa.Name)
		pa.Status.MarkActive()
		pa.Status.SetConditions(apis.Conditions{{
			Type:               apis.ConditionReady,
			Status:             corev1.ConditionTrue,
			Reason:             "I am born",
			LastTransitionTime: apis.VolatileTime{Inner: metav1.Now()},
		}})
		log.Printf("'%v'", pa)
		y, err := c.sampleclientset.AutoscalingV1alpha1().PodAutoscalers(foo.Namespace).Update(context.TODO(), pa, metav1.UpdateOptions{})
		if err != nil {
			log.Printf("Error retrieving podautoscalar")
			return err
		}
		log.Printf("'%v'", y) */

		err = c.updateFooStatus(foo, deployment)
		if err != nil {
			log.Printf("Error retrieving podautoscalar")
			return err
		}
	}

	// Test Metrics Client
	var namespace1 = "default"
	var config = rest.Config{Host: "https://127.0.0.1:61560", APIPath: "/apis",
		BearerToken:     "eyJhbGciOiJSUzI1NiIsImtpZCI6Ii1QRmx4T2ZvYWllSHB2UFZlZmJWVW5YVlRsbUZleHVsa1BCUGVtemNxMU0ifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6Im1ldHJpY3NhZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJtZXRyaWNzLWFkbWluIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiMDRhOGFlMjItMjk3OS00NGRhLTk2MjQtMTI2M2I2OGU5ODM2Iiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlZmF1bHQ6bWV0cmljcy1hZG1pbiJ9.bk0IZJcZkmJep4JTYYL2BEemjHc6lhj9iT0dBUCk4rjQ9hlAr3cG-eOh_MdzUVDNC_a5eg5pinb4aUGzSKawxXN84OBcO46YNsjDnu9rqYhIE3FG6oAzZPkHbXPixyd0MjMdzHxHosdEYROzmUJiTDYNpHY4u2Z6tKRePUBzoWZM-fDPCazxl5YsKV3xE4ttmx6GWlx6Wo6Oyjusf9OTg_u-d56dZVDuZ5s1_ZfEF6mlBAZjT3V02zSVO1SJvhWV_YRLbBfWRS1NwgLZ8UWDdwUxu10dghKZd6wdj4X8bg7sF_7jvHbFvQj83CrR3k5IjIiDjOS3nklB2HN4-JpWyA",
		TLSClientConfig: rest.TLSClientConfig{Insecure: true}}
	var metricsClient = resourceclient.NewForConfigOrDie(&config)
	metrics, err := metricsClient.PodMetricses(namespace1).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf(fmt.Errorf("unable to fetch metrics from resource metrics API: %v", err).Error())
	}

	if len(metrics.Items) == 0 {
		log.Printf(fmt.Errorf("no metrics returned from resource metrics API").Error())
	}
	for _, m := range metrics.Items {
		if m.Name == "autoscale-go-00001-deployment-7df88b9699-jlj8z" {
			podSum := int64(0)
			missing := len(m.Containers) == 0
			for _, c := range m.Containers {
				resValue, found := c.Usage["memory"]
				if !found {
					missing = true
					break
				}
				podSum += resValue.MilliValue()
			}
			log.Printf("Error retrieving podautoscalar '%v'", missing)
			log.Printf("Pod Sum '%v' ", podSum)
		}
	}

	//________________
	c.recorder.Event(foo, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	return nil
}

func (c *Controller) updateFooStatus(foo *samplev1alpha1.PodAutoscaler, deployment *appsv1.Deployment) error {
	// NEVER modify objects from the store. It's a read-only, local cache.
	// You can use DeepCopy() to make a deep copy of original object and modify this copy
	// Or create a copy manually for better performance
	fooCopy := foo.DeepCopy()
	var x int32 = 2
	fooCopy.Status.DesiredScale = &x
	fooCopy.Status.MarkActive()
	fooCopy.Status.SetConditions(apis.Conditions{{
		Type:               apis.ConditionReady,
		Status:             corev1.ConditionTrue,
		Reason:             "I am born",
		LastTransitionTime: apis.VolatileTime{Inner: metav1.Now()},
	}})
	// If the CustomResourceSubresources feature gate is not enabled,
	// we must use Update instead of UpdateStatus to update the Status block of the Foo resource.
	// UpdateStatus will not allow changes to the Spec of the resource,
	// which is ideal for ensuring nothing other than resource status has been updated.
	_, err := c.sampleclientset.AutoscalingV1alpha1().PodAutoscalers(foo.Namespace).UpdateStatus(context.TODO(), fooCopy, metav1.UpdateOptions{})
	return err
}

// enqueueFoo takes a Foo resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than Foo.
func (c *Controller) enqueueFoo(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}

// handleObject will take any resource implementing metav1.Object and attempt
// to find the Foo resource that 'owns' it. It does this by looking at the
// objects metadata.ownerReferences field for an appropriate OwnerReference.
// It then enqueues that Foo resource to be processed. If the object does not
// have an appropriate OwnerReference, it will simply be skipped.
func (c *Controller) handleObject(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		klog.V(4).Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}
	klog.V(4).Infof("Processing object: %s", object.GetName())
	if ownerRef := metav1.GetControllerOf(object); ownerRef != nil {
		// If this object is not owned by a Foo, we should not do anything more
		// with it.
		if ownerRef.Kind != "Foo" {
			return
		}

		foo, err := c.foosLister.PodAutoscalers(object.GetNamespace()).Get(ownerRef.Name)
		if err != nil {
			klog.V(4).Infof("ignoring orphaned object '%s/%s' of foo '%s'", object.GetNamespace(), object.GetName(), ownerRef.Name)
			return
		}
		c.enqueueFoo(foo)
		return
	}
}

/*
// newDeployment creates a new Deployment for a Foo resource. It also sets
// the appropriate OwnerReferences on the resource so handleObject can discover
// the Foo resource that 'owns' it.
func newDeployment(foo *samplev1alpha1.PodAutoscaler) *appsv1.Deployment {
	labels := map[string]string{
		"app":        "nginx",
		"controller": foo.Name,
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      foo.Spec.DeploymentName,
			Namespace: foo.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(foo, samplev1alpha1.SchemeGroupVersion.WithKind("Foo")),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: foo.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx:latest",
						},
					},
				},
			},
		},
	}
}
*/

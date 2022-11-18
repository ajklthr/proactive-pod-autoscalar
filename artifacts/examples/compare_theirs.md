kubectl describe revision hello-00001
Name:         hello-00001
Namespace:    default
Labels:       serving.knative.dev/configuration=hello
serving.knative.dev/configurationGeneration=1
serving.knative.dev/configurationUID=b4054adb-64d2-4fee-8696-f850db4a02a6
serving.knative.dev/routingState=active
serving.knative.dev/service=hello
serving.knative.dev/serviceUID=f0b30dc4-e570-442f-ba17-88e3e41a6996
Annotations:  autoscaling.knative.dev/class: kpa.autoscaling.knative.dev
serving.knative.dev/creator: minikube-user
serving.knative.dev/routes: hello
serving.knative.dev/routingStateModified: 2022-11-18T00:53:17Z
API Version:  serving.knative.dev/v1
Kind:         Revision
Metadata:
Creation Timestamp:  2022-11-18T00:53:17Z
Generation:          1
Managed Fields:
API Version:  serving.knative.dev/v1
Fields Type:  FieldsV1
fieldsV1:
f:metadata:
f:annotations:
.:
f:autoscaling.knative.dev/class:
f:serving.knative.dev/creator:
f:serving.knative.dev/routes:
f:serving.knative.dev/routingStateModified:
f:labels:
.:
f:serving.knative.dev/configuration:
f:serving.knative.dev/configurationGeneration:
f:serving.knative.dev/configurationUID:
f:serving.knative.dev/routingState:
f:serving.knative.dev/service:
f:serving.knative.dev/serviceUID:
f:ownerReferences:
.:
k:{"uid":"b4054adb-64d2-4fee-8696-f850db4a02a6"}:
f:spec:
.:
f:containerConcurrency:
f:containers:
f:enableServiceLinks:
f:timeoutSeconds:
Manager:      controller
Operation:    Update
Time:         2022-11-18T00:53:17Z
API Version:  serving.knative.dev/v1
Fields Type:  FieldsV1
fieldsV1:
f:status:
.:
f:actualReplicas:
f:conditions:
f:containerStatuses:
f:desiredReplicas:
f:observedGeneration:
Manager:      controller
Operation:    Update
Subresource:  status
Time:         2022-11-18T00:53:21Z
Owner References:
API Version:           serving.knative.dev/v1
Block Owner Deletion:  true
Controller:            true
Kind:                  Configuration
Name:                  hello
UID:                   b4054adb-64d2-4fee-8696-f850db4a02a6
Resource Version:        533062
UID:                     9063014f-2e56-4cd9-b46e-0a414392c7ce
Spec:
Container Concurrency:  0
Containers:
Env:
Name:   TARGET
Value:  World
Image:    gcr.io/knative-samples/helloworld-go
Name:     user-container
Ports:
Container Port:  8080
Protocol:        TCP
Readiness Probe:
Success Threshold:  1
Tcp Socket:
Port:  0
Resources:
Enable Service Links:  false
Timeout Seconds:       300
Status:
Actual Replicas:  1
Conditions:
Last Transition Time:  2022-11-18T00:53:19Z
Severity:              Info
Status:                True
Type:                  Active
Last Transition Time:  2022-11-18T00:53:19Z
Status:                True
Type:                  ContainerHealthy
Last Transition Time:  2022-11-18T00:53:19Z
Status:                True
Type:                  Ready
Last Transition Time:  2022-11-18T00:53:19Z
Status:                True
Type:                  ResourcesAvailable
Container Statuses:
Image Digest:       gcr.io/knative-samples/helloworld-go@sha256:5ea96ba4b872685ff4ddb5cd8d1a97ec18c18fae79ee8df0d29f446c5efe5f50
Name:               user-container
Desired Replicas:     1
Observed Generation:  1
Events:
Type    Reason         Age   From                 Message
  ----    ------         ----  ----                 -------
Normal  RevisionReady  2s    revision-controller  Revision becomes ready upon all resources being ready





kubectl describe ksvc hello
Name:         hello
Namespace:    default
Labels:       <none>
Annotations:  serving.knative.dev/creator: minikube-user
serving.knative.dev/lastModifier: minikube-user
API Version:  serving.knative.dev/v1
Kind:         Service
Metadata:
Creation Timestamp:  2022-11-17T12:53:57Z
Generation:          1
Managed Fields:
API Version:  serving.knative.dev/v1
Fields Type:  FieldsV1
fieldsV1:
f:metadata:
f:annotations:
.:
f:kubectl.kubernetes.io/last-applied-configuration:
f:spec:
.:
f:template:
.:
f:metadata:
.:
f:annotations:
.:
f:autoscaling.knative.dev/class:
f:spec:
.:
f:containers:
Manager:      kubectl-client-side-apply
Operation:    Update
Time:         2022-11-17T12:53:57Z
API Version:  serving.knative.dev/v1
Fields Type:  FieldsV1
fieldsV1:
f:status:
.:
f:address:
.:
f:url:
f:conditions:
f:latestCreatedRevisionName:
f:latestReadyRevisionName:
f:observedGeneration:
f:traffic:
f:url:
Manager:         controller
Operation:       Update
Subresource:     status
Time:            2022-11-17T12:53:59Z
Resource Version:  439058
UID:               382b7cf2-b830-4686-bfc9-59caa95e836c
Spec:
Template:
Metadata:
Annotations:
autoscaling.knative.dev/class:  kpa.autoscaling.knative.dev
Creation Timestamp:               <nil>
Spec:
Container Concurrency:  0
Containers:
Env:
Name:   TARGET
Value:  World
Image:    gcr.io/knative-samples/helloworld-go
Name:     user-container
Ports:
Container Port:  8080
Protocol:        TCP
Readiness Probe:
Success Threshold:  1
Tcp Socket:
Port:  0
Resources:
Enable Service Links:  false
Timeout Seconds:       300
Traffic:
Latest Revision:  true
Percent:          100
Status:
Address:
URL:  http://hello.default.svc.cluster.local
Conditions:
Last Transition Time:        2022-11-17T12:53:59Z
Status:                      True
Type:                        ConfigurationsReady
Last Transition Time:        2022-11-17T12:53:59Z
Status:                      True
Type:                        Ready
Last Transition Time:        2022-11-17T12:53:59Z
Status:                      True
Type:                        RoutesReady
Latest Created Revision Name:  hello-00001
Latest Ready Revision Name:    hello-00001
Observed Generation:           1
Traffic:
Latest Revision:  true
Percent:          100
Revision Name:    hello-00001
URL:                http://hello.default.127.0.0.1.sslip.io
Events:
Type    Reason   Age   From                Message
  ----    ------   ----  ----                -------
Normal  Created  110s  service-controller  Created Configuration "hello"
Normal  Created  110s  service-controller  Created Route "hello"



kubectl describe PodAutoscaler hello-00001
Name:         hello-00001
Namespace:    default
Labels:       app=hello-00001
serving.knative.dev/configuration=hello
serving.knative.dev/configurationGeneration=1
serving.knative.dev/configurationUID=c804fe17-eb43-4ca8-aa81-5c34b6803ddc
serving.knative.dev/revision=hello-00001
serving.knative.dev/revisionUID=294581f3-faa0-4b19-a0a8-2605702dbbcd
serving.knative.dev/service=hello
serving.knative.dev/serviceUID=382b7cf2-b830-4686-bfc9-59caa95e836c
Annotations:  autoscaling.knative.dev/class: kpa.autoscaling.knative.dev
autoscaling.knative.dev/metric: concurrency
serving.knative.dev/creator: minikube-user
API Version:  autoscaling.internal.knative.dev/v1alpha1
Kind:         PodAutoscaler
Metadata:
Creation Timestamp:  2022-11-17T12:53:57Z
Generation:          2
Managed Fields:
API Version:  autoscaling.internal.knative.dev/v1alpha1
Fields Type:  FieldsV1
fieldsV1:
f:metadata:
f:annotations:
.:
f:autoscaling.knative.dev/class:
f:serving.knative.dev/creator:
f:labels:
.:
f:app:
f:serving.knative.dev/configuration:
f:serving.knative.dev/configurationGeneration:
f:serving.knative.dev/configurationUID:
f:serving.knative.dev/revision:
f:serving.knative.dev/revisionUID:
f:serving.knative.dev/service:
f:serving.knative.dev/serviceUID:
f:ownerReferences:
.:
k:{"uid":"294581f3-faa0-4b19-a0a8-2605702dbbcd"}:
f:spec:
.:
f:protocolType:
f:reachability:
f:scaleTargetRef:
Manager:      controller
Operation:    Update
Time:         2022-11-17T12:53:59Z
API Version:  autoscaling.internal.knative.dev/v1alpha1
Fields Type:  FieldsV1
fieldsV1:
f:status:
.:
f:actualScale:
f:conditions:
f:desiredScale:
f:metricsServiceName:
f:observedGeneration:
f:serviceName:
Manager:      autoscaler
Operation:    Update
Subresource:  status
Time:         2022-11-17T12:54:01Z
Owner References:
API Version:           serving.knative.dev/v1
Block Owner Deletion:  true
Controller:            true
Kind:                  Revision
Name:                  hello-00001
UID:                   294581f3-faa0-4b19-a0a8-2605702dbbcd
Resource Version:        439060
UID:                     6b9dcb1a-88e0-4ef3-8a23-b7327ffd8a86
Spec:
Protocol Type:  http1
Reachability:   Reachable
Scale Target Ref:
API Version:  apps/v1
Kind:         Deployment
Name:         hello-00001-deployment
Status:
Actual Scale:  1
Conditions:
Last Transition Time:  2022-11-17T12:53:59Z
Status:                True
Type:                  Active
Last Transition Time:  2022-11-17T12:53:59Z
Status:                True
Type:                  Ready
Last Transition Time:  2022-11-17T12:53:59Z
Status:                True
Type:                  SKSReady
Last Transition Time:  2022-11-17T12:53:59Z
Status:                True
Type:                  ScaleTargetInitialized
Desired Scale:           1
Metrics Service Name:    hello-00001-private
Observed Generation:     2
Service Name:            hello-00001
Events:                    <none>


kubectl describe deployment hello-00001-deployment
Name:                   hello-00001-deployment
Namespace:              default
CreationTimestamp:      Thu, 17 Nov 2022 07:53:57 -0500
Labels:                 app=hello-00001
serving.knative.dev/configuration=hello
serving.knative.dev/configurationGeneration=1
serving.knative.dev/configurationUID=c804fe17-eb43-4ca8-aa81-5c34b6803ddc
serving.knative.dev/revision=hello-00001
serving.knative.dev/revisionUID=294581f3-faa0-4b19-a0a8-2605702dbbcd
serving.knative.dev/service=hello
serving.knative.dev/serviceUID=382b7cf2-b830-4686-bfc9-59caa95e836c
Annotations:            autoscaling.knative.dev/class: kpa.autoscaling.knative.dev
deployment.kubernetes.io/revision: 1
serving.knative.dev/creator: minikube-user
Selector:               serving.knative.dev/revisionUID=294581f3-faa0-4b19-a0a8-2605702dbbcd
Replicas:               0 desired | 0 updated | 0 total | 0 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  0 max unavailable, 25% max surge
Pod Template:
Labels:       app=hello-00001
serving.knative.dev/configuration=hello
serving.knative.dev/configurationGeneration=1
serving.knative.dev/configurationUID=c804fe17-eb43-4ca8-aa81-5c34b6803ddc
serving.knative.dev/revision=hello-00001
serving.knative.dev/revisionUID=294581f3-faa0-4b19-a0a8-2605702dbbcd
serving.knative.dev/service=hello
serving.knative.dev/serviceUID=382b7cf2-b830-4686-bfc9-59caa95e836c
Annotations:  autoscaling.knative.dev/class: kpa.autoscaling.knative.dev
serving.knative.dev/creator: minikube-user
Containers:
user-container:
Image:      gcr.io/knative-samples/helloworld-go@sha256:5ea96ba4b872685ff4ddb5cd8d1a97ec18c18fae79ee8df0d29f446c5efe5f50
Port:       8080/TCP
Host Port:  0/TCP
Environment:
TARGET:           World
PORT:             8080
K_REVISION:       hello-00001
K_CONFIGURATION:  hello
K_SERVICE:        hello
Mounts:             <none>
queue-proxy:
Image:       ko.local/queue-39be6f1d08a095bd076a71d288d295b6:2550d4dff9c7a2ab8232ead213f6b3cbb74c5d2600badec4c9d24d5256681ff7
Ports:       8022/TCP, 9090/TCP, 9091/TCP, 8012/TCP, 8112/TCP
Host Ports:  0/TCP, 0/TCP, 0/TCP, 0/TCP, 0/TCP
Requests:
cpu:      25m
Readiness:  http-get http://:8012/ delay=0s timeout=1s period=10s #success=1 #failure=3
Environment:
SERVING_NAMESPACE:                        default
SERVING_SERVICE:                          hello
SERVING_CONFIGURATION:                    hello
SERVING_REVISION:                         hello-00001
QUEUE_SERVING_PORT:                       8012
QUEUE_SERVING_TLS_PORT:                   8112
CONTAINER_CONCURRENCY:                    0
REVISION_TIMEOUT_SECONDS:                 300
REVISION_RESPONSE_START_TIMEOUT_SECONDS:  0
REVISION_IDLE_TIMEOUT_SECONDS:            0
SERVING_POD:                               (v1:metadata.name)
SERVING_POD_IP:                            (v1:status.podIP)
SERVING_LOGGING_CONFIG:
SERVING_LOGGING_LEVEL:
SERVING_REQUEST_LOG_TEMPLATE:             {"httpRequest": {"requestMethod": "{{.Request.Method}}", "requestUrl": "{{js .Request.RequestURI}}", "requestSize": "{{.Request.ContentLength}}", "status": {{.Response.Code}}, "responseSize": "{{.Response.Size}}", "userAgent": "{{js .Request.UserAgent}}", "remoteIp": "{{js .Request.RemoteAddr}}", "serverIp": "{{.Revision.PodIP}}", "referer": "{{js .Request.Referer}}", "latency": "{{.Response.Latency}}s", "protocol": "{{.Request.Proto}}"}, "traceId": "{{index .Request.Header "X-B3-Traceid"}}"}
SERVING_ENABLE_REQUEST_LOG:               false
SERVING_REQUEST_METRICS_BACKEND:          prometheus
TRACING_CONFIG_BACKEND:                   none
TRACING_CONFIG_ZIPKIN_ENDPOINT:
TRACING_CONFIG_DEBUG:                     false
TRACING_CONFIG_SAMPLE_RATE:               0.1
USER_PORT:                                8080
SYSTEM_NAMESPACE:                         knative-serving
METRICS_DOMAIN:                           knative.dev/internal/serving
SERVING_READINESS_PROBE:                  {"tcpSocket":{"port":8080,"host":"127.0.0.1"},"successThreshold":1}
ENABLE_PROFILING:                         false
SERVING_ENABLE_PROBE_REQUEST_LOG:         false
METRICS_COLLECTOR_ADDRESS:
CONCURRENCY_STATE_ENDPOINT:
CONCURRENCY_STATE_TOKEN_PATH:             /var/run/secrets/tokens/state-token
HOST_IP:                                   (v1:status.hostIP)
ENABLE_HTTP2_AUTO_DETECTION:              false
Mounts:                                     <none>
Volumes:                                      <none>
Conditions:
Type           Status  Reason
  ----           ------  ------
Available      True    MinimumReplicasAvailable
Progressing    True    NewReplicaSetAvailable
OldReplicaSets:  <none>
NewReplicaSet:   hello-00001-deployment-985669b75 (0/0 replicas created)
Events:
Type    Reason             Age    From                   Message
  ----    ------             ----   ----                   -------
Normal  ScalingReplicaSet  2m19s  deployment-controller  Scaled up replica set hello-00001-deployment-985669b75 to 1
Normal



kubectl describe ingresses.networking.internal.knative.dev
Name:         hello
Namespace:    default
Labels:       serving.knative.dev/route=hello
serving.knative.dev/routeNamespace=default
serving.knative.dev/service=hello
Annotations:  networking.internal.knative.dev/rollout:
{"configurations":[{"configurationName":"hello","percent":100,"revisions":[{"revisionName":"hello-00001","percent":100}],"stepParams":{}}]...
networking.knative.dev/ingress.class: kourier.ingress.networking.knative.dev
serving.knative.dev/creator: minikube-user
serving.knative.dev/lastModifier: minikube-user
API Version:  networking.internal.knative.dev/v1alpha1
Kind:         Ingress
Metadata:
Creation Timestamp:  2022-11-17T23:46:29Z
Finalizers:
ingresses.networking.internal.knative.dev
Generation:  1
Managed Fields:
API Version:  networking.internal.knative.dev/v1alpha1
Fields Type:  FieldsV1
fieldsV1:
f:metadata:
f:annotations:
.:
f:networking.internal.knative.dev/rollout:
f:networking.knative.dev/ingress.class:
f:serving.knative.dev/creator:
f:serving.knative.dev/lastModifier:
f:labels:
.:
f:serving.knative.dev/route:
f:serving.knative.dev/routeNamespace:
f:serving.knative.dev/service:
f:ownerReferences:
.:
k:{"uid":"12dd9744-a558-4379-824f-8e6ea744226c"}:
f:spec:
.:
f:httpOption:
f:rules:
Manager:      controller
Operation:    Update
Time:         2022-11-17T23:46:29Z
API Version:  networking.internal.knative.dev/v1alpha1
Fields Type:  FieldsV1
fieldsV1:
f:metadata:
f:finalizers:
.:
v:"ingresses.networking.internal.knative.dev":
Manager:      kourier
Operation:    Update
Time:         2022-11-17T23:46:29Z
API Version:  networking.internal.knative.dev/v1alpha1
Fields Type:  FieldsV1
fieldsV1:
f:status:
.:
f:conditions:
f:observedGeneration:
f:privateLoadBalancer:
.:
f:ingress:
f:publicLoadBalancer:
.:
f:ingress:
Manager:      kourier
Operation:    Update
Subresource:  status
Time:         2022-11-17T23:46:29Z
Owner References:
API Version:           serving.knative.dev/v1
Block Owner Deletion:  true
Controller:            true
Kind:                  Route
Name:                  hello
UID:                   12dd9744-a558-4379-824f-8e6ea744226c
Resource Version:        517194
UID:                     79708e6c-0915-4c0b-942e-d7d45a4c464d
Spec:
Http Option:  Enabled
Rules:
Hosts:
hello.default
hello.default.svc
hello.default.svc.cluster.local
Http:
Paths:
Splits:
Append Headers:
Knative - Serving - Namespace:  default
Knative - Serving - Revision:   hello-00001
Percent:                          100
Service Name:                     hello-00001
Service Namespace:                default
Service Port:                     80
Visibility:                             ClusterLocal
Hosts:
hello.default.127.0.0.1.sslip.io
Http:
Paths:
Splits:
Append Headers:
Knative - Serving - Namespace:  default
Knative - Serving - Revision:   hello-00001
Percent:                          100
Service Name:                     hello-00001
Service Namespace:                default
Service Port:                     80
Visibility:                             ExternalIP
Status:
Conditions:
Last Transition Time:  2022-11-17T23:46:29Z
Status:                True
Type:                  LoadBalancerReady
Last Transition Time:  2022-11-17T23:46:29Z
Status:                True
Type:                  NetworkConfigured
Last Transition Time:  2022-11-17T23:46:29Z
Status:                True
Type:                  Ready
Observed Generation:     1
Private Load Balancer:
Ingress:
Domain Internal:  kourier-internal.kourier-system.svc.cluster.local
Public Load Balancer:
Ingress:
Domain Internal:  kourier.kourier-system.svc.cluster.local
Events:
Type    Reason           Age   From                Message
  ----    ------           ----  ----                -------
Normal  FinalizerUpdate  22s   ingress-controller  Updated "hello" finalizers


kubectl describe service hello
Name:              hello
Namespace:         default
Labels:            serving.knative.dev/route=hello
serving.knative.dev/service=hello
Annotations:       serving.knative.dev/creator: minikube-user
serving.knative.dev/lastModifier: minikube-user
Selector:          <none>
Type:              ExternalName
IP Families:       <none>
IP:
IPs:               <none>
External Name:     kourier-internal.kourier-system.svc.cluster.local
Port:              http2  80/TCP
TargetPort:        80/TCP
Endpoints:         <none>
Session Affinity:  None
Events:            <none>
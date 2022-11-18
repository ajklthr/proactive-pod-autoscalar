kubectl describe revision hello-00001
Name:         hello-00001
Namespace:    default
Labels:       serving.knative.dev/configuration=hello
serving.knative.dev/configurationGeneration=1
serving.knative.dev/configurationUID=46ebc798-18f1-4fb2-821b-ee428d5d502c
serving.knative.dev/routingState=active
serving.knative.dev/service=hello
serving.knative.dev/serviceUID=c54f0aa9-f4a3-427a-b4a3-0fdfcd15a70e
Annotations:  autoscaling.knative.dev/class: ppa
serving.knative.dev/creator: minikube-user
serving.knative.dev/routes: hello
serving.knative.dev/routingStateModified: 2022-11-18T00:23:22Z
API Version:  serving.knative.dev/v1
Kind:         Revision
Metadata:
Creation Timestamp:  2022-11-18T00:23:22Z
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
k:{"uid":"46ebc798-18f1-4fb2-821b-ee428d5d502c"}:
f:spec:
.:
f:containerConcurrency:
f:containers:
f:enableServiceLinks:
f:timeoutSeconds:
Manager:      controller
Operation:    Update
Time:         2022-11-18T00:23:22Z
API Version:  serving.knative.dev/v1
Fields Type:  FieldsV1
fieldsV1:
f:status:
.:
f:conditions:
f:containerStatuses:
f:observedGeneration:
Manager:      controller
Operation:    Update
Subresource:  status
Time:         2022-11-18T00:23:28Z
Owner References:
API Version:           serving.knative.dev/v1
Block Owner Deletion:  true
Controller:            true
Kind:                  Configuration
Name:                  hello
UID:                   46ebc798-18f1-4fb2-821b-ee428d5d502c
Resource Version:        525760
UID:                     f6352690-e499-4053-a641-95f9bc65d6e5
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
Conditions:
Last Transition Time:  2022-11-18T00:23:28Z
Severity:              Info
Status:                True
Type:                  Active
Last Transition Time:  2022-11-18T00:23:28Z
Status:                True
Type:                  ContainerHealthy
Last Transition Time:  2022-11-18T00:23:28Z
Status:                True
Type:                  Ready
Last Transition Time:  2022-11-18T00:23:28Z
Status:                True
Type:                  ResourcesAvailable
Container Statuses:
Image Digest:       gcr.io/knative-samples/helloworld-go@sha256:5ea96ba4b872685ff4ddb5cd8d1a97ec18c18fae79ee8df0d29f446c5efe5f50
Name:               user-container
Observed Generation:  1
Events:
Type     Reason         Age   From                 Message
  ----     ------         ----  ----                 -------
Normal   RevisionReady  26m   revision-controller  Revision becomes ready upon all resources being ready
Warning  InternalError  26m   revision-controller  failed to update deployment "hello-00001-deployment": Operation cannot be fulfilled on deployments.apps "hello-00001-deployment": the object has been modified; please apply your changes to the latest version and try again





kubectl describe ksvc hello
Name:         hello
Namespace:    default
Labels:       <none>
Annotations:  serving.knative.dev/creator: minikube-user
serving.knative.dev/lastModifier: minikube-user
API Version:  serving.knative.dev/v1
Kind:         Service
Metadata:
Creation Timestamp:  2022-11-17T12:29:51Z
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
Time:         2022-11-17T12:29:51Z
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
Time:            2022-11-17T12:33:34Z
Resource Version:  435355
UID:               2fe68898-6037-4e2f-99a7-de5ceb269144
Spec:
Template:
Metadata:
Annotations:
autoscaling.knative.dev/class:  ppa
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
Last Transition Time:        2022-11-17T12:33:34Z
Status:                      True
Type:                        ConfigurationsReady
Last Transition Time:        2022-11-17T12:33:34Z
Message:                     Waiting for load balancer to be ready
Reason:                      Uninitialized
Status:                      Unknown
Type:                        Ready
Last Transition Time:        2022-11-17T12:33:34Z
Message:                     Waiting for load balancer to be ready
Reason:                      Uninitialized
Status:                      Unknown
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
Normal  Created  18m   service-controller  Created Configuration "hello"
Normal  Created  18m   service-controller  Created Route "hello"


kubectl describe PodAutoscaler hello-00001
Name:         hello-00001
Namespace:    default
Labels:       app=hello-00001
serving.knative.dev/configuration=hello
serving.knative.dev/configurationGeneration=1
serving.knative.dev/configurationUID=655c3ed9-87b7-4ee1-b815-16fc5d19c7aa
serving.knative.dev/revision=hello-00001
serving.knative.dev/revisionUID=aac6906e-911f-4876-b949-2b58626330fb
serving.knative.dev/service=hello
serving.knative.dev/serviceUID=2fe68898-6037-4e2f-99a7-de5ceb269144
Annotations:  autoscaling.knative.dev/class: ppa
autoscaling.knative.dev/metric:
serving.knative.dev/creator: minikube-user
API Version:  autoscaling.internal.knative.dev/v1alpha1
Kind:         PodAutoscaler
Metadata:
Creation Timestamp:  2022-11-17T12:29:51Z
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
k:{"uid":"aac6906e-911f-4876-b949-2b58626330fb"}:
f:spec:
.:
f:protocolType:
f:reachability:
f:scaleTargetRef:
Manager:      controller
Operation:    Update
Time:         2022-11-17T12:33:34Z
API Version:  autoscaling.internal.knative.dev/v1alpha1
Fields Type:  FieldsV1
fieldsV1:
f:status:
.:
f:actualScale:
f:conditions:
f:metricsServiceName:
f:serviceName:
Manager:      ppa-controller
Operation:    Update
Subresource:  status
Time:         2022-11-17T12:52:04Z
Owner References:
API Version:           serving.knative.dev/v1
Block Owner Deletion:  true
Controller:            true
Kind:                  Revision
Name:                  hello-00001
UID:                   aac6906e-911f-4876-b949-2b58626330fb
Resource Version:        438575
UID:                     297232a8-8400-463b-aa40-f4bb32c540bb
Spec:
Protocol Type:  http1
Reachability:   Reachable
Scale Target Ref:
API Version:  apps/v1
Kind:         Deployment
Name:         hello-00001-deployment
Status:
Actual Scale:  5
Conditions:
Last Transition Time:  2022-11-17T12:52:04Z
Reason:                I am born
Status:                True
Type:                  Ready
Metrics Service Name:
Service Name:
Events:
Type    Reason  Age                   From               Message
  ----    ------  ----                  ----               -------
Normal  Synced  3m29s (x64 over 18m)  sample-controller  Foo synced successfully




kubectl describe deployment hello-00001-deployment
Name:                   hello-00001-deployment
Namespace:              default
CreationTimestamp:      Thu, 17 Nov 2022 07:29:51 -0500
Labels:                 app=hello-00001
serving.knative.dev/configuration=hello
serving.knative.dev/configurationGeneration=1
serving.knative.dev/configurationUID=655c3ed9-87b7-4ee1-b815-16fc5d19c7aa
serving.knative.dev/revision=hello-00001
serving.knative.dev/revisionUID=aac6906e-911f-4876-b949-2b58626330fb
serving.knative.dev/service=hello
serving.knative.dev/serviceUID=2fe68898-6037-4e2f-99a7-de5ceb269144
Annotations:            autoscaling.knative.dev/class: ppa
deployment.kubernetes.io/revision: 1
serving.knative.dev/creator: minikube-user
Selector:               serving.knative.dev/revisionUID=aac6906e-911f-4876-b949-2b58626330fb
Replicas:               5 desired | 5 updated | 5 total | 5 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  0 max unavailable, 25% max surge
Pod Template:
Labels:       app=hello-00001
serving.knative.dev/configuration=hello
serving.knative.dev/configurationGeneration=1
serving.knative.dev/configurationUID=655c3ed9-87b7-4ee1-b815-16fc5d19c7aa
serving.knative.dev/revision=hello-00001
serving.knative.dev/revisionUID=aac6906e-911f-4876-b949-2b58626330fb
serving.knative.dev/service=hello
serving.knative.dev/serviceUID=2fe68898-6037-4e2f-99a7-de5ceb269144
Annotations:  autoscaling.knative.dev/class: ppa
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
Progressing    True    NewReplicaSetAvailable
Available      True    MinimumReplicasAvailable
OldReplicaSets:  <none>
NewReplicaSet:   hello-00001-deployment-6f8db65b8c (5/5 replicas created)
Events:
Type     Reason                 Age   From                   Message
  ----     ------                 ----  ----                   -------
Warning  ReplicaSetCreateError  21m   deployment-controller  Failed to create new replica set "hello-00001-deployment-6f8db65b8c": Unauthorized
Normal   ScalingReplicaSet      21m   deployment-controller  Scaled up replica set hello-00001-deployment-6f8db65b8c to 1
Normal   ScalingReplicaSet      21m   deployment-controller  Scaled up replica set hello-00001-deployment-6f8db65b8c to 5 from 1



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
Creation Timestamp:  2022-11-17T23:39:40Z
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
k:{"uid":"f666df19-a34e-4158-bb6f-240eded018c7"}:
f:spec:
.:
f:httpOption:
f:rules:
Manager:      controller
Operation:    Update
Time:         2022-11-17T23:39:40Z
API Version:  networking.internal.knative.dev/v1alpha1
Fields Type:  FieldsV1
fieldsV1:
f:metadata:
f:finalizers:
.:
v:"ingresses.networking.internal.knative.dev":
Manager:      kourier
Operation:    Update
Time:         2022-11-17T23:39:40Z
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
Time:         2022-11-17T23:40:44Z
Owner References:
API Version:           serving.knative.dev/v1
Block Owner Deletion:  true
Controller:            true
Kind:                  Route
Name:                  hello
UID:                   f666df19-a34e-4158-bb6f-240eded018c7
Resource Version:        516068
UID:                     70a470e4-221a-4c3e-bc3e-db7671b439f2
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
Last Transition Time:  2022-11-17T23:40:44Z
Status:                True
Type:                  LoadBalancerReady
Last Transition Time:  2022-11-17T23:39:40Z
Status:                True
Type:                  NetworkConfigured
Last Transition Time:  2022-11-17T23:40:44Z
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
Type    Reason           Age    From                Message
  ----    ------           ----   ----                -------
Normal  FinalizerUpdate  4m35s  ingress-controller  Updated "hello" finalizers


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




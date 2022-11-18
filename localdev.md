minikube start --memory 7868 --cpus 6


https://minikube.sigs.k8s.io/docs/handbook/pushing/#1-pushing-directly-to-the-in-cluster-docker-daemon-docker-env
https://github.com/imjasonh/ko/blob/master/README.md#with-minikube
eval $(minikube -p minikube docker-env)

Controller Logs 

kubectl -n knative-serving logs $(kubectl -n knative-serving get pods -l app=controller -o name) -c controller


https://github.com/knative/client/blob/main/docs/README.md


kubectl get ksvc

https://minikube.sigs.k8s.io/docs/handbook/accessing/

 $ minikube service hello-minikube1 --url

 curl -H "Host: hello.default.127.0.0.1.sslip.io" http://192.168.49.2:80
 curl -verbose -H "Host: hello.default.127.0.0.1.sslip.io" http://127.0.0.1:80


https://knative.dev/docs/install/yaml-install/serving/install-serving-with-yaml/#install-a-networking-layer


kubectl patch configmap -n knative-serving config-domain -p "{\"data\": {\"127.0.0.1.sslip.io\": \"\"}}"


minikube tunnel

curl -H "Host: hello.default.127.0.0.1.sslip.io" http://127.0.0.1:80




kubectl --namespace kourier-system get service kourier

NAME      TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
kourier   LoadBalancer   10.108.226.60   127.0.0.1     80:31418/TCP,443:30271/TCP   3h21m
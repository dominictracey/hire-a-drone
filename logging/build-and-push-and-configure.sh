echo 'as per: http://stackoverflow.com/questions/36379572/how-to-setup-error-reporting-in-stackdriver-from-kubernetes-pods/36476771#36476771'
docker build -t gcr.io/###your project id###/fluentd-forwarder:v1 .
gcloud docker push gcr.io/###your project id###/fluentd-forwarder:v1
kubectl create -f fluentd-forwarder-controller.yaml
kubectl create -f fluentd-forwarder-service.yaml

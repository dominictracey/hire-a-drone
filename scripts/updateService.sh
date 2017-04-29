gcloud service-management deploy openapi.yaml
echo 'update args in k8s/production/api-production.yaml and kubectl apply -f k8s/production/api-production.yaml --namespace=production'
echo 'it also seems like you then have to push to git to get Jenkins to overwrite the api container because it goes back to the default name (v1)'

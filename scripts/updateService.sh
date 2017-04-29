gcloud service-management deploy openapi.yaml
echo 'update args in k8s/production/api-production.yaml and kubectl apply -f k8s/production/api-production.yaml --namespace=production'

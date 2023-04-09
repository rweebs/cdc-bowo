sed -i '' 's/app: green/app: blue/g' test/example/app/deploy/app-service.yaml
kubectl apply -f ./test/example/app/deploy/app-service.yaml
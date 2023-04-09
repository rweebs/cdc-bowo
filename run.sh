sed -i '' 's/app: blue/app: green/g' test/example/app/deploy/app-service.yaml
kubectl apply -f ./test/example/app/deploy/app-service.yaml
./main start blue-green
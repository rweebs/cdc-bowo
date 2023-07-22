kubectl patch service app -p '{"spec":{"selector":{"app": "blue"}}}'
kubectl exec $(kubectl get pods -l app=green -o jsonpath='{.items[0].metadata.name}') -- sh apk add redis
kubectl exec -it $(kubectl get pods -l app=green -o jsonpath='{.items[0].metadata.name}') sh 
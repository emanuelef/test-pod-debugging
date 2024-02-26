#k8s-local-debugging

Strt:
``` bash
kubectl apply -f deployment-service.yaml
kubectl apply -f deployment-service-second.yaml
kubectl apply -f deployment-service-go-app.yaml
kubectl apply -f deployment-service-rust-app.yaml
```

Stop:
``` bash
kubectl delete -f deployment-service.yaml
kubectl delete -f deployment-service-second.yaml
kubectl delete -f deployment-service-go-app.yaml
kubectl delete -f deployment-service-rust-app.yaml
```
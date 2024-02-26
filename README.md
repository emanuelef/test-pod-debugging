# k8s-local-debugging

This is the repo for the Medium Article [Mastering Local Microservices Debugging with Mirrord](https://medium.com/@emafuma/mastering-local-microservices-debugging-with-mirrord-0a99443c1544). 

Start:
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
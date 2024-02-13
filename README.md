`docker build -t simple-app .`

`docker run -p 8080:8080 simple-app`

``` bash
kubectl apply -f deployment-service.yaml
kubectl apply -f deployment-service-second.yaml
kubectl apply -f deployment-service-go-app.yaml
kubectl apply -f deployment-service-rust-app.yaml
```

``` bash
kubectl delete -f deployment-service.yaml
kubectl delete -f deployment-service-second.yaml
kubectl delete -f deployment-service-go-app.yaml
kubectl delete -f deployment-service-rust-app.yaml
```
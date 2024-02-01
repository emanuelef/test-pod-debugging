`docker build -t simple-app .`

`docker run -p 8098:8098 simple-app`

``` bash
kubectl apply -f deployment-service.yaml
kubectl apply -f deployment-service-second.yaml
```

``` bash
kubectl delete -f deployment-service.yaml
kubectl delete -f deployment-service-second.yaml
```

`kubectl expose deployment test-pod-debugging --type=LoadBalancer --port=8098`


`--kubeconfig /Users/efumagal/Workspace/Gravity/erdk-local-deployment/.kubeconfig`

kubectl apply -f deployment-service.yaml --kubeconfig /Users/efumagal/Workspace/Gravity/erdk-local-deployment/.kubeconfig
kubectl apply -f deployment-service-second.yaml --kubeconfig /Users/efumagal/Workspace/Gravity/erdk-local-deployment/.kubeconfig

kubectl delete -f deployment-service.yaml --kubeconfig /Users/efumagal/Workspace/Gravity/erdk-local-deployment/.kubeconfig
kubectl delete -f deployment-service-second.yaml --kubeconfig /Users/efumagal/Workspace/Gravity/erdk-local-deployment/.kubeconfig
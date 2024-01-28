`docker build -t simple-app .`

`docker run -p 8082:8082 simple-app`

`kubectl apply -f deployment.yaml`

`kubectl expose deployment test-pod-debugging --type=LoadBalancer --port=8082`
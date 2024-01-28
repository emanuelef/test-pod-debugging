`docker build -t simple-app .`

`docker run -p 8098:8098 simple-app`

`kubectl apply -f deployment.yaml`

`kubectl expose deployment test-pod-debugging --type=LoadBalancer --port=8098`
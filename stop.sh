#!/bin/bash

kubectl delete -f deployment-service.yaml
kubectl delete -f deployment-service-second.yaml
kubectl delete -f deployment-service-go-app.yaml
kubectl delete -f deployment-service-rust-app.yaml

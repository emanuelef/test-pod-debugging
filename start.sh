#!/bin/bash

kubectl apply -f deployment-service.yaml
kubectl apply -f deployment-service-second.yaml
kubectl apply -f deployment-service-go-app.yaml
kubectl apply -f deployment-service-rust-app.yaml

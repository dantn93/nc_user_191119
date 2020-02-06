#!/bin/sh
# kubectl delete service tranngocdan-nc-user-service
# kubectl delete deployment tranngocdan-nc-user
kubectl create -f provision/k8s/deployment.yaml
kubectl get service tranngocdan-nc-user-service

# minikube service tranngocdan-nc-user-service --url
#!/bin/bash

# Create skyramp cluster
skyramp config apply local

# Copy kubeconfig to default location

cp ~/.skyramp/kind-cluster.kubeconfig ~/.kube/config

# install worker
helm repo add skyramp https://letsramp.github.io/helm/
helm install mocker skyramp/worker -n default --set rbac=true

# deploy with skaffold

skaffold debug --port-forward

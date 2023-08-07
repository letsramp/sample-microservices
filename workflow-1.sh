#!/bin/bash

# Create skyramp cluster
skyramp config apply local

# Copy kubeconfig to default location
cp ~/.skyramp/kind-cluster.kubeconfig ~/.kube/config

# deploy with skaffold
skaffold debug 

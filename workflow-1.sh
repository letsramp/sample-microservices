#!/bin/bash

# Create skyramp cluster
skyramp cluster create --local

# Copy kubeconfig to default location
cp ~/.skyramp/kind-cluster.kubeconfig ~/.kube/config

# deploy with skaffold
skaffold debug

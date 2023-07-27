#!/bin/bash

helm repo add skyramp https://letsramp.github.io/helm/
helm install mocker skyramp/worker -n default --set rbac=true

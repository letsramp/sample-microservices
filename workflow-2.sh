#!/bin/bash

# Deploy Mocker

pushd /workspaces/sample-microservices/skyramp/rest-demo

skyramp mocker apply -n default
sleep 5

skyramp tester start checkout-test -n default
#!/bin/bash

# Deploy Mocker

pushd /workspaces/sample-microservices/skyramp/rest-demo

skyramp mocker apply -n default
skyramp tester start checkout-test -n default
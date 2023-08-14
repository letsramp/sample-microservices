#!/bin/bash

# Deploy Mocker

pushd /workspaces/sample-microservices/skyramp/rest-demo

skyramp mocker apply -n default

# run functional tests
skyramp tester start checkout-test -n default

# run load test
skyramp tester start test-load -n default

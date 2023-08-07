#!/bin/bash

cd /workspaces/sample-microservices/skyramp/docker-compose

echo "applying skyramp mock"
skyramp  mocker apply --address localhost:35142

echo "start tester"
skyramp tester start checkout-test --address localhost:35142
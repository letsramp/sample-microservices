#!/bin/bash

cd /workspaces/sample-microservices/skyramp/docker-compose

docker pull public.ecr.aws/j1n2c2p2/rampup/worker:latest
docker pull public.ecr.aws/j1n2c2p2/microservices-demo/productcatalogservice:v0.3.9
docker pull public.ecr.aws/j1n2c2p2/microservices-demo/checkoutservice:v0.4.1
docker pull public.ecr.aws/j1n2c2p2/microservices-demo/cartservice:v0.3.9

docker compose up


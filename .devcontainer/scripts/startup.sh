#!/bin/bash
set -eux

# Bringup Skyramp
# echo $ECR_TOKEN | docker login --username AWS --password-stdin 296613639307.dkr.ecr.us-west-2.amazonaws.com 
docker-compose -f /etc/skyramp/skyramp-docker-compose.yaml up -d

# Bringup sample-microservices
docker-compose -f ../../workspaces/sample-microservices/src/docker-compose.yml up -d
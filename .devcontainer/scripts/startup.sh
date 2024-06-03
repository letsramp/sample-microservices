#!/bin/bash
set -eux

# Bringup sample-microservices
docker-compose -f src/docker-compose.yml up -d

# Bringup Skyramp
echo $ECR_TOKEN | docker login --username AWS --password-stdin 296613639307.dkr.ecr.us-west-2.amazonaws.com 
docker-compose -f /etc/skyramp/skyramp-docker-compose.yaml up -d
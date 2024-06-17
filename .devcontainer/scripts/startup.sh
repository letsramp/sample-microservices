#!/bin/bash
set -eux

echo "cs name =============="
echo $SKY_CODESPACE_NAME

echo "new password =============="
echo $NEW_PASSWORD

"vscode:skyramp" | sudo chpasswd

# Bringup Skyramp
# echo $ECR_TOKEN | docker login --username AWS --password-stdin 296613639307.dkr.ecr.us-west-2.amazonaws.com 
docker-compose -f /etc/skyramp/skyramp-docker-compose.yaml up -d

# Bringup sample-microservices
docker-compose -f ../../workspaces/sample-microservices/src/docker-compose.yml up -d
#!/bin/bash
docker compose -f /workspaces/sample-microservices/skyramp/docker/demo/docker-compose.yml up -d --wait || true
nohup skyramp server up -vvv &
gh codespace ports visibility 45132:public -c $CODESPACE_NAME

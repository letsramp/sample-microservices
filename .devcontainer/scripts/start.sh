#!/bin/bash
gh codespace ports visibility 4000:public -c $CODESPACE_NAME
# Start Skyramp
skyramp server up > output.log 2>&1 & || true
docker compose -f /workspaces/sample-microservices/skyramp/docker/demo/docker-compose.yml up -d --wait || true
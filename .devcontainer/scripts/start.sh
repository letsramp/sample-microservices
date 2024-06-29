#!/bin/bash
# gh codespace ports visibility 4000:public -c $CODESPACE_NAME
docker compose -f /workspaces/sample-microservices/skyramp/docker/demo/docker-compose.yml up -d --wait || true
# #install skyramp cli curl -fsSL https://skyramp.dev/installer.sh
# curl -fsSL https://skyramp.dev/installer.sh | bash
# #run skyramp cli
# skyramp --version

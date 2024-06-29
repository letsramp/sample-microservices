#!/bin/bash
docker compose -f /workspaces/sample-microservices/skyramp/docker/demo/docker-compose.yml up -d --wait || true
nohup skyramp server up -vvv &
gh codespace ports visibility 45132:public -c $CODESPACE_NAME
apt-get update && apt-get install -y python3.11 python3-pip && update-alternatives --install /usr/bin/python python /usr/bin/python3.11 1
pip3 install -r /etc/skyramp/requirements.txt

#!/bin/bash
# install sample-microservices and skyramp
docker compose -f /etc/skyramp/docker-compose.yml up -d --wait || true
# start skyramp  server
nohup skyramp server up -vvv &
# open public port for skyramp dashboard
gh codespace ports visibility 4000:public -c $CODESPACE_NAME
gh codespace ports visibility 3000:public -c $CODESPACE_NAME
# install python3.11 and requirements
apt-get update && apt-get install -y python3.11 python3-pip && update-alternatives --install /usr/bin/python python /usr/bin/python3.11 1
pip3 install -r /etc/skyramp/requirements.txt

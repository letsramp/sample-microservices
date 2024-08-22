#!/bin/bash
# install skyramp
yes | bash -c "$(curl -fsSL https://skyramp.dev/installer.sh)"
# install skyramp worker and dashboard
# curl -o /etc/skyramp/docker-compose.yml https://raw.githubusercontent.com/letsramp/sample-microservices/script_branch/.devcontainer/scripts/docker-compose.yml
# docker compose -f /etc/skyramp/docker-compose.yml up -d --wait || true
skyramp dashboard up -d -vvv || true
# start skyramp  server
nohup skyramp server up -vvv &
# # open public port for skyramp dashboard
# gh codespace ports visibility 4000:public -c $CODESPACE_NAME
# gh codespace ports visibility 3000:public -c $CODESPACE_NAME
# install python3.11 and requirements
# apt-get update && apt-get install -y python3.11 python3-pip && update-alternatives --install /usr/bin/python python /usr/bin/python3.11 1
curl -o /etc/skyramp/requirements.txt https://raw.githubusercontent.com/letsramp/sample-microservices/script_branch/.devcontainer/scripts/requirements.txt
chmod +x /etc/skyramp/requirements.txt
pip install -r /etc/skyramp/requirements.txt

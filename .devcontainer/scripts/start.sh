#!/bin/bash

# install skyramp
yes | bash -c "$(curl -fsSL https://skyramp.dev/installer.sh)"

# install skyramp worker and dashboard
curl -o /etc/skyramp/docker-compose.yml https://raw.githubusercontent.com/letsramp/sample-microservices/script_branch/.devcontainer/scripts/docker-compose.yml
docker compose -f /etc/skyramp/docker-compose.yml up -d --wait || true

# start skyramp  server
nohup skyramp server up -vvv &

# install python requirements
curl -o /etc/skyramp/requirements.txt https://raw.githubusercontent.com/letsramp/sample-microservices/script_branch/.devcontainer/scripts/requirements.txt
chmod +x /etc/skyramp/requirements.txt
/usr/local/python/current/bin/pip install -r /etc/skyramp/requirements.txt

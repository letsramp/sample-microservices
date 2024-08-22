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
# pip install -r /etc/skyramp/requirements.txt
# /usr/local/python/3.11.9/bin/pip install -r /etc/skyramp/requirements.txt
# Step 1: Detect the installed Python version
PYTHON_VERSION=$(python3 --version 2>&1 | awk '{print $2}')
# Example output: 3.11.9

# Step 2: Construct the path to pip based on the Python version
PIP_PATH="/usr/local/python/${PYTHON_VERSION}/bin/pip"

# Step 3: Check if the pip path exists
if [ -x "$PIP_PATH" ]; then
    # Step 4: Use pip to install the requirements
    $PIP_PATH install -r /etc/skyramp/requirements.txt
else
    echo "pip not found at $PIP_PATH"
    exit 1
fi
#!/bin/bash

# Start SSH service
sudo service ssh start

# Set up SSH tunnel
PORT=${1:-6363}
ssh -R $PORT:localhost:2222 serveo.net

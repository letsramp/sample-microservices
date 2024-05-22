FROM mcr.microsoft.com/devcontainers/base:alpine

RUN apk add bash

RUN apk update && \
    apk add --no-cache docker-cli python3 && \
    apk add --no-cache --virtual .docker-compose-deps python3-dev libffi-dev openssl-dev gcc libc-dev make && \
    pip3 install docker-compose && \
    apk del .docker-compose-deps

# Install Skyramp
RUN bash -c "$(curl -fsSL https://skyramp.dev/installer.sh)"
 
# # Add Skyramp docker-compose manifest
RUN cat <<EOF > /tmp/skyramp-docker-compose.yaml
services:
    skyramp:
      image: public.ecr.aws/j1n2c2p2/rampup/worker:latest
      volumes:
        - skyramp:/etc/skyramp
        - /var/run/docker.sock:/var/run/docker.sock
      ports:
        - "35142:35142"
      restart: always
    # The services below (dashboard-server, dashboard-client, and mongodb)
    # can be included optionally for Skyramp Dashboard bring up
    dashboard-server:
      image: public.ecr.aws/j1n2c2p2/rampup/dashboard-server:latest
      environment:
        NODE_ENV: production
        DOCKER_ENV: true 
      volumes:
        - /var/run/docker.sock:/var/run/docker.sock
      ports:
        - "4000:4000"
      restart: always
    dashboard-client:
      image: public.ecr.aws/j1n2c2p2/rampup/dashboard-client:latest
      environment:
        NODE_ENV: production
      volumes:
        - /var/run/docker.sock:/var/run/docker.sock
      ports:
        - "3000:3000"
      restart: always
    mongodb:
      image: mongo:6.0.6
      container_name : mongo-datastore-svc
      mem_limit: 2g
      ports:
        - "27017:27017"
      environment:
        MONGO_INITDB_ROOT_USERNAME: admin
        MONGO_INITDB_ROOT_PASSWORD: skyramp
        MONGO_INITDB_DATABASE: dashboarddb
      volumes:
        - mongodb:/data/db
      restart: always
  volumes:
    mongodb:
    skyramp:
EOF

# Bringup Skyramp worker
RUN docker compose -f /tmp/skyramp-docker-compose.yaml up -d

# Note: You can use any Debian/Ubuntu based image you want. 
FROM mcr.microsoft.com/vscode/devcontainers/base:buster

# Install Node.js and npm 18
RUN apt-get update \
    && apt-get install -y curl \
    && curl -fsSL https://deb.nodesource.com/setup_18.x | bash - \
    && apt-get install -y nodejs

# install python3 and pip
RUN apt-get update \
    && apt-get install -y python3 python3-pip

VOLUME [ "/var/lib/docker" ]

# Setting the ENTRYPOINT to docker-init.sh will start up the Docker Engine 
# inside the container "overrideCommand": false is set in devcontainer.json. 
# The script will also execute CMD if you need to alter startup behaviors.
ENTRYPOINT [ "/usr/local/share/docker-init.sh" ]

# Install Skyramp
RUN bash -c "$(curl -fsSL https://skyramp.dev/installer.sh)"

COPY scripts/start.sh /etc/skyramp/start.sh
RUN chmod +x /etc/skyramp/start.sh

CMD [ "sleep", "infinity" ]

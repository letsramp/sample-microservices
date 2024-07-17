# use the latest version of the Python 3.11 container image
FROM mcr.microsoft.com/devcontainers/python:3.11

# install python3 and pip
RUN apt-get update \
    && apt-get install -y python3.11 python3-pip

# install requirements.txt
COPY scripts/requirements.txt /etc/skyramp/requirements.txt

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

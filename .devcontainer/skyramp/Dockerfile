# use the latest version of the Python 3.11 container image
FROM mcr.microsoft.com/devcontainers/python:3.11

# install requirements.txt
COPY scripts/requirements.txt /etc/skyramp/requirements.txt

VOLUME [ "/var/lib/docker" ]

# setting the ENTRYPOINT to docker-init.sh will start up the Docker Engine 
ENTRYPOINT [ "/usr/local/share/docker-init.sh" ]

# install Skyramp
RUN bash -c "$(curl -fsSL https://skyramp.dev/installer.sh)"

# startup script
COPY scripts/start.sh /etc/skyramp/start.sh
RUN chmod +x /etc/skyramp/start.sh

# copy the Skyramp configuration file
COPY scripts/docker-compose.yml /etc/skyramp/docker-compose.yml
RUN chmod +x /etc/skyramp/docker-compose.yml

CMD [ "sleep", "infinity" ]

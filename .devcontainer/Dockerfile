# use the latest version of the Python 3.11 container image
FROM mcr.microsoft.com/devcontainers/python:3.11

# install requirements.txt
COPY scripts/requirements.txt /etc/skyramp/requirements.txt

VOLUME [ "/var/lib/docker" ]

# startup script
ENTRYPOINT [ "/usr/local/share/docker-init.sh" ]

# Install Skyramp
RUN bash -c "$(curl -fsSL https://skyramp.dev/installer.sh)"

COPY scripts/start.sh /etc/skyramp/start.sh
RUN chmod +x /etc/skyramp/start.sh

CMD [ "sleep", "infinity" ]

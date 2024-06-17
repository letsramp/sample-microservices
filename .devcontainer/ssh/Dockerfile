# Use an appropriate base image
FROM mcr.microsoft.com/vscode/devcontainers/base:ubuntu

# Add a new user and set up sudo privileges
ARG USERNAME=user
ARG USER_UID=1000
ARG USER_GID=$USER_UID
ARG USER_PASSWORD=password

RUN apt-get update && apt-get install -y \
    sudo \
    openssh-server \
    && rm -rf /var/lib/apt/lists/* \
    && groupadd --gid $USER_GID $USERNAME \
    && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && echo "$USERNAME:$USER_PASSWORD" | chpasswd \
    && echo "$USERNAME ALL=(ALL:ALL) ALL" >> /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME \
    && mkdir /var/run/sshd \
    && echo 'PasswordAuthentication yes' >> /etc/ssh/sshd_config \
    && echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config \
    && service ssh start

EXPOSE 22

# Start SSH service
CMD ["/usr/sbin/sshd", "-D"]

# Use an appropriate base image
FROM mcr.microsoft.com/vscode/devcontainers/base:ubuntu

# Add a new user and set up sudo privileges
ARG USERNAME=skyramp
ARG USER_UID=100011
ARG USER_GID=$USER_UID

RUN apt-get update && apt-get install -y \
    sudo \
    openssh-server \
    && rm -rf /var/lib/apt/lists/* \
    && groupadd --gid $USER_GID $USERNAME \
    && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && echo "$USERNAME ALL=(ALL:ALL) ALL" >> /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME \
    && mkdir /var/run/sshd \
    && echo 'PasswordAuthentication yes' >> /etc/ssh/sshd_config \
    && echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config \
    && service ssh start

# Set up SSH key (if needed)
# COPY your_public_key.pub /home/$USERNAME/.ssh/authorized_keys
# RUN chown -R $USERNAME:$USERNAME /home/$USERNAME/.ssh

EXPOSE 22

# Start SSH service
CMD ["/usr/sbin/sshd", "-D"]

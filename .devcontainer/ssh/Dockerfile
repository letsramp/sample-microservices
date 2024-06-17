# Use an appropriate base image
FROM mcr.microsoft.com/vscode/devcontainers/base:ubuntu

# Add a new user and set up sudo privileges
ARG USERNAME=user
ARG USER_UID=100011
ARG USER_GID=$USER_UID
ARG USER_PASSWORD=password

RUN apt-get update && apt-get install -y sudo openssh-server

# Create the user and set the password
RUN groupadd --gid $USER_GID $USERNAME && \
    useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME && \
    echo "$USERNAME:$USER_PASSWORD" | chpasswd && \
    echo "$USERNAME ALL=(ALL:ALL) ALL" >> /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME

# # Ensure SSH directory and keys are created
# RUN mkdir /var/run/sshd && \
#     ssh-keygen -A

EXPOSE 2222

# Start SSH service in foreground
CMD ["/usr/sbin/sshd", "-D"]

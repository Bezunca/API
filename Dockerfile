# === Pre Stages global arguments ===========================================
# --- Project Name ---
ARG PROJECT_NAME="bezuncapi"

# --- Create user ---
ARG USERNAME="bezunca"

# === Stage 1 - Build go project ===========================================
FROM golang:1.14.4-buster AS builder
ARG PROJECT_NAME
ARG USERNAME

# --- Environment Variables ---
# Don't allow APT to make question
ENV DEBIAN_FRONTEND=noninteractive

# Add APT config file
ADD "https://gist.githubusercontent.com/HeavenVolkoff/ff7b77b9087f956b8df944772e93c071/raw" /etc/apt/apt.conf.d/99custom

# Update APT
RUN apt-get update -qq \
    && \
    # Install build requirements
    apt-get install \
        git \
        ssh \
        curl \
        upx \
        build-essential \
        openssh-client \
    && \
    git config --global url."git@github.com:".insteadOf "https://github.com/"

# Setting up GOPRIVATE env so we can download from gitlab's private repositories
ENV GOPRIVATE="github.com/Bezunca/*"

# Create build directory
WORKDIR /src/proj

# Build argument. Link for the tar file containing the git ssh key.
ARG SSH_KEY_LINK
ARG SSH_KEY
# Ensure SSH_KEY_LINK link is not empty
RUN test -n "$SSH_KEY_LINK" || test -n "$SSH_KEY" || ( echo "You must provide SSH_KEY_LINK or SSH_KEY" && exit 1 )

# Get SSH keys
RUN test -n "$SSH_KEY_LINK" &&  ( curl -k -# -L ${SSH_KEY_LINK} | tar -C /root -x || exit 1 ) || true

RUN test -n "$SSH_KEY" \
    && ( \
        mkdir -p ~/.ssh \
        && \
        chmod 700 ~/.ssh \
        && \
        echo "$SSH_KEY" | tr -d '\r' > ~/.ssh/id_rsa \
        && \
        chmod 600 ~/.ssh/id_rsa \
    ) || true

RUN ssh-keyscan -p 22 github.com > ~/.ssh/known_hosts \
    && \
    chmod 644 ~/.ssh/known_hosts

# Copying go modules files
COPY go.mod .
COPY go.sum .

# Downloading dependencies
RUN go mod download

# Copy project sources
COPY . .

# Building project executable, cleaning useless stuff and compressing binary
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o "executable" ./cmd/${PROJECT_NAME}
#&& upx --best --ultra-brute ${PROJECT_NAME}

# Fix permissions and create unprivileged user
RUN useradd -b /home -s /bin/sh -u 1001 -g 65534 ${USERNAME} \
    && \
    # Setup data volumes directories
    install -g 65534 -o 1001 -d /home/${USERNAME}/logs \
    && \
    # Remove setuid and setgid permissions
    find / -perm /6000 -type f -exec chmod a-s {} \; || true

# === Stage 2 - Setup runtime ==================================================
FROM ubuntu
ARG PROJECT_NAME
ARG USERNAME

# Copy project data
COPY --from=builder /src/proj/executable /usr/local/bin

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /home/  /home/

# Setup runtime
USER ${USERNAME}
VOLUME /home/${USERNAME}/logs
WORKDIR /home/${USERNAME}

# Exposed ports
EXPOSE 8080

RUN echo ${PROJECT_NAME}
# Run application
ENTRYPOINT ["/usr/local/bin/executable"]
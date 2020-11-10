FROM jenkins/jnlp-slave:alpine

LABEL maintainer="joewalker"

USER root

# ENV PATH=$PATH:/home/sonar-scanner/bin;/home/dependency-check/bin
ENV SDKMAN_DIR=/root/.sdkman

COPY assets .
COPY scripts/setup.sh .

RUN apk update && \
    apk add --no-cache \
    ca-certificates \
    zip \
    unzip \
    git \
    wget \
    curl \
    docker \
    maven \
    jq

RUN chmod u+x *.sh && \
    bash ./setup.sh && \
    rm -f setup.sh

VOLUME [ "/var/run/docker.sock" ]
FROM jenkinsci/jnlp-slave:3.35-2-alpine

LABEL maintainer="joewalker"

USER root

ENV NEXUS_USER="admin"
ENV NEXUS_PASS="admin"
ENV NEXUS_URL="10.101.16.52:8080"
ENV PATH=$PATH:/home/sonar-scanner/bin;/home/dependency-check/bin

COPY setup.sh .
COPY settings.xml /tmp/settings.xml

RUN apk update && \
    apk add --no-cache \
    ca-certificates \
    zip \
    unzip \
    git \
    wget \
    curl \
    docker \
    openjdk8 \
    maven

RUN chmod u+x setup.sh && \
    bash ./setup.sh

VOLUME [ "/var/run/docker.sock" ]
USER jenkins
FROM centos:7

LABEL maintainer="joewalker"
ARG USER
ARG PASS

ENV JAVA_HOME=/root/
ENV SONAR_RUNNER_HOME=/home/sonar-scanner
ENV PATH=$PATH:${SONAR_RUNNER_HOME}/bin

COPY entrypoint.sh .
COPY sshd_config /etc/ssh/sshd_config

RUN chmod u+x entrypoint.sh && \
    ./entrypoint.sh

VOLUME [ "/var/run/docker.sock" ]
EXPOSE 22
CMD [ "/usr/sbin/sshd", "-D"]

FROM centos:7

LABEL maintainer="joewalker"
ARG USER
ARG PASS

COPY entrypoint.sh .

RUN chmod u+x entrypoint.sh && \
    ./entrypoint.sh

VOLUME [ "/var/run/docker.sock" ]
EXPOSE 22
CMD [ "/usr/sbin/sshd", "-D"]
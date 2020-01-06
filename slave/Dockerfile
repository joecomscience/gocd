FROM centos:7

LABEL maintainer="joewalker"
ARG PASS
ARG USER

COPY entrypoint.sh .

RUN chmod u+x entrypoint.sh && \
    ./entrypoint.sh
CMD [ "/usr/sbin/sshd", "-D"]
FROM jenkins/jnlp-slave:alpine as jenkins-base
WORKDIR /home
COPY assets .
COPY scripts/setup.sh .

RUN chmod u+x *.sh && \
    bash ./setup.sh && \
    rm -f setup.sh

FROM openjdk:8-alpine
LABEL maintainer="joewalker"

RUN apk update && \
    apk add --no-cache \
    git \
    docker \
    maven \
    gradle \
    jq

COPY --from=jenkins-base /usr/share/jenkins/ /usr/share/jenkins/
COPY --from=jenkins-base /home/assets/libs/transformxml /usr/local/bin
COPY --from=jenkins-base /usr/local/bin/oc /usr/local/bin
COPY --from=jenkins-base /home/assets/maven_settings.xml /usr/share/java/maven-3/conf/settings.xml

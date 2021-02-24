FROM jenkins/jenkins:lts as jenkins-base
COPY ./assets/certs/ca.crt /home/ca.crt
COPY ./assets/certs/aycapsu01ts472.cer /home/
COPY ./assets/certs/aycapsu01ts472.dso.aycap.bayad.co.th.pem /home/aycapsu01ts472.pem

WORKDIR /home
USER root

RUN openssl \
    pkcs12 \
    -export \
    -out jenkins_keystore.p12 \
    -passout 'pass:changeit' \
    -inkey aycapsu01ts472.pem \
    -in aycapsu01ts472.cer \
    -certfile ca.crt \
    -name aycapsu01ts472.dso.aycap.bayad.co.th && \
    keytool \
    -importkeystore \
    -srckeystore jenkins_keystore.p12 \
    -srcstorepass 'changeit' \
    -srcstoretype PKCS12 \
    -srcalias aycapsu01ts472.dso.aycap.bayad.co.th \
    -deststoretype JKS \
    -destkeystore jenkins_keystore.jks \
    -deststorepass 'changeit' \
    -destalias aycapsu01ts472.dso.aycap.bayad.co.th

FROM jenkins/jenkins:lts
USER root
COPY --from=jenkins-base /home/ca.crt /usr/local/share/ca-certificates/ca.crt
COPY --from=jenkins-base /home/jenkins_keystore.jks /var/jenkins_home/

RUN update-ca-certificates

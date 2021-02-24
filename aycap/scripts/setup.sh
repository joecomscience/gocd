#!/bin/bash

set -e

ca_file=/usr/local/share/ca-certificates/ca.crt

chmod u+x ./libs/transformxml
mv ./libs/transformxml /usr/local/bin

mv maven_settings.xml /usr/share/java/maven-3/conf/settings.xml

wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub
wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.30-r0/glibc-2.30-r0.apk
apk add --no-cache glibc-2.30-r0.apk

echo '--install openjfx--'
wget --quiet --output-document=/etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub
wget https://github.com/sgerrand/alpine-pkg-java-openjfx/releases/download/8.151.12-r0/java-openjfx-8.151.12-r0.apk
apk add --no-cache java-openjfx-8.151.12-r0.apk

# echo '--install sonar-scanner--'
# wget -O sonar-scanner-cli-4.2.0.1873.zip  https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.2.0.1873.zip
# unzip sonar-scanner-cli-4.2.0.1873.zip -d /home
# rm -f sonar-scanner-cli-4.2.0.1873.zip
# cd /home
# mv sonar-scanner-4.2.0.1873 sonar-scanner

echo '--install openshift cli--'
wget -O openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit.tar.gz https://github.com/openshift/origin/releases/download/v3.11.0/openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit.tar.gz
tar xvzf openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit.tar.gz
cp ./openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit/oc /usr/local/bin
rm -rf openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit*

<<<<<<< HEAD
echo '--install internal cert--'
update-ca-certificates
# keytool -import -storepass changeit -file /usr/share/ca-certificates/master.crt -keystore /etc/ssl/certs/java/cacerts -noprompt -alias jenkins_master

=======
echo '--install sdkman manage java version--'
curl -s "https://get.sdkman.io" | bash
rm -rf /var/lib/apt/lists/*
echo "sdkman_auto_answer=true" > $SDKMAN_DIR/etc/config
echo "sdkman_auto_selfupdate=false" >> $SDKMAN_DIR/etc/config
echo "sdkman_insecure_ssl=true" >> $SDKMAN_DIR/etc/config

source $SDKMAN_DIR/bin/sdkman-init.sh

sdk install java $SDKMAN_JAVA_VERSION

echo '--install internal cert--'
update-ca-certificates
java_default=$SDKMAN_DIR/candidates/java/current/jre/lib/security/cacerts

keytool \
-import \
-storepass changeit \
-file $ca_file \
-keystore $java_default \
-noprompt \
-alias jenkins_master
>>>>>>> dccac8c05c8e20e6490a40ae83e0e64489b76c61

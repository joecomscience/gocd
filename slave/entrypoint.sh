#!/bin/bash

set -e
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
yum update -y
yum install -y \
openssh-server \
zip \
unzip \
which \
git \
docker

echo '--check passwd--'
# useradd ${USER}
# sed -i -e 's/#PermitRootLogin/PermitRootLogin/g' sshd_config
echo "${USER}:${PASS}" | chpasswd -c SHA512 

echo '--gen key--'
ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key -N ''
ssh-keygen -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key -N ''
ssh-keygen -t ed25519 -f /etc/ssh/ssh_host_ed25519_key -N ''

echo '--install skdman--'
curl -s "https://get.sdkman.io" | bash > /dev/null
source "${HOME}/.sdkman/bin/sdkman-init.sh"
sdk version

echo '--install dependency--'
java_version="8.0.242.hs-adpt"

sdk install java ${java_version}
sdk install maven

mkdir -p /usr/local/java
ln -s ~/.sdkman/candidates/java/current/bin/ /usr/local/java/bin

echo '--install sonar-scanner--'

curl https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.2.0.1873-linux.zip -o sonar-scanner-cli-4.2.0.1873-linux.zip
unzip sonar-scanner-cli-4.2.0.1873-linux.zip -d /home
rm -rf sonar-scanner-cli-4.2.0.1873-linux.zip
cd /home
mv sonar-scanner-4.2.0.1873-linux sonar-scanner

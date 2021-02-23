#!/bin/bash

set -e

chmod u+x ./libs/transformxml

echo '--install openshift cli--'
wget -O openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit.tar.gz https://github.com/openshift/origin/releases/download/v3.11.0/openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit.tar.gz
tar xvzf openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit.tar.gz
cp ./openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit/oc /usr/local/bin
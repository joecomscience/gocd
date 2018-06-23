#!/bin/bash
set -o pipefail
IFS=$'\n\t'

echo $BUILD | jq . - > /tmp/build-spec.json

cat /tmp/build-spec.json | jq -r '.spec.strategy.customStrategy.env | map([ "export " + .name, "\"" + .value + "\""] | join("=")) | join(" \n")' > /tmp/env-vars

source /tmp/env-vars

if [[ "$DEBUG" = "true" ]]; then
  echo "-- build spec"
  cat /tmp/build-spec.json
  echo
  echo

  echo "-- ENV Vars"
  cat /tmp/env-vars
  echo
  echo
  set -x
fi

if [[ "${SOURCE_REPOSITORY}" != "git://"* ]] && [[ "${SOURCE_REPOSITORY}" != "git@"* ]]; then
  URL="${SOURCE_REPOSITORY}"
  if [[ "${URL}" != "http://"* ]] && [[ "${URL}" != "https://"* ]]; then
    URL="https://${URL}"
  fi
fi

export URL
GIT_USER="$(cat /var/run/secrets/openshift.io/source/username)"
GIT_PASS="$(cat /var/run/secrets/openshift.io/source/password)"
SOURCE_REF="${SOURCE_REF-master}"
echo "git-info: ${URL}@${SOURCE_REF}:${SOURCE_CONTEXT_DIR-./}"
URL="$(echo $URL | sed -e s%://%://$GIT_USER:$GIT_PASS@%g)"
SOURCE_REPOSITORY="${URL}"

if [ -n "${SOURCE_REF}" ]; then
  BUILD_DIR=$(mktemp -d)
  git clone --recursive "${SOURCE_REPOSITORY}" "${BUILD_DIR}"
  if [ $? != 0 ]; then
    echo "Error trying to fetch git source: ${SOURCE_REPOSITORY}"
    exit 1
  fi
  pushd "${BUILD_DIR}"
  git checkout "${SOURCE_REF}"
  if [ $? != 0 ]; then
    echo "Error trying to checkout branch: ${SOURCE_REF}"
    exit 1
  fi

  echo -- source-ref
  if [ -n "$SOURCE_CONTEXT_DIR" ]; then
    cd ${SOURCE_CONTEXT_DIR}
  fi

  ls -lah ./

  echo -- Going to apply this config after substitution
  cat ${CONFIGURATION_FILE}
  echo --

  cat ${CONFIGURATION_FILE} | envsubst > substituted-config.yaml

  if [[ "$DEBUG" = "true" ]]; then
    echo -- Config file w/ substitutions
    cat substituted-config.yaml
    echo -- end of substituted config
  fi

  gocd_configurator substituted-config.yaml
  export EXIT_CODE="$?"

  popd
else
  echo -- no source-ref
fi

exit "$EXIT_CODE"

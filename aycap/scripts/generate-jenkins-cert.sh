
docker run --rm --init joewalker/jenkins-slave:jdk8 -url https://host.docker.internal:8443 -secret f92e08e5bd31f4cf5bc95a603cd6bf2447e1848e012de597611c5d63682af0b5 slave-1

docker run --rm --init joewalker/jenkins-slave:jdk8 -direct host.docker.internal:8443

0C CB 5D 8D 21 EB 77 6C 4C 6A 97 04 25 F8 7D 4C 0A 65 FB 6C
0A:65:FB:6C

docker run --rm --name=agent1 -p 3000:22 \
-e "JENKINS_AGENT_SSH_PUBKEY=~/home/jenkins/.ssh/jenkins_agent_key.pub" \
jenkins/ssh-agent:alpine

VARS1="HOME=|USER=|MAIL=|LC_ALL=|LS_COLORS=|LANG="
VARS2="HOSTNAME=|PWD=|TERM=|SHLVL=|LANGUAGE=|_="
VARS="${VARS1}|${VARS2}"
docker exec agent1 sh -c "env | egrep -v "^(${VARS})" >> /etc/environment"
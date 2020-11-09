build8:
	DOCKER_BUILDKIT=1 docker build -t joewalker/jenkins-slave:jdk8 -f dockerfile/slave/. .

build7:
	DOCKER_BUILDKIT=1 docker build -t joewalker/jenkins-slave:jdk7 --arg JAVA_VERSION="openjdk7" -f dockerfile/slave/. .

push:
	docker push joewalker/jenkins-slave
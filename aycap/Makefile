jdk7:
	DOCKER_BUILDKIT=1 docker build -t joewalker/jenkins-slave:jdk7 -f dockerfile/slave/Dockerfile-jdk7 .

jdk8:
	DOCKER_BUILDKIT=1 docker build -t joewalker/jenkins-slave:jdk8 -f dockerfile/slave/Dockerfile-jdk8 .

jdk11:
	DOCKER_BUILDKIT=1 docker build -t joewalker/jenkins-slave:jdk11 -f dockerfile/slave/Dockerfile-jdk11 .

slave:
	DOCKER_BUILDKIT=1 docker build -t joewalker/jenkins-slave:muti -f dockerfile/slave/Dockerfile .

master:
	DOCKER_BUILDKIT=1 docker build -t joewalker/jenkins-master -f dockerfile/master/Dockerfile .

push:
	docker push -a joewalker/jenkins-slave
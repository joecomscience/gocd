jdk8:
	DOCKER_BUILDKIT=1 docker build -t joewalker/jenkins-slave:jdk8 -f dockerfile/slave/Dockerfile-jdk8 .

jdk7:
	DOCKER_BUILDKIT=1 docker build -t joewalker/jenkins-slave:jdk7 -f dockerfile/slave/Dockerfile-jdk7 .

slave:
	DOCKER_BUILDKIT=1 docker build -t joewalker/jenkins-slave:multi-java -f dockerfile/slave/Dockerfile .

push:
	docker push joewalker/jenkins-slave:multi-java
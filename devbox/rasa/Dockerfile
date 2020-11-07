# Build a devbox for Rasa
FROM ubuntu:20.04

# Never prompt the user for choices on installation/configuration of packages
ENV DEBIAN_FRONTEND=noninteractive
ENV TERM=linux
ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8

RUN apt-get update -y && \
	apt-get install -y emacs \ 
	curl \
	python3-dev python3-pip

RUN pip3 install rasa

# Install code server
RUN curl -fsSL https://code-server.dev/install.sh | sh

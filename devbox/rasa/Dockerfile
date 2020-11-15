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
	git \
	ipython3 \
	python3-dev python3-pip \
	ssh

RUN pip3 install rasa \
	google-cloud-storage


# Install code server
RUN curl -fsSL https://code-server.dev/install.sh | sh

ENV PATH=$PATH:/google-cloud-sdk/bin

RUN wget -q https://dl.google.com/dl/cloudsdk/channels/rapid/google-cloud-sdk.tar.gz && \
    tar xzf google-cloud-sdk.tar.gz -C / && \
    rm google-cloud-sdk.tar.gz && \
    /google-cloud-sdk/install.sh \
    --disable-installation-options \
    --bash-completion=false \
    --path-update=false \
    --usage-reporting=false && \
    gcloud components install alpha beta

RUN gcloud components install anthoscli kpt
FROM ubuntu:wily
MAINTAINER caktux

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
    apt-get upgrade -q -y && \
    apt-get dist-upgrade -q -y && \
    apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 923F6CA9 && \
    echo "deb http://ppa.launchpad.net/elementrem/elementrem/ubuntu wily main" | tee -a /etc/apt/sources.list.d/elementrem.list && \
    apt-get update && \
    apt-get install -q -y gele

EXPOSE 7075
EXPOSE 30707

ENTRYPOINT ["/usr/bin/gele"]

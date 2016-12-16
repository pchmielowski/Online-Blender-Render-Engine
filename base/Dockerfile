FROM golang:1.6

MAINTAINER Piotr Chmielowski

RUN apt-get update
RUN apt-get install -y time blender
RUN apt-get clean
RUN rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

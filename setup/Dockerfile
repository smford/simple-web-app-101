FROM golang:buster

RUN apt-get update -y
RUN apt-get install vim -y
RUN mkdir /go/tutorial
RUN echo "#!/usr/bin/env bash" >> /go/download.sh
RUN echo "git clone https://github.com/smford/simple-web-app-101.git /go/tutorial" >> /go/download.sh
RUN chmod +x /go/download.sh

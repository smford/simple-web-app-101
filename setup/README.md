# Setting up your Computer for the Tutorial

We will be using docker to build and run this tutorial, this will mean you won't have to install much stuff on to your computer and you'll be able to easily remove it in time.

I won't be going in to using docker in much depth, instead I will be giving you the instructions on how to install it on your laptop, and the instructions of how to use docker to run stuff.

Alternatively, you can just install golang and git on to your laptop, but that will require more configuration and it has not been tested thoroughly.

Technologies used in these first two slides:
- docker
- golang
- git

## Install Docker
[Windows](https://docs.docker.com/docker-for-windows/install/)
[OSX](https://docs.docker.com/docker-for-mac/install/)
[Linux](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

## Download the golang image
```
docker pull golang:buster
docker build -t my-golang-app .
```

This will create a container called "my-golang-app" which you will then be able to run by:
```
docker run -it my-golang-app /bin/bash
```


# Setting up your Computer for the Tutorial

We will be using docker to build and run this tutorial, this will mean you won't have to install much stuff on to your computer and you'll be able to easily remove it in time.

I won't be going in to using docker in much depth, instead I will be giving you the instructions on how to install it on your laptop, and the instructions of how to use docker to run stuff.

Alternatively, you can just install golang and git on to your laptop, but that will require more configuration and it has not been tested thoroughly.

Technologies used in these first two slides:
- docker
- golang
- git

## Install Docker
- [Windows](https://docs.docker.com/docker-for-windows/install/)
- [OSX](https://docs.docker.com/docker-for-mac/install/)
- [Linux](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

## Download the golang image and build the container
```
docker pull golang:buster
docker build -t my-golang-app .
```

This will create a container called "my-golang-app" which you will then be able to run by:

## Start the container and get access to the tutorial code
```
mkdir tutorial
docker run -it --mount type=bind,source=$(pwd)/tutorial,target=/go/tutorial my-golang-app /bin/bash
./download.sh
```
- Line 1: makes a directory on your laptop called tutorial, this is where this tutorial will be downloaded to
- Line 2: tells docker:
  - to start your newly built container called my-golang-app with an interative terminal (-it)
  - to mount the newly created tutorial directory into the cointer
  - to start the bash shell
- Line 3: runs the download.sh script which will download this tutorial from github into the tutorial directory

## Things to note
- the container image is based upon ubuntu buster 18.10
- golang, git, vim and other common tools are installed
- to install more stuff use "apt-get install <app to install>"
- the tutorial directory you created is shared with the docker container.  when you create a file in that directory the container will be able to access it, conversely when working in the container you just need to save a file to /go/tutorial and your laptop will be able to see it too

## Cleanup
- To remove the container and delete the image run:
  - docker container delete my-golang-app
  - docker image delete golang:buster


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

## Create the docker configuration file
- create a new empty file called `Dockerfile` (without a .txt or other file extention) and paste the following lines in to it:
```
FROM golang:buster

RUN apt-get update -y
RUN apt-get install vim -y
RUN mkdir /go/tutorial
RUN echo "#!/usr/bin/env bash" >> /go/download.sh
RUN echo "git clone https://github.com/smford/simple-web-app-101.git /go/tutorial" >> /go/download.sh
RUN chmod +x /go/download.sh
```

## Download the golang image and build the container
In your terminal, `cd` to the directory where you created the `Dockerfile` and run the following commands:
```
docker pull golang:buster
docker build -t my-golang-app .
```

This will create a container called "my-golang-app" which you will then be able to run by:

## Start the container and get access to the tutorial code

** If you are using Windows, you will need to grant Docker permission to share a local drive with the container by: right click the whale icon, select settings, then GIVE FULL ACCESS TO YOUR DISK TO DOCKER!

- OSX/Linux:
	```
	mkdir tutorial
	docker run -it --mount type=bind,source=$(pwd)/tutorial,target=/go/tutorial my-golang-app /bin/bash
	./download.sh
	```

- Windows:
	```
	mkdir tutorial
	docker run -it --mount type=bind,source=DRIVELETTER:\tutorial,target=/go/tutorial my-golang-app /bin/bash
	./download.sh
	```

- Line 1: makes a directory on your laptop called tutorial, this is where this tutorial will be downloaded to
- Line 2: tells docker:
  - to `run` your newly built container called `my-golang-app` with an interative terminal (`-it`)
  - to mount the newly created tutorial directory into the container under /go/tutorial `--mount type=bind,source=$(pwd)/tutorial,target=/go/tutorial`
  - to start the bash shell `/bin/bash`
- Line 3: runs the download.sh script which will download this tutorial from github into the tutorial directory you just made

## Important Things to Note
- the container image is based upon ubuntu buster 18.10
- golang, git, vim and other common tools are installed
- to install more stuff use `apt-get install <app to install>`
- the tutorial directory you created is shared with the docker container.  when you create a file in that directory the container will be able to access it, conversely when working in the container you just need to save a file to /go/tutorial and your laptop will be able to see it too
- when you exit the container it will be killed and anything stored in the container will be lost, including anything you installed via `apt-get install`.  To save data it must be placed in the `/go/tutorial` directory

## Starting the container again
If you accidentally exit the container, no problem, you can start it up again by:
- For OSX/Linux: `docker run -it --mount type=bind,source=$(pwd)/tutorial,target=/go/tutorial my-golang-app /bin/bash`
- For Windows: `docker run -it --mount type=bind,source=DRIVELETTER:\tutorial,target=/go/tutorial my-golang-app /bin/bash`

And since you have saved stuff to the /go/tutorial directory it will be saved and available in the newly executed container

## Cleanup
- To remove the container and delete the image run:
  - `docker container rm my-golang-app`
  - `docker image rm golang:buster`
  - `docker image rm my-golang-app`


# A Simple Introduction to Writing a Simple Golang Application using DevOps Principles.

Step by step instructions on creating a simple web app in golang, stored in git, that says "hello world", gets automatically tested and deployed in to the cloud for free, with distributed free logging.

In this introduction I will step you through using some simple and free technologies that you can use to help understand the fundamentals of the process of getting code out to production using a common pipeline, with DevOps based methodologies. I have no particular allegiance to any technology, I use whatever tool works best in a given situation, I selected these technologies because it can all run in the cloud for free, and are relatively modern and popular technologies.


| Technology | Purpose | Why | Link |
|:--|:--|:--|:--|
| golang | programming language | to write the application | [golang.org](https://golang.org/) |
| go dep | programming dependancy management tool | helps manage the dependancies of all the libraries your application will use | [golang/dep](https://github.com/golang/dep) | 
| git | distributed version control system | used to manage versions of the application as we write it | [git.kernel.org](https://git.kernel.org/pub/scm/git/git.git/) | 
| github.com | cloud hosting (storage) for git | a place on the internet to store the code you've just written | [github.com](https://www.github.com/) |
| heroku.com | Cloud hosting of the application | a place to run and host your application | [heroku.com](https://heroku.com/) |
| papertrailapp.com | Cloud hosted log management | a place to see all the logs of your application | [papertrailapp.com](https://www.papertrailapp.com/) |
| travis-ci.org | Test and deploy your application | a system that can test your application for errors, and deploy to heroku if all is well | [travis-ci.org](https://www.travis-ci.org/) |

---

## Preparation for Tutorial
- [Preparation Page](https://github.com/smford/simple-web-app-101/tree/master/setup/prep)

## Things we'll do together in the Tutorial
1. [Setup the tutorial on your laptop](https://github.com/smford/simple-web-app-101/tree/master/setup)

## Slide Overview
1. [Slide 00: DevOps: Background to Modern Practices](https://github.com/smford/simple-web-app-101/tree/master/slide-00)
1. [Slide 01: Structure of a golang application](https://github.com/smford/simple-web-app-101/tree/master/slide-01)
1. [Slide 02: Changing the simple golang application](https://github.com/smford/simple-web-app-101/tree/master/slide-02)
1. [Slide 03: Making the appication run on any operating system](https://github.com/smford/simple-web-app-101/tree/master/slide-03)
1. [Slide 04: Making a simple web application](https://github.com/smford/simple-web-app-101/tree/master/slide-04)

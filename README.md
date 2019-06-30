# A Simple Introduction to Writing a Simple Golang Application usin DevOps principles.

Step by step instructions on creating a simple web app in golang, stored in git, that says "hello world", gets automatically tested and deployed in to the cloud for free, with distributed free logging.

Let us assume that code can be written, the next step/problem is to ensure that code is of high quality, can actually run somewhere, be reliable and resilient,  whilst also being scalable.  In the DevOps world this is called Building, Testing and Deploying, and the flow of code through the each of these steps is called the Pipeline.

Luckily most of the Build, Test and Deploy problem can be solved using cloud technologies and gluing them together.  The problem can be answered in many other different ways, with many different types of technologies, not all of which need to be cloud based.

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


## Environments

In the IT world we often talk about environments, it is best explained in an example:

- Sandbox Environment: a dirty set of servers and technology that I can play around with and test my application on without ruining other peoples work or disrupting any real users
- Pre-production Environment: a set of servers and technologies that are identical to the production environment, to firstly test the application against other components of the system; and secondly test performance of the application on infrastructure which represents what the end user will use
- Production Environment: the actual set of servers and technology that will be used by end users of the application

The names and even the number of the environments do not really matter, different employers work differently.  What is important to understand is that as the environments gets closer to production three things should occur:

1. the cost of running the environments increases
2. confidence that the code will run well in production increases
3. that disruption to any environment effects more and more users


## Pipeline
Once code is written, it needs to be built, tested and deployed.

### 1. Build
This can mean compiling the code in to a binary executable file, or another type of "build artifact", most of the time we will add a version to it.  Other types of build artifacts include: docker images, rpms, debs, virtual machine images, tarballs, pretty much anything


### 2. Test
We need to test that code is of good quality before it gets deployed to production.  Testing can include many different aspects, from simulating user journeys through the web application; down to testing the actual syntax of the code.  We should aspire to make the tests as automated as possible, removing the need for a human to do anything.  Writing the automated tests can sometimes be even more challenging than writing the applications code in the first place. 

Testing does not just include testing the application, it can involve testing parts of the application; testing the application with other components (like a database backend); and testing the application against the same infrastructure that would be used in the production environment.

As code passes a set of tests, it can then be tested on the next environment closer to production.  
 
### 3. Deploy



---

## Lets Get Started 
1. create simple application in golang
1. convert it to a webapp
1. allow for some simple configuration to be passed to the app
1. save this in to github
1. configure testing of your app
1. get your app deployed in to the cloud
1. configure logging of your app
 

## Alternative Technologies



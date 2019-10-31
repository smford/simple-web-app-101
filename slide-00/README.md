# DevOps: Background to Modern Practices

Here I will be talking about some of the basic principles of DevOps

- General Principles of DevOps
- Agile Working
- Environments
- Pipelines

## General Principles of DevOps

There is no clear definition of what DevOps is, it is flexible and everyone has their own opinion of it.  A good generic description is:

DevOps is a set of practices that combines software development (Dev) and information-technology operations (Ops) which aims to shorten the systems development life cycle and provide continuous delivery with high software quality. [source](https://en.wikipedia.org/wiki/DevOps)

My definition: It is a set of techniques and methods that try and bring developers and infrastructure engineers closer together.  We leverage technology, ways of working, and foster collaborative behaviours to allow these two types of workers to work more efficiently.

Who cares?  Well, it means as technical people we get to play with new and exciting technogies faily regularly.  As an Infrastructure person I get to meddle in code more, and developers get to learn infrastructure more.  DevOps creates a better understanding of each sides needs and desires, and that allows better work to be delivered.

The primary output in the DevOps world is code.  Code makes applications, applications give competive advantage, and that means profit.

DevOps:
- ensures code is of high quality (less bugs mean happier customers)
- can actually run somewhere
- be reliable (high uptime) and resilient (so it can heal itself)
- whilst also being scalable (grow when needed)
- and work as autonomously as possible


## Agile Working

One way of working that most teams use is called *Agile*.  This essentially involves:
- chopping up work in to small tasks
- working through these small tasks and regularly checking whether what we are doing is what we should be doing; and quickly pivoting when the direction changes
- failing fast: it's a fact that not everything succeeds, so it's best to monitor progress and if it isn't working, stop and re-evaluate, and try something new early before things get too far off track
- having regular meetings to monitor work progression and priority.  This is like keeping your finger on the pulse
- check in on the teams health regularly.  Ask them what they aren't happy about and try to address them quickly, having someone empowered and responsible for addressing problems means the techies can focus on technical stuff, whilst someone else keeps the distractions away
- rather than trying to complete a task/problem perfectly on the first attempt, which almost always fails, accept that you might need to come back to a problem a number of times and polish it before you get it right.  This is called iterative development, you create a simple application (called minimum viable product), then present it back to the end user, making sure it's what they want, then add a little more, and present back to the user, and repeat over and over



## Environments

In the IT world we often talk about environments, it is best explained in an example:

- Sandbox Environment: a dirty set of servers and technology that I can play around with and test my application on without ruining other peoples work or disrupting any real users
- Pre-production Environment: a set of servers and technologies that are identical to the production environment, to firstly test the application against other components of the system; and secondly test performance of the application on infrastructure which represents what the end user will use
- Production Environment: the actual set of servers and technology that will be used by end users of the application

The names and even the number of the environments do not really matter, different organisations work differently.  What is important to understand is that as the environments gets closer to production three things should occur:

1. the cost of running the environments increases
2. confidence that the code will run well in production increases
3. that disruption to any environment effects more and more users

## Pipeline

Once code is written, it needs to be built, tested and deployed; the flow of code through the each of these steps is called the Pipeline.

Luckily most of the Build, Test and Deploy pipeline can use cloud technologies and DevOps involves us gluing them together.

### 1. Build
This can mean compiling the code in to a binary executable file, or another type of "build artifact", most of the time we will add a version to it.  Other types of build artifacts include: docker images, rpms, debs, virtual machine images, tarballs, pretty much anything


### 2. Test
We need to test that code is of good quality before it gets deployed to production.  Testing can include many different aspects, from simulating user journeys through the web application; down to testing the actual syntax of the code.  We should aspire to make the tests as automated as possible, removing the need for a human to do anything.  Writing the automated tests can sometimes be even more challenging than writing the applications code in the first place.

Testing does not just include testing the application, it can involve testing parts of the application; testing the application with other components (like a database backend); and testing the application against the same infrastructure that would be used in the production environment.

As code passes a set of tests, it can then be tested on the next environment closer to production.

### 3. Deploy

Deploying a build artifact (an application) can happen many ways depending on the technology used to run the application, and the technologies used in an organisation.  Many places use specialised tools such as Jenkins, Travis, Concourse or CircleCI to deploy the application.  But you can use scripts or orchestration tools (like ansible, salt) to deploy also.

Deployment is often to a sandbox environment first, then the tests run (either by hand or ideally automated), then if the tests passed succesfully deployed to a pre-production environment where a further set of tests are run.  These second set of tests are typically used to test two things - firstly the performance of an application when used on production like infrastructure, and secondly against the complete set of technologies that the application relies upon.  Furthermore, the performance tests will often simulate heavy load on the application to discern wether any performance related issues arise (in the application or in the infrastructure).  If it is a web application, simulations of users navigating the website and the application will occur.

Once this round of testing is complete, the deployment system will either automatically deploy to production, or wait until sign off by a quality assurance person before deploying to production.

We aspire to make deployments have minimal effect on users, and we aspire to have zero down time deployments.  There are various techniques and technologies that allow us to achieve it.  For example, some deployments will create a brand new production environment, then instruct a load balancer to migrate users across on to the new production environment.  Once the old environment has drained of users, it will shut the old production environment down.  The methods and technologies used to do these sorts of deployments depend heavily on the application being created in such a way to allow the deployments to succeed.  The most popular approach is called [12 Factor](https://12factor.net/).

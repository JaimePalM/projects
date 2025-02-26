# Introduction

Jenkins is an open-source automation server that can be used to automate all sorts of tasks related to building, testing, and delivering or deploying software. Jenkins can be installed through native system packages, Docker, or even run standalone by any machine with a Java Runtime Environment (JRE) installed.

High level learning objectives:
- CI/CD at a Glance (Continuous Integration/Continuous Deployment)
- Jenkins: an Overview
- Introduction to Jenkins DSLs (Domain Specific Languages)
- Configure and Implement Jenkins DSLs
- Jenkins Job Structure

The Jenkins server checks the repository at regular intervals for changes. If changes are detected, Jenkins will automatically download the changes, build the project, run tests, and generate reports. If the build is successful, Jenkins will deploy the changes to the production server. This way developers can focus on writing code and Jenkins will take care of the rest.

Jenkins has a community ecosystem of plugins that can be used to extend the functionality of Jenkins. Jenkins can be integrated with other tools like GitHub, Docker, Slack, and more.

Who uses Jenkins?
- Developers
- DevOps Engineers
- System Administrators

# CI/CD 

CI/CD stands for Continuous Integration/Continuous Deployment. CI/CD is a set of practices that automate the process of integrating code changes from multiple contributors into a shared repository several times a day, testing the changes (CI), and deploying the changes to production (CD). CI/CD helps to automate the software delivery process, reduce manual errors, and increase the speed of software delivery.

CI Example:
During the course of regular software development, one or more team members check in their code regularly to a shared repository (using Git for example). The Ci server then fetches these fresh changes and integrates them and then runs the build process and produces a bundle of artifacts. These artifacts are then versioned and stored in a repository. The CI server then runs several automated tests (unit tests, integration tests, etc.) on the artifacts. Then it notifies the team of the results. 

CI
- Code integration practice
- Centrally managed code repository
- Efficient build and test process
- Immediate issue detection and resolution
- Successful delivery pipeline

CD Example:
Once the CI process is successful, the CD process takes over. The CD process automates the deployment of the artifacts to the production environment. The CD process can be manual or automated. The CD process can be triggered by the CI process or by a manual trigger.

CD
- Sustainable product delivery
- On-Demand deployment
- Cost efficient process
- Continuous testing and feedback
- High quality deliverables

# Jenkins 

## Introduction

You begin by logging into the Jenkins server. 
1. Initialize a **new job** on the Jenkins dashboard.
2. Specify the **source code repository** from you can pull the code for automation such as GitHub. 
3. Specify **triggers** for the job (start to build once another build is complete, kick off at intervals, etc.). 
4. Specify the **build process** (build, test, deploy). 
5. Specify the **post-build** actions (send email, deploy to production, etc.).

Jenkins Master: the Jenkins server that manages the jobs and the build process. Its job is to pull the SCM repository, schedule build jobs, dispatch builds to the Jenkins slave, monitor the build process and present the build results.

Jenkins Slave: a Jenkins slave is a Java executable that runs in a remote machine that is connected to the Jenkins master. The Jenkins slave is used to run the build process received from the Jenkins master.

Jenkins Jobs:
- Using Jenkins UI
- Can Copy/Clone Jobs 
- Cumbersome & Error prone
- Jenkins DSL (configure as code, easier to manage)

## Installation

Pre-requisites:
- Java Runtime Environment (JRE) 1.8 or higher
- Oracle Java Development Kit (JDK) 1.8 or higher

Installation Steps:
- Download Jenkins
- Service Logon Credentials: "Run Service as Local System" 
- Open Jenkins in Browser: http://localhost:8080
- Unlock Jenkins: initial password is stored in the Jenkins log file (follow the instructions)

## Configuration
Now we've Jenkins installed and unlocked, we can configure Jenkins. 

Plugins: Jenkins has a community ecosystem of plugins that can be used to extend the functionality of Jenkins. Jenkins can be integrated with other tools like GitHub, Docker, Slack, and more.

Steps:
- Install suggested plugins (wait for installation to complete)
- Create first admin user (all can be invited)
- Leave the Jenkins URL as default
- Save and Finish

Installing a plugin:
- Go to "Manage Jenkins"
- Go to "Manage Plugins"
- Go to "Available" tab
- Search for "Job DSL" and install it
Jenkins DSLs (Domain Specific Languages) are used to define Jenkins jobs as code.

There is a lot more options to configure Jenkins, like setting up security, configuring Jenkins nodes, etc.

## Building a Job
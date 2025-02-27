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

## Jenkins DSL

It's a plugin which enables developers to define jobs using conditional states, variables and loops. It's a way to define Jenkins jobs as code.
- Variables: defined using `def` keyword.
- Loops: defined using `for` keyword, e.g. number of times a job is created.
- Branching: defined using `if` keyword, e.g. if a job is created based on a condition.
- Functions: defined using `def` keyword, e.g. a function to create a job.

Jenkins offers an API to create jobs. E.g. `job('job-name')` creates a job with the specified name (the name is the only required parameter). Jobs definition follow a XML format.

### Playground

It's a web application that allows you to write and test Jenkins DSL scripts. It's a great way to learn Jenkins DSL. You'll write the job in the left pane and see the XML representation of the job in the right pane.

### Creating a Job 

Jenkins Job: a Jenkins job is a set of instructions that tells Jenkins what to do when it receives code changes from the source code repository. A Jenkins job can be created using the Jenkins UI or using Jenkins DSL.

Creating a Job using Jenkins UI:
- Click on "New Item"
- Enter a name for the job
- Select "Freestyle project" (general purpose job)
- Configure the job
  - Source Code Management: specify the source code repository (e.g. GitHub)
  - Triggers: specify the triggers for the job (e.g. build periodically, build when a change is pushed to the repository)
  - Build Steps: specify the build process (e.g. build, test, deploy). We choose "Process Job DSLs" and specify the DSL script.
    - We're going to pass the DSL script, so select "Use the provided DSL script"
- Save the job

DSL script example:
```groovy
def jobName = 'my-first-job'

job(jobName) { }
```
This script creates a job with the name `my-first-job`.

Now, we can build the new job, pressing "Build Now" from the project page. The job will be scheduled and eventually will be successfully built. Note: it's possible we need to disable the script security in Jenkins.

Now we can go the main page and see the new job. Click on it and see the job details and it was sent by the project we created.

## Jenkins Roles and Configuration

From "Manage Jenkins" we can go to "Users" and see the users that have access to Jenkins. There it can be created new users, delete users, change passwords, etc.

Once a user log in, they can configure their profile. There are several options like the API token (useful when using Jenkins API), SSH public keys, etc.

## Security

By default Jenkins give to every authenticated user the "Overall" role. This role gives the user the ability to read and configure Jenkins. This isn't the ideal scenario, so it's recommended to create new roles and assign them to non admin users. For that we can download the "Role-based Authorization Strategy" plugin.
Now we go to "Security" -> "Authorization" and select "Role-Based Strategy".

There are th  ree types of roles:
- Global Roles: apply to all Jenkins objects (users, jobs, nodes, etc.)
- Project/Item Roles: apply to specific jobs (pattern: match job name, user with the role can access the items with name that match the pattern)
- Slave/Agent Roles: apply to specific nodes

Now we can go to "Assign Roles" and assign roles to users. 

Finally, create two projects (DevProject1 and TestProject1) and assign roles to users. Now users with developer role can only access DevProject1, because the name matches the "Dev" pattern. Same for the tester role. If we log out and log in with a user we won't see the other project.

## Basic Configuration

In system configuration we can set several options:

**System Message**: we can set a message that will be displayed to all users when they log in. Go to "Manage Jenkins" -> "Configure System" and set the message. It follows HTML format (changing it from security settings).

**\# of executors**: it's the number of concurrent jobs that can be run on the Jenkins master. It's recommended to set it to the number of cores in the machine.

**Labels**: in Jenkins we have a system for distributing jobs to different nodes. We can assign labels to nodes and jobs. Labels are used to group nodes and jobs and they are used to restrict jobs to run on specific nodes. When a node is created, it's assigned a label and then it can be set to only build jobs with label match. On the other hand, we can go the job configuration and set the "Restrict where this project can be run" option to a label, this way the job will only run on nodes with that label expression.

**Quiet Period**: it's the time Jenkins waits before starting a build after a change is detected. It's useful when there are multiple changes in a short period of time. It's expressed in seconds. In jobs there is a build trigger called "Poll SCM" that will check the repository for changes and if there are changes it will wait the quiet period before starting the build.

**SCM Checkout Retry Count**: it's the number of times Jenkins will retry to checkout the code from the repository. It's useful when the repository is down or there are network issues.

**Restrict Project Naming**: it's a regex pattern that restricts the project names. It's useful to enforce a naming convention.

**Global Properties**: we can set global properties that can be used in jobs. We can set environment variables (key-value, used later with $key), etc.

**Jenkins Location**: we can set the Jenkins URL and the system admin e-mail address.

**Shell**: we can set the shell that Jenkins will use to run shell scripts. It can be set to the default shell or to a custom shell.

## Jobs

Checklist:
- How to create a basic job [already covered](#creating-a-job)
- Basic job configuration 
- How to run a job remotely
- How to chain a job execution

There are different "**Build Triggers**" that can be used to start a job:
- Build periodically: it's used to schedule a job to run at specific intervals. It uses the cron syntax (e.g. `H/5 * * * *` runs the job every 5 minutes and the `H` means a random minute to avoid all jobs running at the same time).
- Poll SCM: it's used to check the repository for changes and if there are changes it will start the build. It's useful when the repository doesn't have a webhook to notify Jenkins of changes.

Regarding "**Build Steps**" that can be used to define the build process. Most common is to use the "Execute shell" step to run shell commands. We can also use the "Invoke top-level Maven targets" step to run Maven commands.

There are different "**Post-build Actions**" that can be used to define what to do after the build is complete. For example, we can use the "Email Notification" action to send an e-mail after the build is complete.

### Dashboard fields

S (State): it's the current state of the job. It can be:
- Green: the job is successful
- Blue: the job hasn't run yet
- Red: the job failed
- Yellow: the job is unstable
- Grey: the job is disabled

W (Weather): it's a weather icon that shows the health of the job. In other words, it shows how often the job fails. The more it fails, the worse the weather icon.

### Triggering a Job Remotely

We can trigger a job remotely using the Jenkins API. We go to the job configuration and select "Trigger builds remotely". We can set an authentication token that will be used to authenticate the request. Then with the URL `http://localhost:8080/job/<job-name>/build?token=<token-name>` we can trigger the job.

Now we can use the URL to trigger the job from a script, for example from a website.

### Chaining Job Execution

A job execution can be chained by using the "Build other projects" post-build action. We can set the projects to build and the conditions to build them. For example, declare a job to build another job only if the current job is successful.

There, it requires a list of projects to watch and a condition (if the project is stable, unstable, fails or always).

Example: we have 3 jobs (Job1, Job2, Job3) and we want to chain them. We go to Job2 configuration and set the "Build Trigger" to "Build after other projects are built" and set the projects to watch to Job1 and the condition to "Successful". Then in the "Post-build Actions" of Job2 we set the "Build other projects" action to build Job3. So configuring one job we chain all the three jobs.

## Integrating Jenkins with GitHub

Learning objectives:
- Create a Java program and run it through CLI
- Create a Jenkins job to run the Java program
- Add the program to a GitHub repository
- Configure Jenkins to pull the program from GitHub when a change is made

### Java Program

We create a simple Java program that prints "Hello, Jenkins!" to the console. We create a new directory and create a new file called `HelloJenkins.java` with the following content:
```java
public class HelloJenkins {
    public static void main(String[] args) {
        System.out.println("Hello, Jenkins!");
    }
}
```
Then we compile the program using the command `javac HelloJenkins.java` and run it using the command `java HelloJenkins`.

## Automated Deployment

Software deployment has 4 main stages:
- Build: compile the code and generate artifacts
- Deploy: deploy the artifacts to the server (test environment)
- Test: run tests on the deployed artifacts
- Release: deploy the artifacts to the production server

For this we'll need the "Deploy to container" plugin

### Deploy to Container Plugin

Create a build job. 
  1. Download a sample war (this is a Java web application) file from the tomcat website.
  2. Add any build step (e.g. execute shell) and as post build action select "Deploy war/ear to a container".
  3. Add the war file with `**/*.war` and as context path `sample.war`.
  4. Select Tomcat 9.x for the container. 
     1. Download and Install Tomcat with an admin user.
     2. Put the credentials in the Jenkins configuration.
  5. Save and build the job.

If we go to the Tomcat server we can see the application deployed (http://localhost:8090/sample.war).

### Notifications

We can configure Jenkins to send e-mail notifications when a build fails. We go to "Manage Jenkins" -> "Configure System" and set the SMTP server. Then we go to the job configuration and set the "Post-build Actions" to "E-mail Notification". 

There it requires a SMTP server (e.g. smpt.gmail.com), the sender e-mail address, its password (use a token) and check SSL.

Now in the post-build actions we can set the recipients e-mail addresses and the conditions to send the e-mail (e.g. always, unstable, etc.). Therefore, there are some plugins to enhance the e-mail notifications, like Extreme Notification or Email extension plugins.

### Pipelines

They are a way to define the build process as code. Pipeline is a workflow with a group of events or jobs that are chained and integrated with each other in sequence. Pipelines can be defined using a Jenkinsfile. A Jenkinsfile is a text file written in Groovy that contains the definition of the build process.

Delivery Pipeline: we'are going to simulate a real case so create three simple jobs (Build, Test and Deploy). Now we can chain the jobs using the build trigger "Build after other projects are built". Then install the "Delivery Pipeline Plugin" and create a new view with the three jobs. Now we can create a new view with the plugin. At the end, we'll configure the components of the pipeline (initial job).

There are some interesting options like "Enable start of new pipeline build" that will trigger the pipeline directly from the view and see the progress of the pipeline. We can also set the "Enable rebuild" option to rebuild an specific job in the pipeline (without rebuilding the whole pipeline).

Build Pipeline: it's a way to visualize the build process. We can install the "Build Pipeline Plugin" and create a new view with the jobs. We can drag and drop the jobs to create the pipeline. 

The difference between the two plugins is that the Delivery Pipeline Plugin is more focused on showing the progress of the pipeline and the Build Pipeline Plugin is more focused on managing the pipeline (it's more interactive).

## Tips & Tricks
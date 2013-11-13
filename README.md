# sputnik environments
This repository contains the proposal for `sputnik-environments`.

## What are `environments`
In a nutshell, environments are a collection of "things" needed to develop software. They exist to help provide a uniform development environment for a user.

Environments consist of the following components:

- languages: the programming language used in the project.
- frameworks: Language specific software layered on top
- people: Defines customer setup for particular usernames (optional)

Sputnik environments are about optimizing ramp-up time to new developers. An open-source project might publish a sputnik environment for contributors. A company might define an environment to help shorten the onboarding of new team members.

## Projects
These are wrapped into a `project`. A project defines which language the developer is using and any additional resources. These definitions are a very basic yaml file with very few required fields. The general rules proposed are:

- a project MUST have a primary language
- a project CAN have secondary languages
- a project CAN have frameworks
- a project CAN have external resources

## External resources
Frequently when developing more than the most naive of software projects, you need access to external resources. In the past we might spin up a virtual machine but even with current tooling this is a time consuming AND resource intensive process. The alternatives are to install these resources locally or to utilize a shared resource. None of these are ideal.

Projects in sputnik can define a list of `containers` they would like to have as part of the project. This is done by leveraging linux containers via the `docker` project. Utilizing the docker.io public registry, a user can pull in any external resources they need such as a database, a memcached server or whatever.

# Opinionated minimalism and composition
In each subdirectory there is a README describing the for of the components. You will note that the number of required fields is fairly minimal. This is intentional. A user's development environment is very personal. Dictating that a developer would be required to use a specific editor or directory structure would be insane.

sputnik environments are only opinionated where neccessary. At a minimum you need a project and a language. Your language defintion doesn't even have to do anything. Your project can consist of nothing but installing containers. The flexibility exists for you, the end user, to define what commands are run at what stage to provide a consistent development environment. Your project could call `chef` or `puppet` or just make a directory structure for you and install some base packages.

The proposed structure is in place to help scope feedback from the community before tooling is created to make it happen.

# Feedback and best practices
The purpose of this repository is to solicit feedback from development communities about "best practices" for developers of that community. As we all know best practices are sometimes neither best or practiced. Our hope is that before the solicitation deadline ends, users of the appropriate communities will contribute to defining the out of the box experience with Sputnik as a developer

# How to contribute
There are many ways to contribute:

- Open issues on this repository with suggestions on proposed definition syntax
- Fork the repository and propose your own definition for a language
- Share this initiative with your respective communities
- Offer insight into how your community or company can use this
- Be honest and tell us what works and what doesn't

While almost everything is open to modification, we have a few small requirements.

- The definition format MUST be yaml. YAML was chosen because it allows the appropriate data structures, is both human and machine readable and allows comments.
- Respect the goals of minimalism and composition defined above
- This is based on Ubuntu 12.04 amd64 (for now) as the primary use case is for inclusion in the Sputnik laptop.

# What's next
At the end of the solictation period, we will curate the best proposals from each community along with any feedback and start developing the first version of the tool.

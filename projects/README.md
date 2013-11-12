# Projects

The idea with projects is to define a development project. A project consists of one or more languages and one or more frameworks.

Using the ficticious `myproject` which is a django app, we would need:

- python as the language of choice
- django as the framework

## Fields
The following fields are proposed.

### name (required)
The name of the project

### version (required)
The version of the project

### language (optional)
The primary language of the project

### additional\_languages (optional)
Any secondary languages that need to be installed.

### base\_directory (optional. default: $HOME)
The base directory for the project on the local filesystem

### init (optional)
A collection of commands to run to initialize this project. Called after any languages and frameworks are installed

### containers (optional)
A list of named containers from the docker registry. These are useful for things like a PGSQL server or a Memcached instance so that you don't have to litter your local system.

Note that containers are not stopped when switching projects.

### start (optional)
A collection of ordered commands to call to "start" this project. This is analagous to the `activate` functionality in virtualenv.

## Format
The proposed format for a language definition is the following yaml:

```yaml
name: myproject
version: 1.0
language: python
additional_languages:
- javascript
framework:
- django
base_directory: "/home/me/development/"
init:
- "pip freeze > requirements.txt"
containers:
- jbarbier/memcached
- jpetazzo/pgsql
start:
- "docker run -p 5432 jpetazzo/pgsql /init YourSecretPassword"
- "docker run -p 11211 jbarbier/memcached memcached -u daemon"
```

In this case we would install the listed packages via apt. When a `python` project is created, the project name will be passed to the listed command.

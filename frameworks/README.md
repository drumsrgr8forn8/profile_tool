# Frameworks

The idea with frameworks is to define any language-level packages needed as part of a project. Note that frameworks are not required and are mainly a convienince. However they can help shorten the project definition requirements.

For instance if you are developing a django application, you might want the following additional python packages:

- django
- django-jinja2

We would then leverage the language's native packaging tool so we can install those packages. In the case of python, we would switch into our virtualenv for the project first before installing the packages

## Fields
The following fields are proposed.

### name (required)
The name of the framework.

### version (required)
The version of the framework. Note that this does not correspond to any packages

### language (optional)
The language for use with the framework. If specified, the named language definition will be used

### modules (optional)
A collection of language-specific packages to be installed with the language's native utility (e.g. gem, pip)

### create (optional)
A command to be run after installing the framework requirements. Will be passed the project name.

## Format
The proposed format for a framework definition is the following yaml:

```yaml
name: django
version: 1.5.4
language: python
modules:
- django
- django-jinja2
create:
- "django-admin.py startproject"
```

In this example we would first use the `python` language definition to create our base project environment. This would provide us with a virtualenv named after our project. We would then switch to the virtualenv and install `django` and `django-jinja2` via `pip` inside the virtualenv. Afterwards we would run, `django-admin.py startproject` passing it the name of our project.

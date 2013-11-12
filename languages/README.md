# Languages

The idea with languages is to define any base packages needed to install a new "language".

For instance if you are a python developer, you might want the following:

- python
- pip
- virtualenv

This would be the bare minimum. Think of languages like "toolchains".

## Fields
The following fields are proposed.

### name
The name of the language

### version
The version of the language

### packages
Any packages in the default ubuntu 12.04 repository that are needed.

### prep
Any commands that need to be run to setup a project based on this language.

## Format
The proposed format for a language definition is the following yaml:

```yaml
name: python
version: 2.7.4
packages:
- python
- python-dev
- virtualenv
- python-pip
prep:
- "virtualenv --no-site-packages --distribute"
```

In this case we would install the listed packages via apt. When a `python` project is created, the project name will be passed to the listed command.

# Additional Note
Languages are totally arbitrary. They don't have to actually map to a REAL programming language. You can define `mylang` as simply requiring X system packages and then move additional logic into your project. 

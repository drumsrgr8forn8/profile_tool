# POC
This is a quick POC to muck about with the CLI api. You need go 1.1+ installed.

## Building
```
export GOPATH=`pwd`
make clean all
```

## Example usage
Note that this doesn't actually do what it's saying JUST yet. This is an mvp for how the utility MIGHT work:

### Installing a language
```
>> bin/profile_tool language install javascript
2013/11/12 02:01:56 Setting up language: javascript 0.10.21
2013/11/12 02:01:56 Running: curl -O http://nodejs.org/dist/v0.10.21/node-v0.10.21-linux-x64.tar.gz
2013/11/12 02:01:56 Running: tar zxvf node-v0.10.21-linux-x64.tar.gz
```

```
>> bin/profile_tool language install python
2013/11/12 02:02:59 Setting up language: python 2.7.4
2013/11/12 02:02:59 Installing package: python
2013/11/12 02:02:59 Installing package: python-dev
2013/11/12 02:02:59 Installing package: virtualenv
2013/11/12 02:02:59 Installing package: python-pip
2013/11/12 02:02:59 Running: virtualenv --no-site-packages --distribute
```

### Installing a framwork
Note that frameworks pull in any defined language yaml files first:

```
>> bin/profile_tool framework install django
013/11/12 02:03:26 Setting up framework: django 1.5.4
2013/11/12 02:03:26 Setting up language: python 2.7.4
2013/11/12 02:03:26 Installing package: python
2013/11/12 02:03:26 Installing package: python-dev
2013/11/12 02:03:26 Installing package: virtualenv
2013/11/12 02:03:26 Installing package: python-pip
2013/11/12 02:03:26 Running: virtualenv --no-site-packages --distribute
2013/11/12 02:03:26 Installing module: django
2013/11/12 02:03:26 Installing module: django-jinja2
2013/11/12 02:03:26 Running: django-admin.py startproject
```

### Starting a defined project
As with frameworks, the languages and frameworks are pulled in before the final steps:

```
>> bin/profile_tool project start myproject                                                                1 ↵
2013/11/12 02:05:01 Setting up project: myproject 1.0
2013/11/12 02:05:01 Setting up language: python 2.7.4
2013/11/12 02:05:01 Installing package: python
2013/11/12 02:05:01 Installing package: python-dev
2013/11/12 02:05:01 Installing package: virtualenv
2013/11/12 02:05:01 Installing package: python-pip
2013/11/12 02:05:01 Running: virtualenv --no-site-packages --distribute
2013/11/12 02:05:01 Setting up language: javascript 0.10.21
2013/11/12 02:05:01 Running: curl -O http://nodejs.org/dist/v0.10.21/node-v0.10.21-linux-x64.tar.gz
2013/11/12 02:05:01 Running: tar zxvf node-v0.10.21-linux-x64.tar.gz
2013/11/12 02:05:01 Setting up framework: django 1.5.4
2013/11/12 02:05:01 Setting up language: python 2.7.4
2013/11/12 02:05:01 Installing package: python
2013/11/12 02:05:01 Installing package: python-dev
2013/11/12 02:05:01 Installing package: virtualenv
2013/11/12 02:05:01 Installing package: python-pip
2013/11/12 02:05:01 Running: virtualenv --no-site-packages --distribute
2013/11/12 02:05:01 Installing module: django
2013/11/12 02:05:01 Installing module: django-jinja2
2013/11/12 02:05:01 Running: django-admin.py startproject
```

Obviously there are some gaps here. This would rely on the underlying step to be idempotent. In the case of package managers, you would hope that to be true. If you're curl-bashing a file then likely not.

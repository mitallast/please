Please
======

[![Build Status](https://travis-ci.org/mitallast/please.svg)](https://travis-ci.org/mitallast/please)

Simple console utility to combine miltiple package managers at one.

See more at https://mitallast.github.io/please/

Supported package managers
==========================

- `homebrew` - for os x and linux
- `apt` - for apt based linux like debuan, ubuntu 
- `yum` - for yum based linux like centos, fedora
- `npm` - cross platform node.js package manager

How to build
============

```sh
go build
```

Example of usage
================

```sh
$ please install python
[brew install python]
Warning: python-2.7.10_2 already installed
```

How to contribute
=================

Before commit changes, use standard tool go fmt.

Before push changes to master branch, use rebase to avoid merges!
example:

```sh
git co master
git pull
git co feature-branch-name
git rebase master
....
git co master
git merge feature-branch-name
git push
```

Credits
=======

- Alexey Korchevsky @mitallast
- Alexey Tabakman @samosad

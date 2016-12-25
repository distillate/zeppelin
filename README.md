# Zeppelin

[![Build Status](https://travis-ci.org/distillate/zeppelin.svg?branch=master)](https://travis-ci.org/distillate/zeppelin)

This repository contains Zeppelin runtime.

## Installing dependencies
We use [govendor](https://github.com/kardianos/govendor) for dependencies management. To install dependencies type:
```
$ govendor sync
```

## Building
To build API binary type:
```
$ go build -o api cmd/api/*
```
To build worker binary type:
```
$ go build -o worker cmd/worker/*
```

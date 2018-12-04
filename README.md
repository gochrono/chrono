![Alt chrono](./.github/full-logo-with-tagline.png)

[![GitHub (pre-)release](https://img.shields.io/github/release/gochrono/chrono/all.svg)](https://github.com/gochrono/chrono/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/gochrono/chrono)](https://goreportcard.com/report/github.com/gochrono/chrono)
[![Build Status](https://travis-ci.org/gochrono/chrono.svg?branch=master)](https://travis-ci.org/gochrono/chrono)
![GitHub](https://img.shields.io/github/license/gochrono/chrono.svg)



## Overview ##

Chrono is a time tracking tool written in Go.
It is fast and simple to use.

Want to know what you did with your time? Chrono will help you with that. Track how
long you spend on a project. Add notes so you know exactly what
you did.

Chrono can manage thousands of frames in less than a second.

To begin using Chrono, you can start tracking a project with `chrono start [project name] [tags]`

To stop tracking the project use `chrono stop`.

For more on using Chrono, check out the [quick start guide][1]

#### Supported Architectures ####

We provide pre-built Chrono binaries for Windows, Linux, and macOS (Darwin) for x64, i386 and ARM architectures.

Alternatively, Chrono can be compiled from source where ever the Go compiler tool chain can run.

**For more information on which architectures you can install Chrono on, check out the [Go documentation](https://golang.org/doc/install).**

# Quickstart #

## Installation ##

The simplest way to install Chrono is to download the latest binary from the [releases page](https://github.com/gochrono/chrono/releases).

The binaries have no external dependencies. After downloading it, place it somewhere on the PATH (such as `/usr/local/bin` on Linux).

Alternatively you can install Chrono by building it yourself. This ensures you're running the absolute bleeding edge version.

## Usage ##

To start tracking time a project, use the `start` command:

``` bash
$ chrono start development +chrono
```

This creates a new __frame__ for the development project with the chrono tag.

Keep notes of what you do for a project with the `notes add` command:

``` bash
$ chrono notes add "made some awesome changes to the README"
$ chrono notes show
[0]: made some awesome changes to the README
```

The notes are added to the current __frame__.

Get information about the current frame with the `status` command:

``` bash
$ chrono status
Project development [chrono] started 10 seconds ago.
```
To stop tracking time for the current frame, use the `stop` command:

``` bash
$ chrono stop
Stopping project development [chrono], started 5 minutes ago (id: 073bbf).
```

You can show a chronolical list of the current day's session (or __frames__) through the `log` command:

``` bash
$ chrono log
Monday  3 December 2018
    (ID: 0d3131) 10:15 to 10:20     0h 05m 00s  development [chrono]
```

For a list of all available commands, use the `help` command:

```
$ chrono help
```

For a list of all available options and arguments for a command, use the `---help` flag:

```
$ chrono log --help
```

## Building Chrono from the Source ##

### Prequisite Tools ###

* [Git](https://git-scm.com/)
* [Go (at least Go 1.11)](https://golang.org/dl/)


### Downloading the source ###

Chrono uses [Go Modules](https://github.com/golang/go/wiki/Modules) to handle dependencies.

The easiest to get the source is to clone Chrono in a directory outside of `GOPATH`, for example:

``` bash
mkdir ~/src && cd ~/src
git clone https://github.com/gochrono/chrono.git
cd chrono
go install
```

## Contributing to Chrono ###

To contribute to the Chrono project or documentation, you should [fork the GitHub project](https://github.com/gochrono/chrono#fork-destination-box) and clone it to your machine.

For a complete guide to contributing to Chrono, see the [Contribution Guide](CONTRIBUTING.md).

We welcome contributions of many kinds from updating documentation, feature requests, bug reports & issues,
feature implementation, pull requests, answering other users questions, etc.

### Asking Support Questions ###

We currently don't have a discussion forum. For now, use the issue tracker to ask questions.

### Reporting Issues ###

If you believe you have found an issue or bad documentation, use
the GitHub issue tracker to report the problem to the Chrono maintainers.

When reporting an issue, please provide the version of chrono is use (`chrono version`)

[1]: https://github.com/gochrono/chrono/wiki/Quick-Start

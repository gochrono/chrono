![Alt chrono](./.github/full-logo-with-tagline.png)

[![GitHub (pre-)release](https://img.shields.io/github/release/gochrono/chrono/all.svg)](https://github.com/gochrono/chrono/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/gochrono/chrono)](https://goreportcard.com/report/github.com/gochrono/chrono)
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

## Choose How to Install ##

The simplest way to install Chrono is to download the latest binary from the [releases page](https://github.com/gochrono/chrono/releases).
The binaries have no external dependencies.

To contribute to the Chrono project or documentation, you should [fork the GitHub project](https://github.com/gochrono/chrono#fork-destination-box) and clone it to your machine.

Alternatively you can install Chrono by building it yourself. This ensures you're running the absolute bleeding edge version.

### Building Chrono from the Source ###

#### Prequisite Tools ####

* [Git](https://git-scm.com/)
* [Go (at least Go 1.11)](https://golang.org/dl/)


#### Downloading the source ####

As of right now, Chrono uses [dep](https://github.com/golang/dep) to manage dependencies. We'll be moving to Go Modules in the near future.

The easiest to get the source is to clone Chrono in a directory outside of `GOPATH`, for example:

``` bash
mkdir $HOME/src && cd $HOME/src
git clone https://github.com/gochrono/chrono.git
cd chrono
dep ensure
go install
```

**If you are a Windows user, substitute the `$HOME` environment variable above with `%USERPROFILE%`.**

## Contributing to Chrono ###

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

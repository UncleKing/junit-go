# go-junit-report

Simple Library that can be used to generate xml reports, suitable for applications that
expect junit xml reports (e.g. [Jenkins](http://jenkins-ci.org)).

[![Build Status][travis-badge]][travis-link]
[![Report Card][report-badge]][report-link]
[![Code Coverage](https://codecov.io/gh/UncleKing/junit-go/branch/master/graph/badge.svg)](https://codecov.io/gh/UncleKing/junit-go)

## Installation

Go version 1.3 or higher is required. Install or update using the `go get`
command:

```bash
go get -u github.com/UncleKing/junit-go
```

## Contribution

Create an Issue and discuss the fix or feature, then fork the package.
Clone to github.com/UncleKing/junit-go.  This is necessary because go import uses this path.

## Run Tests
go test

## Usage
Please check the test code to see how the library can be used.

[travis-badge]: https://travis-ci.org/UncleKing/junit-go.svg
[travis-link]: https://travis-ci.org/UncleKing/junit-go
[report-badge]: https://goreportcard.com/badge/github.com/UncleKing/junit-go
[report-link]: https://goreportcard.com/report/github.com/UncleKing/junit-go

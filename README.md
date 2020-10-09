[![Logo](https://www.ga4gh.org/wp-content/themes/ga4gh-theme/gfx/GA-logo-horizontal-tag-RGB.svg)](https://ga4gh.org)

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://opensource.org/licenses/Apache-2.0)
[![Go Report](https://goreportcard.com/badge/github.com/ga4gh/htsget-refserver-utils)](https://goreportcard.com/badge/github.com/ga4gh/htsget-refserver-utils)
[![Travis (.org)](https://img.shields.io/travis/ga4gh/htsget-refserver-utils?style=flat-square)](https://travis-ci.org/ga4gh/htsget-refserver-utils)
[![Coveralls github](https://img.shields.io/coveralls/github/ga4gh/htsget-refserver-utils?style=flat-square)](https://coveralls.io/github/ga4gh/htsget-refserver-utils?branch=master)

# htsget-refserver-utils
Helper utilities for the [htsget reference server](https://github.com/ga4gh/htsget-refserver)

https://img.shields.io/coveralls/github/ga4gh/htsget-refserver-utils?style=flat-square

## Installation

Clone repo and build from source:
```
git clone https://github.com/ga4gh/htsget-refserver-utils.git
cd htsget-refserver-utils
go build -o ./htsget-refserver-utils .
./htsget-refserver-utils
```

Install binary to $GOBIN:
```
go get -u "go get -u "github.com/ga4gh/htsget-refserver-utils"
go install "github.com/ga4gh/htsget-refserver-utils"
htsget-refserver-utils
```

## Usage

`htsget-refserver-utils` contains multiple subcommands that can be specified on command line, i.e.
```
htsget-refserver-utils ${SUBCOMMAND} ${ARG1} ${ARG2} ... ${ARGN}
```

### Subcommands

* modify-sam
    * streams a SAM file from stdin, emitting custom fields and tags to stdout
    * ex: `htsget-refserver-utils modify-sam -fields QNAME,FLAG -tags NM,MD -notags HI`
* help
    * prints help message

## Testing

Run all tests
```
go test ./...
```

Run all tests and produce coverage report
```
go test ./... -coverprofile=cp.out
```
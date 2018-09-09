# pqerrors

[![Build Status](https://img.shields.io/travis/dhui/pqerrors/master.svg)](https://travis-ci.org/dhui/pqerrors)
[![GoDoc](https://godoc.org/github.com/dhui/pqerrors?status.svg)](https://godoc.org/github.com/dhui/pqerrors)
![Supported Go Versions](https://img.shields.io/badge/Go-1.11-lightgrey.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/dhui/pqerrors)](https://goreportcard.com/report/github.com/dhui/pqerrors)
[![GitHub Release](https://img.shields.io/github/release/dhui/pqerrors.svg)](https://github.com/dhui/pqerrors/releases)

pqerrors is a Go library that provides constants to make handling [pq](https://github.com/lib/pq) errors easier

## Example Usage:

### Error Classes

```golang
    var err error // err from database/sql
    if e, ok := err.(pq.Error); ok {
        switch e.Code.Class() {
        case pqerrors.PqErrClassNoData:
             // Handle error
        case pqerrors.PqErrClassIntegrityConstraintViolation:
             // Handle error
        default:
            // Handle unexpected error
        }
    }
```

### Error Codes

```golang
    var err error // err from database/sql
    if e, ok := err.(pq.Error); ok {
        switch e.Code {
        case pqerrors.PqErrCodeDataExceptionNullValueNotAllowed:
             // Handle error
        case pqerrors.PqErrCodeIntegrityConstraintViolationUniqueViolation:
             // Handle error
        default:
            // Handle unexpected error
        }
    }
```



## How to update error class and code constants:
1. Copy the [table of Postgres errors](https://www.postgresql.org/docs/current/static/errcodes-appendix.html) into `errors_table.txt`
1. Run `go generate`
1. Ensure that tests still pass: `go test -v`
1. Commit the changes
1. [Open a PR](https://github.com/dhui/pqerrors/pulls)
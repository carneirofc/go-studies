# go-studies

Notes and practice code

## Table of Contents

[[TOC]]

## ["The Go Programming Language"](http://www.gopl.io/)

### Chapter 1

A series of simple tools, such as web clients and web servers.
The gif bellow was generated and served using the program [./go-programming-language/ch1/lissajous.go](./go-programming-language/ch1/lissajous.go)

![lissajous gif](./docs/ch1/lissajous.gif)

### Chapter 2

## Notes
###  Build
On should not specify the direct file path when building something. Point it to the module instead *~the directory~* !

```command
go build ./cmd/nvim-man/
```

### Tests
The same applies to tests, always specify a directory. One might consider changing the directory into the module itself and run the tests from there.

```command
go test -v ./utilities_tests/
```

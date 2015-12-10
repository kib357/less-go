# less-go

[![Build Status](https://secure.travis-ci.org/kib357/less-go.png?branch=master)](http://travis-ci.org/kib357/less-go)

Less compiler for golang

Builds css using original less compiler and Duktape embeddable Javascript engine

## Installation

    go get github.com/kib357/less-go

## Command Line usage

    go build
    cd $GOPATH/src/github.com/kib357/less-go/lessc
    ./lessc --input="inputFile" --output="outputFile"

Example:

    ./lessc --input="./styles.less" --output="./styles.css"

## Programmatic usage

    err := less.RenderFile("./styles.less", "./styles.css", map[string]interface{}{"compress": true})

## Limitations

Because of using C++ Javascript engine, cross compilation not supported.

CLI interface doesnt support options

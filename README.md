# less-go

[![Build Status](https://secure.travis-ci.org/kib357/less-go.png?branch=master)](http://travis-ci.org/kib357/less-go)

Less compiler for Golang

Builds CSS using original LESS compiler and Duktape embeddable Javascript engine

## Status

This project is a work-in-progress, we accept pull requests.

## Installation

    go get github.com/kib357/less-go

## Command Line usage

    cd $GOPATH/src/github.com/kib357/less-go/lessc
    go build
    ./lessc --input="inputFile" --output="outputFile"

Example:

    ./lessc --input="./styles.less" --output="./styles.css"

## Programmatic usage

    err := less.RenderFile("./styles.less", "./styles.css", map[string]interface{}{"compress": true})
    
### Function reference

#### RenderFile(input, output string, mods ...map[string]interface{}) error

Renders LESS and generates output CSS

#### SetReader(customReader Reader)

    type Reader interface {
	    ReadFile(string) ([]byte, error)
    }

Sets a custom reader for .less files. You can use it to replace standard input from file system to string. Example:

    type LessReader struct{}

    func (LessReader) ReadFile(path string) ([]byte, error) {
	    return []byte(".class { width: (1 + 1) }"), nil
    }
    
    less.SetReader(LessReader)
    
#### SetWriter(customWriter Writer)

    type Writer interface {
	    WriteFile(string, []byte, os.FileMode) error
    }
    
Analogue of custom reader, but for output CSS

## Limitations

Because of using C++ Javascript engine, cross compilation not supported. Not tested on Windows.

CLI interface doesnt support options

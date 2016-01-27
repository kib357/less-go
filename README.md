# less-go

[![Build Status](https://secure.travis-ci.org/kib357/less-go.png?branch=master)](http://travis-ci.org/kib357/less-go)

[Less](http://lesscss.org/) compiler for [Golang](https://golang.org/)

Builds CSS using original [Less compiler](https://github.com/less/less.js) and [Duktape](http://duktape.org/) embeddable Javascript engine

## Status

This project is a work-in-progress, we accept pull requests.

## Installation

```
    go get github.com/kib357/less-go
```

## Command Line usage

```
    cd $GOPATH/src/github.com/kib357/less-go/lessc
    go get
    go build
    ./lessc --input="inputFile" --output="outputFile"
    ./lessc -i inputFile -o outputFile
```

Examples:

```
    ./lessc --input="./styles.less" --output="./styles.css"
    ./lessc -i styles.less -o styles.css
```

More about usage you can see in cli help:

```
    ./lessc -h
```

## Programmatic usage

```go
    import "github.com/kib357/less-go"

    func main() {
        err := less.RenderFile("./styles.less", "./styles.css", map[string]interface{}{"compress": true})
    }
```

### Function reference

#### RenderFile(input, output string, mods ...map[string]interface{}) error

Renders Less and generates output CSS.

#### SetReader(customReader Reader)

```go
    type Reader interface {
	    ReadFile(string) ([]byte, error)
    }
```

Sets a custom reader for .less files. You can use it to replace standard input from file system to another. Example:

```go
    type LessReader struct{}

    var lessFiles = map[string][]byte{"styles": []byte{".class { width: (1 + 1) }"}}

    func (LessReader) ReadFile(path string) ([]byte, error) {
	    lessFile, ok := lessFiles[path]
        if !ok {
            return "", errors.New("path not found")
        }
        return lessFile, nil
    }

    func main() {
        less.SetReader(LessReader)
        ...
    }

```

#### SetWriter(customWriter Writer)

```go
    type Writer interface {
	    WriteFile(string, []byte, os.FileMode) error
    }
```

Analogue of custom reader, but for output CSS.

## Current limitations

Because of using C Javascript engine, cross compilation not supported. Not tested on Windows.

CLI interface doesnt support options

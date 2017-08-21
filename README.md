# less-go 
## THIS PROJECT IS NO LONGER MAINTAINED, feel free to fork and use as boilerplate for further usages

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
    import (
        "github.com/kib357/less-go"
    	"bytes"
    )

    func main() {
        w := new(bytes.Buffer)
        lessCompiler := less.NewLessCompiler()

        err := lessCompiler.Compile("assets/test/signature.less", w, map[string]interface{}{"compress": true})

        if err != nil {
            t.Error("Render error", err)
        }
        res := w.String()
    }
```

## Current limitations

Because of using C Javascript engine, cross compilation not supported. Not tested on Windows.

CLI interface doesnt support options

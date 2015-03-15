Native Go Telegram library
==========================

This is a WIP integration of C Telegram library (https://github.com/vysheng/tgl) into Go

[![Build Status](https://travis-ci.org/gloob/go-telegram.svg?branch=master)](https://travis-ci.org/gloob/go-telegram)

Version: 0.1-alpha

Documentation: http://godoc.org/github.com/gloob/go-telegram/tgl

LICENSE
-------

go-telegram is licensed under the Apache License, Version 3.0 (http://www.apache.org/licenses/LICENSE-2.0.html).

Current Development Status
--------------------------

* Alpha version
 * In planning phase.
 * ~~Any useful code yet.~~
 * Exposing the library into Go without any aliasing.

***

* Beta version
 * TbD

How to Install
--------------

Sadly we are facing two problems with the current project approach:
1. Unsupported git submodules in go get command. (https://github.com/golang/go/issues/7764)
2. Currently tgl library doesn't have a pkg-config package defined, so we need to compile it by ourselves in the correct location. (TODO: Need to create a issue for this or provide a pull request to the tgl project)

So you need to do some manual arrangment to succesfully install the package.

    $ go get -u github.com/gloob/go-telegram/tgl # It will fail tgl.go.h:4:21: fatal error: tgl/tgl.h: No such file or directory
    $ cd $GOPATH/src/github.com/gloob/go-telegram/tgl
    $ git submodule update --init --recursive
    $ cd lib/tgl && ./configure && make
    $ go get -u github.com/gloob/go-telegram/tgl # It will works now.

Current Status
--------------

 * Linux
   * Not yet!
 * Mac OS
   * Not yet!
 * Windows
   * planned for 2015 Q4

Sample Program
--------------

```go
package main

import (
    "fmt"
    "github.com/gloob/go-telegram/tgl"
)

func main() {
    // TODO: Create sample program. :-)
}
```

You can contact me at gloob@litio.org

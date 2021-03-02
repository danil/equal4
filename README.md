equal4
======

[![Build Status](https://cloud.drone.io/api/badges/danil/equal4/status.svg)](https://cloud.drone.io/danil/equal4)
[![Go Reference](https://pkg.go.dev/badge/github.com/danil/equal4.svg)](https://pkg.go.dev/github.com/danil/equal4)

Comparison of an error messages for Go.  
Source files are distributed under the BSD-style license
found in the [LICENSE](./LICENSE) file.

Install
-------

    go get github.com/danil/equal4@v0.13.0

Usage
-----

```go
package main

import (
    "errors"
    "fmt"

    "github.com/danil/equal4"
)

func main() {
    fmt.Println(equal4.ErrorEqual(errors.New("foobar"), errors.New("bar")))
    fmt.Println(equal4.ErrorContains(errors.New("foobar"), "bar"))
    fmt.Println(equal4.ErrorMatch(errors.New("foobar"), regexp.MustCompile("bar$")))
}
```

Output:

    false
    true
    true

# hello
say hello

## Install
import code
```bash
go get github.com/codeExpert666/hello@latest
```

import cmd
````bash
go install github.com/codeExpert666/hello/cmd/hello@latest
````

## Example
Here's a simple example as follows:
```go
package main

import (
    "fmt"
    "github.com/codeExpert666/hello"
)

func main() {
    result := hello.Hello("Jack")
    fmt.Println(result)
}
```
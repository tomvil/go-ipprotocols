## go-ipprotocols

**go-ipprotocols** is a Go library that provides a mapping of IP protocol numbers to names and vice versa. It uses a static mapping based on IANA's protocol numbers, allowing developers to easily retrieve protocol names from numbers and numbers from protocol names.

### Features
- **Get protocol name by number**: Quickly look up the name of an IP protocol using its number.
- **Get protocol number by name**: Find the number associated with a given IP protocol name.

### Installation

You can install this package by using `go get`:

`go get github.com/tomvil/go-ipprotocols`

### Usage

After installation, you can import the library into your Go project and use the functions `GetProtocolName` and `GetProtocolNumber` to perform protocol lookups.

#### Example Code:
```
package main

import (
    "fmt"
    "log"
    "github.com/tomvil/go-ipprotocols"
)

func main() {
    // Get protocol name by number
    name, err := ipprotocols.GetProtocolName(6)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Protocol number 6 is:", name) // Output: Protocol number 6 is: TCP

    // Get protocol number by name
    number, err := ipprotocols.GetProtocolNumber("TCP")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Protocol TCP has number:", number) // Output: Protocol TCP has number: 6
}
```

### Contributing

Contributions are welcome! If you'd like to add more features, fix bugs, or improve documentation, feel free to open a pull request. Please ensure that all changes include relevant tests.

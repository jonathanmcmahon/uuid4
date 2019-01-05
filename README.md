# uuid4 #

## Go library for generating RFC 4122 version 4 UUIDs ##

Package uuid4 provides a pure Go implementation of UUID version 4 variant 1 as defined in RFC-4122. This package currently supports:

* Creation of new UUIDs in string representation

* Creation of new UUIDs in byte representation
q
[See the RFC 4122 spec here.](https://tools.ietf.org/html/rfc4122)

## Requirements ##

This package has been tested on Go 1.10.

## Installation ##

```
go get github.com/jonathanmcmahon/uuid4
```

## Usage ##

```
import (
	"fmt"

	"github.com/jonathanmcmahon/uuid4"
)

func main() {
	u, err := uuid4.New()
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	fmt.Println()

	b, err := uuid4.NewBytes()
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}
```

Example output:

```
311378dd-1e12-40d0-852f-c209bf34a5b0

[191 225 203 239 127 91 71 230 169 110 55 39 242 134 245 44]
```

## License ##

This code is released under the [MIT License](./LICENSE). 

## References ##

* [RFC 4122](https://tools.ietf.org/html/rfc4122)

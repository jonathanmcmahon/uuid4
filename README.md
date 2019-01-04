# uuid4 #

### Go library for generating RFC 4122 version 4 UUIDs ###

See the spec here: [https://tools.ietf.org/html/rfc4122](https://tools.ietf.org/html/rfc4122)

### Usage ###

Go get it:

```
go get github.com/jonathanmcmahon/uuid4
```

Then:

```
import (
	"fmt"

	"github.com/jonathanmcmahon/uuid4"
)

func main() {
	u, err := uuid4.NewString()
	if err == nil {
		fmt.Println(u)
	}
}
```

Output:

```
2dea57d2-e37e-e584-bdb3-921e1f0e5284
```
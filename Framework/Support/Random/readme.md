
```go
package main

import (
	"fmt"
	"github.com/gohouse/random"
)

func main() {
	fmt.Println(random.Rand())
	fmt.Println(random.Random(12))
	fmt.Println(random.Random(12,random.T_ALL))
	fmt.Println(random.RandomBetween(6, 11))
	fmt.Println(random.RandomBetween(6,11, random.T_ALL))
}
```

```go
StrNumberic  = `0123456789`
StrLowercase = `abcdefghijklmnopqrstuvwxyz`
StrCapital   = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
```

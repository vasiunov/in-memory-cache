## Example

```go
package main

import (
	"fmt"
	"github.com/vasiunov/in-memory-cache"
)

func main() {

	someCache := NewCache()
	fmt.Println(someCache)

	someCache.Set("userID", 59)
	fmt.Println(someCache)

	userId := someCache.Get("userID")
	fmt.Println(userId)

	someCache.Delete("userID")
	userId = someCache.Get("userID")

	fmt.Println(userId)
}
```
Description is available here - https://pkg.go.dev/github.com/vasiunov/in-memory-cache

## Example

```go
package main

import (
	"fmt"
	"github.com/vasiunov/in-memory-cache"
)

func main() {

	someCache := cache.NewCache()
	fmt.Println(someCache) // {map[]}

	someCache.Set("userID", 59)
	fmt.Println(someCache) // {map[userID:59]}

	userId := someCache.Get("userID")
	fmt.Println(userId) // 59

	someCache.Delete("userID")
	userId = someCache.Get("userID")

	fmt.Println(userId) // <nil>
}
```
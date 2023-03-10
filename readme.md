Package cache is used for caching data. Each key-value pair expires in 5 sec.
Description is available here - https://pkg.go.dev/github.com/vasiunov/in-memory-cache
Play code - https://go.dev/play/p/dBzJ5Uv-Kjc

## Example

```go
package main

import (
	"fmt"
	"github.com/vasiunov/in-memory-cache"
)

func main() {
	someCache := NewCache()

	for i := 0; i <= 10; i++ {
		time.Sleep(10 * time.Millisecond)
		go someCache.Set(fmt.Sprintf("%d", i), i*i)
	}

	for i := 0; i <= 10; i++ {
		time.Sleep(10 * time.Millisecond)
		go fmt.Println(someCache.Get(fmt.Sprintf("%d", i))) // 0 1 4 9 16 25 36 49 64 81 100
	}

	time.Sleep(time.Second * 3)
	fmt.Println("2sec left", someCache) // 2sec left {map[0:0 1:1 10:100 2:4 3:9 4:16 5:25 6:36 7:49 8:64 9:81] 5000000000 0xc0000b2000}

	time.Sleep(time.Second * 3)
	fmt.Println(someCache.Get("1"), "time is out") // <nil> time is out
	fmt.Println(someCache) // {map[] 5000000000 0xc0000b2000}
}
```
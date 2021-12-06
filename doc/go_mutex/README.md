# Go: mutex

- [map print](#map-print)

[↑ top](#go-map)

<br><br><br><br>

#### concurrent bug
In concurrent cases, there may be bugs if locks are not used

[Code](https://go.dev/play/p/0_ekQwdAFeN)

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
```
# Go: mutex

- [map print](#map-print)

[↑ top](#go-map)


#### concurrent bug
In concurrent cases, there may be bugs if locks are not used.

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

#### principle of realization

##### simple mutex
Use a variable to mark whether the current lock is held by a certain goroutine.

```go
// CAS
func cas(val *int32, old, new int32) bool
// block
func semacquire(*int32)
// wake up
func semrelease(*int32)

type Mutex struct {
	key int32 // whether the lock is held
	sema int32 // semaphore used to block/wake up the goroutine 
}

func xadd(v *int32, delta int32) (new int32) {
	for {
		v := *val
		if cas(val, v, v + delta) {
			return v + delta
		}
	}

	panic("unreached")
}

func (m *Mutex) Lock() {
	if xadd(&m.key, 1) == 1 {
		return
	}

	semcquire(&m.sema) // block
}

func (m *Mutex) Unlock() {
	if xadd(&m.key, -1) == 0 {
		return
	}

	semrelease(&m.sema) // wake up
}
```

##### Give the new goroutine a chance
todo
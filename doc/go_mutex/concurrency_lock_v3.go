package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Incr() {
	c.Lock()
	c.count++
	c.Unlock()
}

func (c *Counter) Count() int {
	c.Lock()
	defer c.Unlock()
	return c.count

}

func main() {
	counter := Counter{}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Incr()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Count())
}

package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, lastUpdated: %v", c.total, c.lastUpdated)
}

func main() {
	counter := Counter{
		total:       10,
		lastUpdated: time.Now(),
	}
	fmt.Println(counter.String())
	// total: 10, lastUpdated: 2022-12-24 11:46:11.3784929 +0300 MSK m=+0.001583101
	time.Sleep(2 * time.Second)
	counter.Increment()
	fmt.Println(counter.String())
	// total: 11, lastUpdated: 2022-12-24 11:46:13.4013493 +0300 MSK m=+2.024439501
}

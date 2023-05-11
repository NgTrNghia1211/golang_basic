package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var count int = 1
var buffer [10]int

func Producer() {
	for {
		x := rand.Intn(99)
		time.Sleep(time.Second)

		// Add to buffer
		m.Lock()
		if count < 10 {
			buffer[count] = x
			count++
		}
		m.Unlock()
		fmt.Printf("Prod %v\n", x)
	}
	wg.Done()
}

func Consumer() {
	for {
		y := -1
		m.Lock()
		if count > 0 {
			y = buffer[count-1]
			count--
		}
		m.Unlock()
		// Consume
		fmt.Printf("Got %v\n", y)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {
		wg.Add(2)
		if i%2 == 1 {
			go Producer()
		} else {
			go Consumer()
		}
	}
	wg.Wait()
}

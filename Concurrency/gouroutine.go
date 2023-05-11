package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	fmt.Println("Goroutine ----------")
	/* go CountSth("Sheep") // lam di - t di lam nhung ham duoi
	// goroutine main end - program end
	CountSth("Fish") */

	/* var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		CountSth("SGB win win")
		wg.Done()
	}()
	go func() {
		CountSth("SKT win win")
		wg.Done()
	}()
	wg.Wait() */

	// MUTEX
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // khoa de doc
		go SayHello()
		m.Lock() // khoa de ghi
		go Increment()
	}
	wg.Wait()
	fmt.Println("Done Program")
}

func CountSth(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}

func SayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func Increment() {
	counter++
	m.Unlock()
	wg.Done()
}

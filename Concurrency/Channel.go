package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch := make(chan int) // make(chan int, 50) // channel chua dc 50 so nguyen
	/* go func() {
		i := <-ch // cho de channel co gia tri
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second)
		ch <- 43
		wg.Done()
	}() */
	go func(ch <-chan int) { // Only got
		i := <-ch
		fmt.Println("Got: ", i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) { // Only send
		ch <- 3
		//fmt.Println("Recive", ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

// deadlock : gui qua nhieu, nhan qua nhieu
// close -> ngừng nhận
// select : khi nhieu channel

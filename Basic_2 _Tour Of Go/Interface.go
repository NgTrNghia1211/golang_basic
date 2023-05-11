package main

import "fmt"

func main() {
	fmt.Println("Interface --- ----")
	var v Writer = ConsoleWriter{}
	n, er := v.Write([]byte("Hello"))
	fmt.Println(n, er)
	v.Write([]byte("Hello World"))
	fmt.Println("Ex: 2 ----------")
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 5; i++ {
		fmt.Println(inc.Increment())
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Incrementer interface {
	Increment() int //prototype
}

type IntCounter int // dung con tro de thao tac tren bo nho

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

type ConsoleWriter struct{}

func (c ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

/*

 */

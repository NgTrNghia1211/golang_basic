package main

import "fmt"

func main() {
	fmt.Println("Interface --- ----")

}

type Writer interface {
	Write([]byte) (int, error)
}
type Closer interface {
	Close() error
}

type WriteClose_er interface {
	Writer
	Closer
}

/*
	interface co the chua interface
*/

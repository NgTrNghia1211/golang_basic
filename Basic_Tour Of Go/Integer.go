package main

import (
	"fmt"
)

var (
	j int     = 10
	k int     = 3
	i uint64  = 8
	c uint64  = i << 3 // left
	f float32 = 1.5
	z         = complex(3, 5)
)

func main() {
	fmt.Printf("Value: %v, Type: %T\n", c, c)
	fmt.Printf("Value: %v, Type: %T\n", z, z)
	fmt.Printf("Value: %v, Type: %T\n", real(z), real(z))
	fmt.Println(i, c)
	fmt.Println(j & k)
	fmt.Println(j | k)
}

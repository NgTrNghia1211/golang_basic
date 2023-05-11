package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Pick a number: ", rand.Intn(10))
	fmt.Println(math.Pi)
}

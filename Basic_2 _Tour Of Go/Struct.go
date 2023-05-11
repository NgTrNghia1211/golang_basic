package main

import "fmt"

type Student struct {
	X string
	Y int
}

func main() {
	student_1 := Student{"Nguyen Trong Nghia", 0} //type 1
	var i int = 72
	var t *int
	p := &i
	fmt.Println(*p)
	*p = 72 / 6
	fmt.Println(i)
	fmt.Println(t)

	k := &student_1
	k.X = "Nghia Ng"
	(*k).Y = 2013
	fmt.Println(student_1.X)
	fmt.Println(student_1.Y)
	fmt.Println(k)
}

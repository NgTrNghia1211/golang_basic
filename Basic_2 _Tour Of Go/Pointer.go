package main

import "fmt"

func main() {
	fmt.Println("Pointers ---------------------------")
	/* fmt.Println("Gán giá trị:")
	var a int = 12
	var b int = a
	a = 24
	fmt.Println("Giá trị: ", a, b)
	fmt.Println("Con trỏ: ")
	var c *int = &a
	a = 26
	fmt.Println("Giá trị: ", a, (*c))
	*c = 27
	fmt.Println("Giá trị: ", a, *c)
	fmt.Println("Array/Slice ------------------------------")
	arrA := [3]int{1, 2, 3}
	slcA := []int{1, 2, 3}
	d := arrA
	e := slcA
	fmt.Println("arrA, slcA, b, c: ", arrA, slcA, d, e)
	arrA[2] = 4
	slcA[2] = 4
	fmt.Println("arrA, slcA, b, c: ", arrA, slcA, d, e)
	// khong duoc &a - 4 */
	var a myStruct = myStruct{2}
	var b *myStruct = new(myStruct)
	var c *myStruct
	b.number = 12 //(*b.number)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	// map cung la mot dang pointer nhu slice
}

type myStruct struct {
	number int
}

// sach, ki truoc, trac nghiem, giay,
// but chi

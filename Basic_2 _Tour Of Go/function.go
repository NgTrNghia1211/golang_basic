package main

import "fmt"

func main() {
	fmt.Println("Function -------------")
	fmt.Println("Basic: ", larger(3, 4))
	b := []int{1, 3, 2, 5, 8, 6}
	fmt.Println("Basic: ", findMax(b))
	fmt.Println("Basic: ", findMaxRec(b, len(b)))
	fmt.Println("Basic: ", findSum(b, len(b)))
}

func larger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMax(a []int) (max int) {
	for index := range a {
		if a[index] > max {
			max = a[index]
		}
	}
	return max
}

func findMaxRec(a []int, i int) int {
	if i == 1 {
		return a[0]
	} else {
		return larger(findMaxRec(a, i-1), a[i-1])
	}
}

func findSum(a []int, i int) int {
	if i == 1 {
		return a[0]
	} else {
		sum := a[i-1]
		return sum + findSum(a, i-1)
	}
}

/*
	toán tử ... vào vị trí cuối cùng
	Hàm ẩn danh
	func() {

	} ()
	func (<name> typestruct) ten ham() //method

	Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:
	while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
	-There are two reasons to use a pointer receiver.

	The first is so that the method can modify the value that its receiver points to.

	The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

	In this example, both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.

	In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

*/

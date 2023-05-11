package main

import "fmt"

type Vertex struct {
	x int
	y int
}
type Vertex1 struct {
	X, Y int
}

var (
	v1 = Vertex1{1, 2}  // has type Vertex
	v2 = Vertex1{X: 1}  // Y:0 is implicit  return {1, 0}
	v3 = Vertex1{}      // X:0 and Y:0 		return {0, 0}
	p  = &Vertex1{1, 2} // has type *Vertex	return &{1, 2}
)

/*
	Struct cho phép con trỏ truy cập giá trị bên trong không thông qua dấu *
	==> (*p).X = p.X
*/

func main() {
	fmt.Println("Hello, World")

	a := make([]int, 0, 2)
	fmt.Printf("a %v, %v, %v\n", a, len(a), cap(a))
	a = append(a, 1, 2, 3, 4)
	fmt.Printf("a %v, %v, %v\n", a, len(a), cap(a))
	a = append(a, []int{5, 6, 7, 8}...) // them mot slice khac
	fmt.Printf("a %v, %v, %v\n", a, len(a), cap(a))
	// hien thuc stack va queue

}

/*
	Array
		Khởi tạo array
			<name> := [số lượng]<type>{giá trị}
			<name> := [...]<type>{giá trị}
			var <name> [số lượng]{giá trị}
			<name>[index]
			len(<name>) chiều dài
			<name_2> := <name_1> (copy 1 đưa vào 2 : 2 đổi 1 không đổi)
			<name_2> := &<name_1> (trở thành con trỏ)
	Slice (khắc phục fixed weight của Array) ~~ Con trỏ
			<name> := []<type>{giá trị}
			<name_2> := <name_1> (name_2 trỏ name_1)
			len(<name>) -- cap(<name>)
			a[(start):(end)]
		Khởi tạo
			a := make ([]<type>, num_1, num_2) (num_1 : len - num_2 : cap)
			append (slice, giá trị thêm) --> grow cap --> cap mới = cap cũ + 1)*2
	Unlike C, Go has no pointer arithmetic
*/

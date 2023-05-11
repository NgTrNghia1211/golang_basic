package main

import "fmt"

const (
	red    = 0
	yellow = 2
	green  = iota
	black
)

/*
const (
	_ = iota + (x)
	red		//1
	yellow	//2
	green	//3
	black	//4
)
*/

func main() {
	const Pi float64 = 3.14

	fmt.Printf("%v, %T\n", Pi, Pi)
	fmt.Printf("%v, %T\n", red, red)
	fmt.Printf("%v, %T\n", yellow, yellow)
	fmt.Printf("%v, %T\n", green, green)
}

/*
	Khai báo hằng số
		const <name> <type> = giá trị
	Đặc điểm
		cũng có đặc tính nếu có const trong block cùng tên const global
		giá trị khởi tạo không cần trình biên dịch phải tính toán
	Kiểu dữ liệu Enum _ tập hợp nhiều hằng số
	Khởi tạo Enum bằng itoa

*/

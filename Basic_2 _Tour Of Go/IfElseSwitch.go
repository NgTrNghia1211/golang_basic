package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")
	/* temp := make(map[string]int)
	temp["Nguyen Trong Nghia"] = 20
	fmt.Println(temp)
	// if age, isExist := temp[]; isExist == true {} //- vẫn chấp nhận
	age, isExist := temp["Nguyen Trong Nghia"]
	if isExist == true {
		fmt.Println(age)
	} */
	var number int = 3
	switch {
	case number <= 3:
		fmt.Println("<=3")
		fallthrough // tiếp tục kiểm tra các trường hợp khác
	case number <= 5:
		fmt.Println("<=5")
	default:
		fmt.Println("Not other")
	}
}

/*
Golang luôn bắt buộc khối lệnh trong if else phải đặt trong {}
Thứ tự ktra biểu thức trái qua phải.
Variables declared by the statement are only in scope until the end of the if.
! =  not
switch case nếu đã thỏa 1 case sẽ tự động thoát luôn
*/

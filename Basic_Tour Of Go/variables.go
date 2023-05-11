package main

import (
	"fmt"
	"strconv"
)

var ( //Khai báo biến global
	n   int    = 10
	m   int    = 12
	str string = "Hello, World"
)

func main() {
	var i int
	i = 100
	k := 1000
	fmt.Println(k + i)
	fmt.Printf("Export: %v, Type: %T", str, str)
	var j string = strconv.Itoa(i)
	fmt.Printf("\n%v, %T", j, j)
}

/* Shift + Alt + A
Khai báo biến trong trong Go
	var <name> <Type> ( = giá trị khởi tạo)
	<name> := (giá trị khởi đầu)
	Println - Printf
	%v , %T : type
Global Scope - Block Scope
	Global Scope chỉ có thể khai báo: var <name> <type> = giá trị
	Khai báo nhiều biến
Shadow
	Ưu tiến biến block scope trong trường hợp khai báo biến block cùng tên với biến global (phạm vi block)
Không cho phép biến dư thừa, unused variable

Exported Variable --> Đặt tên bắt đầu bằng in hoa ==> cho phép các package khác truy xuất
Đặt tên theo kiểu camel -- > chỉ bắt đầu bằng alphabet

Convert Type : <type>(giá trị) (nhỏ --> sang lớn)
			   int -> string: --> ra unicode
			   import strconv --> thỏa mãn
*/

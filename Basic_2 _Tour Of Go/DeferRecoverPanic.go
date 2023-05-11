package main

import "fmt"

func main() {
	/* fmt.Println("Hello, Guys")
	// Defer - truoc khi ham ket thuc moi thuc hien code sau chu defer
	// Cac lenh defer duoc dua vao mot stack nen khi se xuat ra theo LIFO
	fmt.Println("Line 1")
	defer fmt.Println("Line 2")
	fmt.Println("Line 3")
	defer fmt.Println("Line 4")
	defer fmt.Println("Line 5") */
	// Panic - loi khong mong muon khi thuc thi chuong trinh
	// nen xem de lay vi du
	// khi gặp panic ctrinh se dung luon
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error: ", err)
		}
	}()
	a := 10
	b := 1
	if b == 0 {
		panic("Chia cho so 0")
	} else {
		fmt.Println("Result: ", a/b)
	}
	fmt.Println("End program.")
}

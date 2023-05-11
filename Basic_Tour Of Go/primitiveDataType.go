package main

import "fmt"

func main() {
	fmt.Println("hello, World")
	var a int = 8
	fmt.Printf("%v ,%T\n", a<<3, a<<3) // dịch trái
	fmt.Printf("%v ,%T\n", a>>3, a>>3) // dịch phải
}

/*
3 Nhóm Kiểu nguyên thủy
	1. Boolean (2 giá trị true, false --  mặc định là false)
		var a bool = true
	2. Numerics
		Integer (signed and unsigned) Hỗ trợ các phép tính
			signed 		:	int8	int16	..
			unsigned	:	uint8	uint16	..
			Dịch bit : & ! ^ &^
		Float (float32 float64) Hỗ trợ tính toán
		Complex (complex 64, complex 128) (float 32.32 | float 64.64)
			real(complex) - imag(complex)
			complex(real, imag)
	3. Text
		String () Hỗ trợ append +
		Byte UTF-8 ''
			[]byte : số nguyên uint8
		Rune UTR-32 (chữ tượng hình) byte lưu được thì rune lưu được
*/

package main

import "fmt"

func main() {
	fmt.Println("Nice, Hello World!")
	/* for i := 1; i <= 10; i++ { // for basic
			fmt.Printf("Integer: %v\n", i)
		}
		fmt.Println("Only one condition")
		j := 1
		for j < 10 {
			fmt.Println(j)
			j++
		}
		fmt.Println("Nested Loop:")
	Loop: // nhãn cho loop
		for k := 1; k <= 5; k++ {
			for l := 1; l <= 5; l++ {
				if k > 1 {
					break Loop
				}
				fmt.Println("Integer: ", k*l)
			}
		} */
	mini_Slice := []int{2, 3, 4}
	for index, value := range mini_Slice { // range ho tro lay value, index
		fmt.Println("Index, Value: ", index, value)
	}
	mapping := map[string]int{
		"Nghia":      10,
		"NghiaNgu":   0,
		"NghiaOcCho": -100,
	}
	for key, val := range mapping {
		fmt.Println("Key, Value:", key, val)
	}
}

/*
Ngoai cu phap thi tuong tu C/C++
*/

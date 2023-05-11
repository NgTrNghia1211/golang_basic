package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	number int `Max is 100`
	name   string
	isMale bool
}

func main() {
	/* fmt.Println("Hello, World")
	studentNameGradeMap := map[string]int{ //slices không được key của map
		"Yone":   8,
		"Yasuo":  10,
		"Graves": 7,
	}
	fmt.Println(studentNameGradeMap)    // xuất theo bảng alphabet
	delete(studentNameGradeMap, "Yone") // delete khỏi map ==> sau khi xóa nếu truy cập lại trả về 0
	// isExist để kiểm tra
	fmt.Println(studentNameGradeMap)

	championHardMap := make(map[string]int)
	championHardMap["Zoe"] = 25
	fmt.Println(championHardMap) */
	studentNghia := Student{
		number: 2013873,
		name:   "NguyenNghia",
		isMale: true,
	}
	fmt.Println(studentNghia)

	studentYasuo := Student{}
	studentYasuo.number = 123
	studentYasuo.name = "Yasuo"
	studentYasuo.isMale = true
	fmt.Println(studentYasuo)

	t := reflect.TypeOf(studentYasuo)
	field, _ := t.FieldByName("number")
	fmt.Println(field)
}

/*
	isExist kiểm tra tồn tại
	len() độ dài
	copy --> con trỏ trỏ vào

	STRUCT ---------------
	// co the ke thua
	// tag
*/

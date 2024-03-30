package main

import (
	"fmt"
)

func main() {
	data := map[string]map[string]int{}
	// student, _ := std.MakeStudentDetails()

	// fmt.Println(student.GetStudent(""), "")

	// teacher, _ := std.MakeTeaher()

	// fmt.Println(" Teacher ", teacher.GetTeacher(""))
	data["test"] = nil
	data["test"]["1"] = 1
	data["test"]["2"] = 1
	fmt.Println("data ", data)

}

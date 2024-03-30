package main

func main() {

}

type Student struct {
	Name string
}

type StudentInterface interface {
	AddStudent(Student) int
	GetStudent(Student) Student
}

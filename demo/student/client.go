package student

import "fmt"

func (s StudentModel) GetStudent(string) StudentModel {

	fmt.Println("Student is ", s.Name)
	return StudentModel{}
}

func (s StudentModel) GetStudentList(string) []StudentModel {
	return []StudentModel{}
}

func (s StudentModel) SaveStudent(StudentModel) string {
	return ""
}

func (t TeacherModel) GetTeacher(name string) string {
	return t.Name
}

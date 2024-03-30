package student

type Student struct {
	Name string
}

type StudentInterface interface {
	AddStudent(Student) int
	GetStudent(Student) Student
}

func MakeStudentClient() StudentInterface {
	return &Student{}
}
func (s *Student) AddStudent(stu Student) int {

	return 10
}

func (s *Student) GetStudent(stu Student) Student {
	return stu
}

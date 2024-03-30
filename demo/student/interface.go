package student

type Student interface {
	GetStudent(string) StudentModel
	GetStudentList(string) []StudentModel
	SaveStudent(StudentModel) string
}

type Teacher interface {
	GetTeacher(string) string
}

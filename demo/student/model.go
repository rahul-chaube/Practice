package student

type StudentModel struct {
	Name        string
	RolleNumber string
	Class       string
}

func MakeStudentDetails() (Student, error) {

	student := new(StudentModel)
	student.RolleNumber = "20"
	student.Class = "FYJC"
	student.Name = "RAHUL CHUABE"

	return student, nil
}

type TeacherModel struct {
	Name string
}

func MakeTeaher() (Teacher, error) {
	return TeacherModel{Name: "rahul chaube"}, nil
}

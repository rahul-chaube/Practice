package main

import (
	"fmt"
	"unsafe"
)

type CustomeError struct {
	Err         error
	ErroMessage string
	Code        int
}

func (err *CustomeError) Error() string {

	return fmt.Sprintf("Error occured : %s with code %d", err.ErroMessage, err.Code)

}

type Persone struct {
	int
	string
	float32
}

func main() {
	_, err := Operation()
	fmt.Printf(" %+v ", err)
	p := Persone{
		12, "Rahul Chaube", 34.0,
	}
	fmt.Printf("%+v \n", p.float32)
	fmt.Println("Size 1 ", unsafe.Sizeof(p))
	empty := struct {
		a struct{}
		b struct{}
	}{}
	fmt.Println("Size 2", unsafe.Sizeof(empty))
}

func Operation() (int, error) {

	return 0, &CustomeError{Err: fmt.Errorf("Error is occured "), Code: 200}

}

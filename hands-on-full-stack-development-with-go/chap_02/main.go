package main

type Person struct {
	name string
	age  int
}

func (p Person) getName() {
	p.name = "Manh Cuong"
}

type Student struct {
	Person
	studentId int
}

func (s *Student) getStudentId() int {
	s.name = "Manh Cuong"
	return s.studentId
}

type MyInterface1 interface {
	getName()
}

type MyInterface2 interface {
	getStudentId() int
}

type FinalInterface interface {
	MyInterface1
	MyInterface2
}

func main() {
	p := Student{Person{"bob", 21}, 1}

	var in FinalInterface = &p
	in.getStudentId()
}

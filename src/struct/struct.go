package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	name  string
	class int
	grade Grade
}

type Grade struct {
	subject string
	score   string
}

func (s Student) ViewScore() {
	fmt.Println(s.grade)
}

func (s Student) AddScore(subject, score string) {
	s.grade.subject = subject
	s.grade.score = score
}

func ViewScore(s Student) {
	fmt.Println(s.grade)
}

// Person이 가진 기능
// 메소드: struct 밖에서 정의
func (p Person) PrintName() {
	fmt.Print(p.name)
}

func main() {

	p1 := Person{"Tom", 15}
	p2 := Person{name: "Nina", age: 21}
	p3 := Person{name: "Mike"}
	p4 := Person{}

	fmt.Println(p1, p2, p3, p4)

	p4.name = "Smith"
	p4.age = 24

	fmt.Println(p4)

	p1.PrintName()
	p2.PrintName()

	var s Student
	s.name = "rakan"
	s.class = 1

	s.grade.subject = "math"
	s.grade.score = "C"

	s.ViewScore()
	// 같은 코드
	ViewScore(s)

	// 값 복사
	// 호출하는 s와 메소드의 입력값 s는 서로 다른 객체
	// 기능으로 속성을 변경하려면 포인터 필요
	s.AddScore("science", "A")
	s.ViewScore()

}

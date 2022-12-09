package main

import "fmt"

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintGrade() {
	fmt.Println(s.class, s.grade)
}

func (s *Student) AddGrade(class string, grade string) {
	// s가 가리키는 변수에 해당하는 속성 변경
	// c언어에서는 화살표로 속성에 접근
	// go에서는 구분하지 않고 컴파일러가 알아서 처리
	// 메모리 주소만 복사하므로 성능에 유리
	// 값을 함수 인자로 받으면 전체 속성을 복사(메모리 낭비)
	s.class = class
	s.grade = grade
}

func main() {
	var a int
	var b int
	var p *int

	a = 3
	b = 2
	p = &a

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	p = &b

	fmt.Println(b)
	fmt.Println(p)
	fmt.Println(*p)

	c := 1

	Increase(c)
	fmt.Println(c) // 1

	IncreaseWithPointer(&c)
	fmt.Println(c)

	var s Student = Student{name: "Ashe", age: 23, class: "Math", grade: "A+"}

	s.AddGrade("Science", "B")
	s.PrintGrade()
}

// 값 복사
// x와 c는 다른 변수
// 함수가 실행될 때까지만 유효한 값
func Increase(x int) {
	x++
}

// p를 통해 c의 주소로 접근
// c의 값을 바꿈
// 복사되는 것은 같지만 주소가 들어있기 때문에 주소로 접근 가능
func IncreaseWithPointer(p *int) {
	*p++
}

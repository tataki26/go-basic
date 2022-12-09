package main

import "fmt"

// interface 이름은 메서드 이름에 er을 붙이는 형태
// interface 내부 메서드는 최대 3개
type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

// Student가 Stringer 구현
// Stringer의 메소드를 가지고 있기 때문
func (s Student) String() string {
	// 값을 그대로 문자열로 만듦
	// 출력하지 않고 문자열 반환
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func (s Student) GetAge() int {
	return s.Age
}

func main() {
	student := Student{"Teemo", 5}

	// Stringer 타입 변수 stringer
	// struct가 interface의 메소드를 포함하고 있으므로 해당 interface를 가리킬 수 있다
	stringer := student

	fmt.Printf("%s\n", stringer.String())
	// undefined error
	// fmt.Printf("%d\n", stringer.GetAge())
}

package impl

import "fmt"

type Service struct {
}

func (s Service) Foo() {
	fmt.Println("foo")
}

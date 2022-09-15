package main

import (
	"fmt"
	"stack"
)

func main() {
	s := new(stack.Mystack)
	for i := 1; i < 4; i++ {
		s.Push(i)
	}

	for i := 1; i < 5; i++ {
		p, err := s.Pop()
		if err == nil {
			fmt.Println(p)
		} else {
			fmt.Println(err)
		}
	}
	//s.isempty() this is error sice is empty is not exported in Mystack
}

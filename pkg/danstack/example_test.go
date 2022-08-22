package danstack_test

import (
	"fmt"

	"github.com/danwhitford/danstack/pkg/danstack"
)

func ExampleDanStack_Push() {
	stack := danstack.New[int]()
	stack.Push(1)
}

func ExampleDanStack_Pop() {
	stack := danstack.New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	v, _ := stack.Pop()
	fmt.Println(v)
	v, _ = stack.Pop()
	fmt.Println(v)
	v, _ = stack.Pop()
	fmt.Println(v)
	// Output:
	// 3
	// 2
	// 1
}

func ExampleDanStack_Empty() {
	stack := danstack.New[int]()
	if stack.Empty() {
		fmt.Println("Is empty")
	} else {
		fmt.Println("Not empty")
	}

	stack.Push(1)
	if stack.Empty() {
		fmt.Println("Is empty")
	} else {
		fmt.Println("Not empty")
	}

	stack.Pop()
	if stack.Empty() {
		fmt.Println("Is empty")
	} else {
		fmt.Println("Not empty")
	}

	// Output:
	// Is empty
	// Not empty
	// Is empty
}
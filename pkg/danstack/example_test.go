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
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for !stack.Empty() {
		v, _ := stack.Pop()
		fmt.Println(v)
	}
	// Output:
	// 3
	// 2
	// 1
}

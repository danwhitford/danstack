// Package danstack provides a simple, minimal stack implementation.
package danstack

import "fmt"

const initialSize = 8

// A DanStack represents an array-backed stack.
type DanStack[T any] struct {
	top   int
	items []T
}

// New will initialise a DanStack struct with the default array
// capacity.
func New[T any]() DanStack[T] {
	st := DanStack[T]{0, make([]T, initialSize)}
	return st
}

// Push adds to the top of the stack, growing the array if necessary.
func (stack *DanStack[T]) Push(item T) {
	if stack.top < len(stack.items) {
		stack.items[stack.top] = item
	} else {
		stack.items = append(stack.items, item)
	}
	stack.top++
}

// Pop removes an item from the top of the stack and returns it.
// Pop will return an error if the stack is empty.
func (stack *DanStack[T]) Pop() (T, error) {
	if stack.top-1 < 0 {
		var t T
		return t, fmt.Errorf("stack underflow")
	}
	prev := stack.items[stack.top-1]
	stack.top--
	return prev, nil
}

// Empty returns true is there are no items in
// the stack.
func (stack DanStack[T]) Empty() bool {
	return stack.top < 1
}

package main

// Stack represents a LIFO stack of integers.
type Stack struct {
	items []int
}

// Push adds a value to the top of the stack
func (stack *Stack) Push(value int) {
	stack.items = append(stack.items, value)
}

// Pop removes and returns the value from the top of the stack
// Returns false if the stack is empty.
func (stack *Stack) Pop() (int, bool) {
	if len(stack.items) == 0 {
		return 0, false
	}

	top := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return top, true
}

// Peek returns the value at the top of the stack without removing it
// Returns false if the stack is empty.
func (stack *Stack) Peek() (int, bool) {
    if len(stack.items) == 0 {
        return 0, false
    }
    return stack.items[len(stack.items)-1], true
}

// IsEmpty checks if the stack is empty.
func (stack *Stack) IsEmpty() bool {
    return len(stack.items) == 0
}

// Length returns the number of elements in the stack.
func (stack *Stack) Length() int {
	return len(stack.items)
}

package structures

import "fmt"

// StackNode represents a single element in a stack.
// Each node contains an integer Value and pointers to the next nodes.
type StackNode struct {
    Value int   // The value stored in the node
    Next  *StackNode // Pointer to the next node in the list (nil if this is the last node)
}

// Stack represents a LIFO stack of integers.
type Stack struct {
	Top *StackNode
	Height int
}

// Push adds a value to the top of the stack
func (stack *Stack) Push(value int) {
	node := &StackNode{Value: value}
	if stack.Height == 0 {
		stack.Top = node
		stack.Height = 1
	} else {
		stack.Height += 1

		temp := stack.Top
		stack.Top = node
		stack.Top.Next = temp
	}
}

// Pop removes and returns the value from the top of the stack
// Returns false if the stack is empty.
func (stack *Stack) Pop() {
	if stack.Height != 0 {
		top := stack.Top.Next
		stack.Top = top
		stack.Height -= 1
	}
}

// Peek returns the value at the top of the stack without removing it
// Returns false if the stack is empty.
func (stack *Stack) Peek() (int, bool) {
    if stack.Height == 0 {
		return 0, false
	}

	return stack.Top.Value, true
}

// IsEmpty checks if the stack is empty.
func (stack *Stack) IsEmpty() bool {
    return stack.Height == 0
}

// Print prints all values in the stack.
func (list *Stack) Print() {
    current := list.Top
    for current != nil {
        fmt.Printf("%d -> ", current.Value)
        current = current.Next
    }
    fmt.Println("nil")
}

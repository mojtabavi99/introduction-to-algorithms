package structures

import (
	"fmt"
)

// BinarySearchTreeNode represents a single node in a binary search tree.
// Each node contains an integer Value and pointers to its left and right children.
type BinarySearchTreeNode struct {
	Value int                     // The value stored in the node
	Left  *BinarySearchTreeNode   // Pointer to the left child (contains smaller values)
	Right *BinarySearchTreeNode   // Pointer to the right child (contains larger values)
}

// BinarySearchTree represents a binary search tree (BST).
// The Root pointer refers to the topmost node in the tree.
type BinarySearchTree struct {
	Root *BinarySearchTreeNode // The root node of the BST
}

// Insert adds a new value to the binary search tree.
// Returns (true, parentValue) if insertion succeeds,
// or (false, 0) if the value already exists in the tree.
func (bst *BinarySearchTree) Insert(value int) (result bool, parent int) {
	node := &BinarySearchTreeNode{Value: value}

	// If the tree is empty, set the root.
	if bst.Root == nil {
		bst.Root = node
		return true, 0
	}

	current := bst.Root
	for {
		// If the value already exists, do not insert.
		if current.Value == node.Value {
			return false, 0
		}

		// Go to the left subtree if value is smaller.
		if node.Value < current.Value {
			if current.Left == nil {
				current.Left = node
				return true, current.Value
			}
			current = current.Left
		} else {
			// Go to the right subtree if value is larger.
			if current.Right == nil {
				current.Right = node
				return true, current.Value
			}
			current = current.Right
		}
	}
}

// Delete removes a value from the binary search tree.
// Returns true if the value was found and deleted, otherwise false.
func (bst *BinarySearchTree) Delete(value int) bool {
    var deleted bool
    bst.Root, deleted = deleteNode(bst.Root, value)
    return deleted
}

// Contains checks whether the given value exists in the tree.
// Returns true if found, otherwise false.
func (bst *BinarySearchTree) Contains(value int) bool {
	current := bst.Root
	for current != nil {
		if value == current.Value {
			return true
		}
		if value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return false
}

// InOrder performs an in-order traversal of the tree (Left → Root → Right)
// and prints all node values in ascending order.
func (bst *BinarySearchTree) InOrder() {
	if bst.Root != nil {
		bst.Root.InOrder()
	}
}

// PreOrder performs a pre-order traversal of the tree (Root → Left → Right)
// and prints all node values.
func (bst *BinarySearchTree) PreOrder() {
	if bst.Root != nil {
		bst.Root.PreOrder()
	}
}

// PrettyPrint prints the tree structure sideways for easier visualization.
func (bst *BinarySearchTree) PrettyPrint() {
	if bst.Root != nil {
		bst.Root.PrettyPrint()
	}
}

// ===== Node-level recursive helper methods =====

// deleteNode is a recursive helper that deletes a node with the given value
// and returns the updated subtree and whether deletion occurred.
func deleteNode(node *BinarySearchTreeNode, value int) (*BinarySearchTreeNode, bool) {
    if node == nil {
        return nil, false
    }

    if value < node.Value {
        node.Left, _ = deleteNode(node.Left, value)
        return node, true
    } else if value > node.Value {
        node.Right, _ = deleteNode(node.Right, value)
        return node, true
    }

    // value == node.Value → found the node to delete
    if node.Left == nil && node.Right == nil {
        // no children
        return nil, true
    } else if node.Left == nil {
        // only right child
        return node.Right, true
    } else if node.Right == nil {
        // only left child
        return node.Left, true
    } else {
        // two children
        successor := findMin(node.Right)
        node.Value = successor.Value
        node.Right, _ = deleteNode(node.Right, successor.Value)
        return node, true
    }
}

// findMin returns the node with the smallest value in a subtree.
func findMin(node *BinarySearchTreeNode) *BinarySearchTreeNode {
    current := node
    for current.Left != nil {
        current = current.Left
    }
    return current
}

// InOrder prints the node values in-order recursively.
func (node *BinarySearchTreeNode) InOrder() {
	if node == nil {
		return
	}
	node.Left.InOrder()
	fmt.Printf("%d ", node.Value)
	node.Right.InOrder()
}

// PreOrder prints the node values in pre-order recursively.
func (node *BinarySearchTreeNode) PreOrder() {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Value)
	node.Left.PreOrder()
	node.Right.PreOrder()
}

// PrettyPrint prints the subtree rooted at the current node
// in a visually structured, sideways format.
func (node *BinarySearchTreeNode) PrettyPrint() {
	printSideways(node, "", true)
}

// printSideways is a helper function that recursively prints
// the tree structure sideways, using ASCII branches for clarity.
func printSideways(node *BinarySearchTreeNode, prefix string, isTail bool) {
	if node == nil {
		return
	}
	printSideways(node.Right, prefix+"    ", false)
	fmt.Printf("%s", prefix)
	if isTail {
		fmt.Printf("└── ")
	} else {
		fmt.Printf("┌── ")
	}
	fmt.Printf("%d\n", node.Value)
	printSideways(node.Left, prefix+"    ", true)
}

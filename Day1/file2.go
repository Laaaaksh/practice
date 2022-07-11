// Question 2

package main

import (
	"fmt"
	"io"
	"os"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) insert(data int) *Tree {
	if t.root == nil {
		t.root = &TreeNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}
func (n *TreeNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &TreeNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &TreeNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func print(w io.Writer, node *TreeNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	tree := &Tree{}
	tree.insert(100)
	tree.insert(-20)
	tree.insert(-9)
	tree.insert(-5)
	tree.insert(-1)
	tree.insert(44)
	tree.insert(6)
	tree.insert(42)
	tree.insert(35)
	tree.insert(5)
	print(os.Stdout, tree.root, 0, 'M')
}

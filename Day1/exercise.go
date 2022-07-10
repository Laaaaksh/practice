// Question 1
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type matrix struct {
	rows int
	columns int
	ele [3][3] int

}

func (r * matrix) getRows() int{
	return r.rows
}
func (r * matrix) getColumn() int{
	return r.columns
}
func (r * matrix) setEle(){
	row:= r.rows
	col:=r.columns
	arr:=r.ele // declared an empty 2d array
	for i:=0; i<row; i++{
		for j:=0 ; j < col; j++{
			fmt.Print(arr[i][j], " ") // assigned values to it
		}
		fmt.Print("\n")
	}
}
func (r *matrix) addMatrix(){
	row:= r.rows
	col:=r.columns
	arr:=r.ele
	var mat = [3][3]int{{1,1,1},{1,1,1},{1,1,1}}
	for i:=0; i<row; i++{
		for j:=0 ; j < col; j++{
			arr[i][j] = arr[i][j] + mat[i][j]
		}
	}
	tempFunc := matrix{3,3,arr}
	tempFunc.setEle()
}
// Unable to print array as json
func (r * matrix) printJson(){
	arr:=r.ele
	convArr, err:= json.Marshal(arr)
	if err == nil{
		log.Fatal(err)
	}
	fmt.Println(convArr)

}
func main()  {
	fmt.Println("Five Functions")
	var one= [3][3] int {{1,2,3},{4,5,6},{7,8,9}}
	var emptyArr= [3][3] int{}
	firstFunc := matrix{3,3, emptyArr}
	secondFunc := matrix{3,3, one}
	fmt.Println(firstFunc.rows)
	fmt.Println(firstFunc.columns)
	secondFunc.setEle()
	fmt.Print("\n")
	secondFunc.addMatrix()
	secondFunc.printJson()
}
// Question one end here

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
	tree.insert(100).
		 insert(-20).
		 insert(-9).
		 insert(-5).
		 insert(-1).
		 insert(44).
		 insert(6).
		 insert(42).
		 insert(35).
		 insert(5).
	print(os.Stdout, tree.root, 0, 'M')
 }
// Question 2 ends here

// Question 3
 package main

 import "fmt"

 type salary interface {
 	calCi() int
 }
 type employee struct {
 	daysWorked int
 	wage int
 }

 func (e employee) calCi() int {
	 return e.daysWorked * e.wage
 }
 func calculate(i salary){
	println(i.calCi())
 }
 func main(){
	 fmt.Println("Third question")
	 fullTime := employee{23, 500}
	 conTract := employee{24, 100}
	 freeLancer := employee{28, 10}

	 calculate(fullTime)
	 calculate(conTract)
	 calculate(freeLancer)
 }

 // Question 3 ends here


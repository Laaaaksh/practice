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
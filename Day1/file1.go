package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Matrix struct {
	rows    int `json:"rows"`
	columns int `json:"columns"`
	ele     [][]int `json:"ele"`
}
// Get rows
func (r *Matrix) getRows() int {
	return r.rows
}
// Get columns
func (r *Matrix) getColumn() int {
	return r.columns
}
// Set values in matrix
func (r *Matrix) setEle(val int) {
	row := r.rows
	col := r.columns
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			r.ele[i][j] = val
		}
	}
}
// Add matrix
func (r *Matrix) addMatrix(mat [][]int) {
	row := r.rows
	col := r.columns
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			r.ele[i][j] = r.ele[i][j] + mat[i][j]
		}
	}
}

// Print Json Matrix
func (r *Matrix) printJson() {
	val, err := json.Marshal(r.ele)
	if err!=nil {
		log.Default()
	}
	fmt.Println(string(val))

}
func main() {
	var one Matrix
	one.rows = 3
	one.columns = 3
	one.ele = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//var mat = [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}} // here declare the type of matrix you wanna add
	//one.setEle(5)
	//one.addMatrix(mat)
	one.printJson()


}

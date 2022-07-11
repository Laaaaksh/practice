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

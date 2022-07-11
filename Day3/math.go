package main

import "fmt"
func add (a int, b int) int{
	return a + b
}
func main()  {
	fmt.Println("Day 3")
	var sum  = add(1,2)
	fmt.Println(sum)
}

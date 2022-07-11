package main

import (
	"fmt"
	"strings"
	"sync"
)
func calculate(string2 string){
	fmt.Println(string2)
	for _,val := range string2{
		ch :=strings.Count(string2,string(val))
		fmt.Println(string(val), " - ", ch)
	}
}

func process(i string, vr * sync.WaitGroup, ms *sync.Mutex) {
	ms.Lock()
	calculate(i)
	ms.Unlock()
	vr.Done()
}
func main(){
	fmt.Println("Question one")
	str := []string{"Hi", "hey", "laksh", "ball", "razorpay"}
	var ch sync.WaitGroup
	var ms sync.Mutex
	for _,bt := range str {
		ch.Add(1)
		go process(bt, &ch, &ms)
	}
	ch.Wait()
	fmt.Println("Printed count of alphabets in all the strings")
}
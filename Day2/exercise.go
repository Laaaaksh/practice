//package main
//
//import (
//	"fmt"
//	"strings"
//	"sync"
//)
//func calculate(string2 string){
//	fmt.Println(string2)
//	for _,val := range string2{
//		ch :=strings.Count(string2,string(val))
//		fmt.Println(string(val), " - ", ch)
//	}
//}
//
//func process(i string, vr * sync.WaitGroup, ms *sync.Mutex) {
//	ms.Lock()
//	calculate(i)
//	ms.Unlock()
//	vr.Done()
//}
//func main(){
//  fmt.Println("Question one")
//	str := []string{"Hi", "hey", "laksh", "ball", "razorpay"}
//	var ch sync.WaitGroup
//	var ms sync.Mutex
//	for _,bt := range str {
//		ch.Add(1)
//		go process(bt, &ch, &ms)
//	}
//	ch.Wait()
//	fmt.Println("Printed count of alphabets in all the strings")
//
//
//}

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"sync"
//	"time"
//)
//
//func review(i int, sc * sync.WaitGroup, sum *int, mt *sync.Mutex)  {
//	rand.Seed(time.Now().UnixNano())
//	mt.Lock()
//	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
//	pt := rand.Intn(10)
//	*sum += pt
//	fmt.Println("Student : " ,i, "Rating is", pt)
//	mt.Unlock()
//	sc.Done()
//}
//func main(){
//	fmt.Println("Question 2")
//	num:=200
//	var sum int
//	var sc sync.WaitGroup
//	var mt sync.Mutex
//	for i:=0 ; i < num ; i++{
//		sc.Add(1)
//		go review(i, &sc,&sum,&mt)
//	}
//	sc.Wait()
//	fmt.Println("Average Rating : ", sum/200 )
//}


package main

import (
	"fmt"
	"math/rand"
	"sync"
)
type account struct {
	balance int
}

func deposit (bal *int, add int){
		*bal = *bal + add
}

func withdrawl (bal *int, sub int){
	if *bal < 0{
		fmt.Println("Less that 0 account balance")
	} else {
		*bal = *bal - sub
	}
}
func main(){
	fmt.Println("Question 3")
	bal := 500
	var vs sync.Mutex
	for i:=0 ; i < 10 ; i++{
		sub := rand.Intn(100)
		add:= rand.Intn(100)
		vs.Lock()
		go deposit(&bal, add)
		vs.Unlock()
		vs.Lock()
		go withdrawl(&bal, sub)
		vs.Unlock()
		fmt.Println("Remaining Balance", bal)
	}
	fmt.Println("Final Balance", bal)

// need to make change here as the updated balance equation should be placed in the critical section
}
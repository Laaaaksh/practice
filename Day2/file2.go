package main
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
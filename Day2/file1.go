package main

import (
	"fmt"
	"sync"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var f = make(map[string]int, 26)

func countFrequency(word string) {
	m.Lock()
	for i := 0; i < len(word); i++ {
		f[string(word[i])] += 1
	}
	m.Unlock()
	wg.Done()
}

func main() {
	strings := []string{"quick", "brown", "fox", "lazy", "dog"}
	for j := range strings {
		wg.Add(1)
		go countFrequency(strings[j])
	}
	wg.Wait()
	fmt.Println(f)
	// checking
}

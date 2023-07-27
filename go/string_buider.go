package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	var str strings.Builder
	wg.Add(2)
	go stringAppend(&str)
	go stringAppend(&str)
	wg.Wait()
	fmt.Println(len(str.String()))

}

func stringAppend(str *strings.Builder) {

	for i := 0; i < 1000; i++ {
		mu.Lock()
		str.WriteString("a")
		mu.Unlock()
	}
	wg.Done()
}

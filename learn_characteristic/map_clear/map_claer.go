package main

import (
	"fmt"
	"runtime"
)

func main() {

	memStats := &runtime.MemStats{}

	v := struct{}{}
	i := 0
	m := make(map[int]struct{})
	for {
		runtime.ReadMemStats(memStats)
		fmt.Printf("Map begin memory: %d bytes\n", memStats.Alloc)
		fmt.Println()
		BuildMap(i, m, v)
		for k, _ := range m {
			delete(m, k)
		}
		i++
		runtime.ReadMemStats(memStats)
		fmt.Printf("Map last memory: %d bytes\n", memStats.Alloc)
		fmt.Println()
		if i == 10 {
			runtime.GC()
			fmt.Println("手动gc")
		}
	}

}

func BuildMap(index int, m map[int]struct{}, v struct{}) {
	for i := index; i < (index+1)*10000000; i++ {
		m[i] = v
	}
}

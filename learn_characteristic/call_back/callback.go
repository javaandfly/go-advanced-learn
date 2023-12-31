package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []string{"€hello", "$malong", " gaoxiao"}
	sort.SliceStable(s1, func(i, j int) bool {
		// 按字节大小顺序降序排序
		return len(s1[i]) > len(s1[j])
	})

	fmt.Println(s1)
}

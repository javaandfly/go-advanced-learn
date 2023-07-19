package main

import "fmt"

func main() {
	fmt.Println(replaceSpace("a b c"))
}

func replaceSpace(s string) string {
	var str string
	for _, v := range s {
		if v == ' ' {
			str += "%20"
			continue
		}
		str += string(v)

	}
	return str
}

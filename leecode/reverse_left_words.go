package main

import "fmt"

func main() {
	fmt.Println(reverseLeftWords("abcd", 2))
}

func reverseLeftWords(s string, n int) string {
	str := s[n:]
	str2 := s[:n]

	return str + str2
}

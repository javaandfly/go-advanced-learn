package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aa"
	fmt.Println(&str)
	var str2 strings.Builder
	fmt.Printf("%p\n", &str2)
	str2.WriteString("addaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa19908979231231231312312423214215115151aaaaaaaaaaaaaaa" + str)
	fmt.Printf("%p\n", &str2)
}

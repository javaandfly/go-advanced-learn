package main

import (
	"bytes"
	"fmt"
)

func main() {

	BenchmarkStringOperation3()
}

func BenchmarkStringOperation3() {

	strBuf := bytes.NewBufferString("11")

	strBuf.WriteString("golang")
	fmt.Println(strBuf)

}

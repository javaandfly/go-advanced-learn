package main

import (
	"fmt"
)

func main() {
	n := 2000
	// tsGap := time.Now().Unix() - stake.StartTime

	pointForSecond := 24 * 60 * 60 / n
	fmt.Println(pointForSecond)
}

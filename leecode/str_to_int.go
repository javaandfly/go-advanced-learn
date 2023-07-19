package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(StrToInt("9223372036854775808"))
}

func StrToInt(str string) int {
	for _, v := range str {
		if v == ' ' {
			str = str[1:]
			continue
		}
		break
	}
	var isBurdenNumber bool
	length := len(str)
	var n int
	for _, v := range str {
		if v == '-' {
			str = str[1:]
			isBurdenNumber = true
			length--
			n++
			continue
		}
		if v == '+' {
			str = str[1:]
			length--
			n++
			continue
		}
	}
	if n >= 2 {
		return 0
	}

	for i, v := range str {
		if v > 57 || v < 48 {
			str = str[:i]
			break
		}
	}
	length = len(str)

	var sum int
	for _, v := range str {

		lengthZero := length - 1
		threshold := 1
		for {
			if lengthZero != 0 {
				threshold = threshold * 10
				lengthZero--
				continue
			}
			break
		}
		sum = sum + (int(v)-48)*threshold
		length--
	}
	if isBurdenNumber {
		sum = -sum
	}
	if sum > math.MaxInt32 {
		return math.MaxInt32
	}
	if sum < math.MinInt32 {
		return math.MinInt32
	}
	return sum
}

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	min := int(1e9 + 7)
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		if t < min {
			min = t
		}
	}

	fmt.Println(min)
}

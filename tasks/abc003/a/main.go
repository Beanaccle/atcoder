package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := 0
	for i := 0; i < n; i++ {
		x += (i + 1) * 10000
	}

	fmt.Println(x / n)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	c := make([][]string, 4)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			var s string
			fmt.Scan(&s)
			c[i] = append([]string{s}, c[i]...)
		}
	}

	for i := 3; i >= 0; i-- {
		fmt.Println(strings.Join(c[i], " "))
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var w string
	fmt.Scan(&w)

	var x []string
	z := "aiueo"

	for _, y := range strings.Split(w, "") {
		if strings.Index(z, y) == -1 {
			x = append(x, y)
		}
	}
	fmt.Println(strings.Join(x, ""))
}

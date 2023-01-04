package main

import (
	"fmt"
)

var x, y int

func main() {
	fmt.Scan(&x, &y)
	fmt.Println(y / x)
}

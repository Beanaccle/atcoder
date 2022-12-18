package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x > y {
		fmt.Println(x)
	} else {
		fmt.Println(y)
	}
}

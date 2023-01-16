package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	ans := []string{"-1", "-1", "-1"}
loop:
	for a := 0; a <= n; a++ {
		if a == n {
			if a*2 == m {
				ans[0] = strconv.Itoa(a)
				ans[1] = "0"
				ans[2] = "0"
			}
		} else {
			for b := 0; b <= 1; b++ {
				c := n - a - b
				if a*2+b*3+c*4 == m {
					ans[0] = strconv.Itoa(a)
					ans[1] = strconv.Itoa(b)
					ans[2] = strconv.Itoa(c)
					break loop
				}
			}
		}
	}

	fmt.Fprintln(writer, strings.Join(ans, " "))
}

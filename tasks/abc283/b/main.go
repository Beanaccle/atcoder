package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	aList := make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(r, &a)
		aList[i] = a
	}

	var q int
	fmt.Fscan(r, &q)
	for i := 0; i < q; i++ {
		var s, k int
		fmt.Fscan(r, &s, &k)
		if s == 1 {
			var x int
			fmt.Fscan(r, &x)
			aList[k-1] = x
		} else {
			fmt.Fprintln(w, aList[k-1])
		}
	}
}

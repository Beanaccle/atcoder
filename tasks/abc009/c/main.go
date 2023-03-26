package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var N, K int
	var S string
	fmt.Fscan(reader, &N, &K, &S)

	s := S
	for i := 0; i < N; i++ {
		j := i
		for k := i + 1; k < N; k++ {
			if s[k] < s[j] && count(S, swap(s, i, k)) <= K {
				j = k
			}
		}
		if j > i {
			s = swap(s, i, j)
		}
	}

	fmt.Fprintln(writer, s)
}

func count(a, b string) int {
	cnt := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt
}

func swap(s string, i, j int) string {
	return s[:i] + s[j:j+1] + s[i+1:j] + s[i:i+1] + s[j+1:]
}

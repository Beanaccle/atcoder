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

	var N, M int
	fmt.Fscan(reader, &N, &M)

	S := make([]string, N)
	for i := 0; i < N; i++ {
		var s string
		fmt.Fscan(reader, &s)
		S[i] = s[3:]
	}

	T := make([]string, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(reader, &T[i])
	}

	cnt := 0
	for _, s := range S {
		for _, t := range T {
			if s == t {
				cnt++
				break
			}
		}
	}

	fmt.Fprintln(writer, cnt)
}

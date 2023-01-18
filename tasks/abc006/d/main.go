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

	var n int
	fmt.Fscan(reader, &n)

	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &cards[i])
	}

	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = 10e9
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			if dp[j] < cards[i] && cards[i] < dp[j+1] {
				dp[j+1] = cards[i]
			}
		}
	}

	ans := 0
	for _, c := range dp {
		if c > n {
			ans++
		}
	}

	fmt.Fprintln(writer, ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

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

	fmt.Fprintln(writer, trib(n))
}

func trib(n int) int {
	if n < 3 {
		return 0
	} else {
		dp := make([]int, n+1)
		dp[0] = 0
		dp[1] = 0
		dp[2] = 0
		dp[3] = 1
		for i := 4; i <= n; i++ {
			dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 10007
		}
		return dp[n]
	}
}

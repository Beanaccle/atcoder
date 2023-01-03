package main

import (
	"fmt"
)

var r, g, b int

func main() {
	fmt.Scan(&r, &g, &b)

	ttl := r + g + b
	max := 1000
	big := 1e9 + 7
	dp := make([][]int, max)

	// 初期化
	for i := 0; i < max; i++ {
		dp[i] = make([]int, ttl+1)
		for j := range dp[i] {
			if j == ttl {
				dp[i][j] = 0
			} else {
				dp[i][j] = int(big)
			}
		}
	}

	ans := int(big)
	for i := 1; i < max; i++ {
		for j := 0; j < ttl; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]+cost(i, j))
		}
		ans = min(ans, dp[i][0])
	}

	fmt.Println(ans)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func cost(pos, remain int) int {
	if remain >= g+b {
		// 残数がg+bより多い -> r
		return abs(350 - pos)
	} else if remain >= b {
		// 残数がbより多い -> g
		return abs(450 - pos)
	}
	return abs(550 - pos)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

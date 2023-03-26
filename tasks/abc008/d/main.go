package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type XY struct {
	minx, miny, maxx, maxy int
}

var w, h, n int
var x, y [30]int
var dp = map[XY]int{}

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &w, &h, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
		x[i]--
		y[i]--
	}

	fmt.Fprintln(writer, dfs(0, 0, w, h))
}

func dfs(minx, miny, maxx, maxy int) int {
	xy := XY{minx: minx, miny: miny, maxx: maxx, maxy: maxy}
	if _, ok := dp[xy]; ok {
		return dp[xy]
	}
	res := 0
	for i := 0; i < n; i++ {
		if minx <= x[i] && x[i] < maxx && miny <= y[i] && y[i] < maxy {
			dfsa := dfs(minx, miny, x[i], y[i])
			dfsb := dfs(x[i]+1, miny, maxx, y[i])
			dfsc := dfs(minx, y[i]+1, x[i], maxy)
			dfsd := dfs(x[i]+1, y[i]+1, maxx, maxy)
			res = max(res, dfsa+dfsb+dfsc+dfsd+(maxx-minx)+(maxy-miny)-1)
		}
	}
	dp[xy] = res
	return dp[xy]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

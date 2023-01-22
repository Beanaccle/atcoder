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

	var a, b int
	fmt.Fscan(reader, &a, &b)

	aCnt := count(digit(a - 1))
	bCnt := count(digit(b))

	fmt.Fprintln(writer, bCnt-aCnt)
}

func count(d []int) int {
	l := len(d)
	dp := make([][][]int, l+1)
	for i := 0; i <= l; i++ {
		dp[i] = make([][]int, 2)
		for smaller := 0; smaller < 2; smaller++ {
			dp[i][smaller] = make([]int, 2)
			for j := 0; j < 2; j++ {
				dp[i][smaller][j] = 0
			}
		}
	}

	dp[0][0][0] = 1
	for i := 0; i < l; i++ {
		for smaller := 0; smaller < 2; smaller++ {
			for j := 0; j < 2; j++ {
				y := d[i]
				if smaller == 1 {
					y = 9
				}
				for x := 0; x <= y; x++ {
					s := 0
					if smaller == 1 || x < d[i] {
						s = 1
					}
					jj := j
					if x == 4 || x == 9 {
						jj = 1
					}

					dp[i+1][s][jj] += dp[i][smaller][j]
				}
			}
		}
	}

	return dp[l][0][1] + dp[l][1][1]
}

func digit(n int) []int {
	var ret []int
	for n != 0 {
		ret = append([]int{n % 10}, ret...)
		n /= 10
	}
	return ret
}

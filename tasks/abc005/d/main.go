package main

import (
	"fmt"
)

var grid [][]int

func main() {
	var n int
	fmt.Scan(&n)

	grid = make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			var d int
			fmt.Scan(&d)
			grid[i][j] = d
		}
	}

	var q int
	fmt.Scan(&q)

	ansList := make([]int, n*n+1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 起点 i, j

			for x := 1; x <= (n - i); x++ {
				for y := 1; y <= (n - j); y++ {
					// 長方形 x * y
					ansList[x*y] = max(ansList[x*y], calc(i, j, x, y))
				}
			}
		}
	}

	for r := 0; r < q; r++ {
		var p int
		fmt.Scan(&p)

		max := 0
		for _, ans := range ansList[:p+1] {
			if max < ans {
				max = ans
			}
		}
		fmt.Println(max)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calc(i, j, x, y int) int {
	var tmp int

	for xx := 0; xx < x; xx++ {
		for yy := 0; yy < y; yy++ {
			tmp += grid[i+xx][j+yy]
		}
	}

	return tmp
}

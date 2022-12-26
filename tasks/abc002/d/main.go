package main

import (
	"fmt"
)

func main() {
	var n, m, x, y int
	fmt.Scan(&n, &m)

	xy := make([][]int, n)
	for i := range xy {
		xy[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&x, &y)
		xy[x-1][y-1] = 1
		xy[y-1][x-1] = 1
	}

	groups := make([][]int, n)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if x != y && xy[x][y] == 1 {
				group := []int{x, y}

				for z := 0; z < n; z++ {
					flg := true
					for _, g := range group {
						if xy[z][g] != 1 {
							flg = false
						}
					}
					if flg {
						group = append(group, z)
					}
				}

				groups[x] = group
			}
		}
	}

	max := 1
	for _, group := range groups {
		if max < len(group) {
			max = len(group)
		}
	}

	fmt.Println(max)
}

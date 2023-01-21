package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var r, c int
	fmt.Fscan(reader, &r, &c)
	var sy, sx int
	fmt.Fscan(reader, &sy, &sx)
	sy--
	sx--
	var gy, gx int
	fmt.Fscan(reader, &gy, &gx)
	gy--
	gx--

	grid := make([][]string, r)
	for i := 0; i < r; i++ {
		var str string
		fmt.Fscan(reader, &str)
		grid[i] = make([]string, c)
		for j, s := range str {
			grid[i][j] = string(s)
		}
	}

	grid[sy][sx] = "0"
	vy := [4]int{0, 1, 0, -1}
	vx := [4]int{1, 0, -1, 0}
	var que [][]int
	que = append(que, []int{sy, sx})
	for len(que) > 0 {
		n := que[0]
		que = que[1:]

		for i := 0; i < 4; i++ {
			ny := n[0] + vy[i]
			nx := n[1] + vx[i]

			if grid[ny][nx] == "." {
				val, _ := strconv.Atoi(grid[n[0]][n[1]])
				grid[ny][nx] = strconv.Itoa(val + 1)
				que = append(que, []int{ny, nx})
			}
		}
	}

	fmt.Fprintln(writer, grid[gy][gx])
}

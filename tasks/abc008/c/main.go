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

	coins := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &coins[i])
	}

	var cnt float64
	for i := 0; i < n; i++ {
		num := 0
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if coins[i]%coins[j] == 0 {
				num++
			}
		}

		cnt += float64((num+2)/2) / float64(num+1)
	}

	fmt.Fprintf(writer, "%.12f\n", cnt)
}

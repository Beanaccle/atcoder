package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	t0 := make([][]int, n)
	for i := 0; i < n; i++ {
		var r string
		fmt.Scan(&r)

		arr0 := strings.Split(r, "-")
		f0, _ := strconv.ParseFloat(arr0[0], 64)
		a0 := int(math.Floor(f0/5) * 5)
		f1, _ := strconv.ParseFloat(arr0[1], 64)
		a1 := int(math.Ceil(f1/5) * 5)
		if (a1 % 100) == 60 {
			a1 += 40
		}
		t0[i] = []int{a0, a1}
	}

	sort.Slice(t0, func(i, j int) bool { return t0[i][0] < t0[j][0] })

	t := t0[0]
	for i := 0; i < n; i++ {
		if t[1] >= t0[i][0] && t0[i][1] > t[1] {
			t[1] = t0[i][1]
		} else if t[1] < t0[i][0] {
			fmt.Printf("%04d-%04d\n", t[0], t[1])
			t = t0[i]
		}
	}
	fmt.Printf("%04d-%04d\n", t[0], t[1])
}

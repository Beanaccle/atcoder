package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var r int
	var rates []int
	for i := 0; i < n; i++ {
		fmt.Scan(&r)
		rates = append(rates, r)
	}

	sort.Ints(rates)

	var rate float64
	for i := (n - k); i < n; i++ {
		rate = (rate + float64(rates[i])) / 2
	}

	fmt.Println(rate)
}

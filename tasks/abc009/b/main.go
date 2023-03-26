package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	list := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &list[i])
	}

	uniqList := uniq(list)
	sort.Ints(uniqList)

	fmt.Fprintln(writer, uniqList[len(uniqList)-2])
}

func uniq(arr []int) []int {
	res := []int{}
	for _, v := range arr {
		if !contains(res, v) {
			res = append(res, v)
		}
	}
	return res
}

func contains(arr []int, v int) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	win := true
	ss := strings.Split(s, "")
	tt := strings.Split(t, "")
	for i := 0; i < len(s); i++ {
		if ss[i] == tt[i] {
			continue
		} else if ss[i] == "@" && strings.Contains("atcoder@", tt[i]) {
			continue
		} else if tt[i] == "@" && strings.Contains("atcoder@", ss[i]) {
			continue
		} else {
			win = false
			break
		}
	}

	if win {
		fmt.Println("You can win")
	} else {
		fmt.Println("You will lose")
	}
}

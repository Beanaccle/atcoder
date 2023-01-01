package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	cards := []int{1, 2, 3, 4, 5, 6}

	for i := 0; i < (n % 30); i++ {
		mod := i % 5
		cards[mod], cards[mod+1] = cards[mod+1], cards[mod]
	}

	for _, card := range cards {
		fmt.Print(card)
	}
	fmt.Print("\n")
}

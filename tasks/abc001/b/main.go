package main

import (
	"fmt"
)

func main() {
	var m, vv float32
	fmt.Scan(&m)
	km := m / 1000

	if km < 0.1 {
		vv = 0
	} else if 0.1 <= km && km <= 5 {
		vv = km * 10
	} else if 6 <= km && km <= 30 {
		vv = km + 50
	} else if 35 <= km && km <= 70 {
		vv = (km-30)/5 + 80
	} else {
		vv = 89
	}

	fmt.Printf("%02.0f\n", vv)
}

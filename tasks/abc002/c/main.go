package main

import (
	"fmt"
	"math"
)

func main() {
	var xa, ya, xb, yb, xc, yc float64
	fmt.Scan(&xa, &ya, &xb, &yb, &xc, &yc)
	area := math.Abs(((xb-xa)*(yc-ya) - (yb-ya)*(xc-xa)) / 2)
	fmt.Printf("%.1f\n", area)
}

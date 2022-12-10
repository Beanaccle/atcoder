package main

import (
	"fmt"
	"math"
)

func main() {
	var deg, dis int
	var dir string
	fmt.Scan(&deg, &dis)

	w := windPower(dis)

	if w == 0 {
		dir = "C"
	} else {
		dir = windDirection(deg)
	}

	fmt.Println(dir, w)
}

func windDirection(deg int) string {
	d := [...]string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW", "N"}
	i := ((deg*10 + 1125) / 2250) % 16
	return d[i]
}

func round(x float64, y int) float64 {
	shift := math.Pow(10, float64(y))
	return math.Floor(x*shift+.5) / shift
}

func windPower(dis int) int {
	wind := round(float64(dis)/60, 1)

	if wind <= 0.2 {
		return 0
	} else if wind <= 1.5 {
		return 1
	} else if wind <= 3.3 {
		return 2
	} else if wind <= 5.4 {
		return 3
	} else if wind <= 7.9 {
		return 4
	} else if wind <= 10.7 {
		return 5
	} else if wind <= 13.8 {
		return 6
	} else if wind <= 17.1 {
		return 7
	} else if wind <= 20.7 {
		return 8
	} else if wind <= 24.4 {
		return 9
	} else if wind <= 28.4 {
		return 10
	} else if wind <= 32.6 {
		return 11
	} else {
		return 12
	}
}

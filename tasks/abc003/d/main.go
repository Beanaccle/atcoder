package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var r, c, x, y, d, l int
	fmt.Scan(&r, &c)
	fmt.Scan(&x, &y)
	fmt.Scan(&d, &l)

	// x * yの範囲で、すべての置き方NからABCDの和集合を引いた順列nを求める
	//  A   B   C   D
	// --- +++ -++ ++-
	// +++ +++ -++ ++-
	// +++ --- -++ ++-

	// n = N - (A || B || C || D)
	// 包除原理
	// n = N - (A + B + C + D - (A && B) - (A && C) - (A && D) - (B && C) - (B && D) - (C && D)
	//       + (A && B && C) + (A && B && D) + (A && C && D) + (B && C && D)
	//       - (A && B && C && D))
	// n = N - A - B - C - D + (A && B) + (A && C) + (A && D) + (B && C) + (B && D) + (C && D)
	//       - (A && B && C) - (A && B && D) - (A && C && D) - (B && C && D)
	//       + (A && B && C && D))

	n := big.NewInt(0)
	// 1 << 4 = 16回でABCDのパターンを網羅できる
	for i := 0; i < (1 << 4); i++ {
		xx := x
		yy := y
		bitCount := 0

		// A(上辺に接しないパターン)
		// ???1
		if (i & 1) == 1 {
			bitCount++
			yy--
		}
		// B(下辺に接しないパターン)
		// ??1?
		if (i >> 1 & 1) == 1 {
			bitCount++
			yy--
		}
		// C(左辺に接しないパターン)
		// ?1??
		if (i >> 2 & 1) == 1 {
			bitCount++
			xx--
		}
		// D(右辺に接しないパターン)
		// 1???
		if (i >> 3 & 1) == 1 {
			bitCount++
			xx--
		}

		// 0以下は区画にならない
		if xx <= 0 || yy <= 0 {
			continue
		}

		// デスクとラックの数より区画が少ない
		if (d + l) > (xx * yy) {
			continue
		}

		// xx * yy区画からデスクとラックと空(デスクとラック以外)の順列
		// (xx * yy)! / d! * l! * (xx * yy - (d + l))!
		v := new(big.Int).Div(
			factorial(xx*yy),
			new(big.Int).Mul(
				new(big.Int).Mul(factorial(d), factorial(l)),
				factorial(xx*yy-d-l)))

		if (bitCount % 2) == 0 {
			// iの２進数の1の数が偶数もしくは0の場合
			// ４辺が接している集合もしくは、2辺か4辺が接していない集合の順列を足す
			// N + (A && B) + (A && C) + (A && D) + (B && C) + (B && D) + (C && D) + (A && B && C && D)
			n = new(big.Int).Add(n, v)
		} else {
			// iの２進数の1の数が奇数の場合
			// 1辺か3辺が接していない集合の順列を引く
			// - A - B - C - D - (A && B && C) - (A && B && D) - (A && C && D) - (B && C && D)
			n = new(big.Int).Sub(n, v)
		}
	}

	diff := (r - x + 1) * (c - y + 1)
	nn := new(big.Int).Mul(n, big.NewInt(int64(diff)))

	mod := math.Pow(10, 9) + 7
	ans := new(big.Int).Mod(nn, big.NewInt(int64(mod)))
	fmt.Println(ans)
}

// 階乗
func factorial(n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	} else {
		return new(big.Int).Mul(factorial(n-1), big.NewInt(int64(n)))
	}
}

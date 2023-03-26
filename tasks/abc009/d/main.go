package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var K, M int
	fmt.Fscan(reader, &K, &M)

	A := make([]int, K)
	for i := 0; i < K; i++ {
		fmt.Fscan(reader, &A[i])
	}

	ans := 0
	if K >= M {
		ans = A[M-1]
	} else {
		C := make([]int, K)
		for i := 0; i < K; i++ {
			fmt.Fscan(reader, &C[i])
		}

		mat := make([][]int, K)
		for i := 0; i < K; i++ {
			mat[i] = make([]int, K)
			mat[0][i] = C[i]
			if i != 0 {
				mat[i][i-1] = math.MaxUint32
			}
		}

		mat = matrixPower(mat, M-K)
		for i := 0; i < K; i++ {
			ans ^= mat[0][i] & A[K-1-i]
		}
	}

	fmt.Fprintln(writer, ans)
}

func matrixMultiply(a, b [][]int) [][]int {
	mat := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		mat[i] = make([]int, len(b[0]))
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				mat[i][j] ^= a[i][k] & b[k][j]
			}
		}
	}
	return mat
}

func matrixPower(a [][]int, n int) [][]int {
	mat := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		mat[i] = make([]int, len(a))
		mat[i][i] = math.MaxUint32
	}
	for n > 0 {
		if n&1 == 1 {
			mat = matrixMultiply(mat, a)
		}
		a = matrixMultiply(a, a)
		n >>= 1
	}
	return mat
}

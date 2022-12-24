package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var s string
	fmt.Fscan(r, &s)

	ans := len(strings.Split(s, "00")) - 1
	for _, s := range strings.Split(s, "00") {
		ans += len(s)
	}
	fmt.Fprintln(w, ans)
}

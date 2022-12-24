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

	ans := "Yes"
	tmp := ""
	for _, str := range strings.Split(s, "") {
		if str == "(" {
			continue
		} else if str == ")" {
			tmp = ""
		} else if strings.Contains(tmp, str) {
			ans = "No"
			break
		} else {
			tmp += str
		}
	}

	fmt.Fprintln(w, ans)
}

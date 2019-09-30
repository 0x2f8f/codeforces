package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("in.txt")
	if err == nil {
		os.Stdin, err = os.OpenFile("in.txt", os.O_RDONLY, 0666)
	}

	var n, m, a, c, d uint64
	fmt.Scan(&n, &m, &a)

	c = n / a
	if (n%a > 0) {
		c++
	}

	d = m / a
	if (m%a>0) {
		d++
	}
	fmt.Println(c*d)
}

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

	var n int
	fmt.Scan(&n)
	if (n%3 == 0) {
		l := n/3
		fmt.Println(l-1, l, l+1)
	} else {
		fmt.Print(-1)
	}
}

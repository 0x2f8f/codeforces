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

	var n, a int
	j := 0
	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		fmt.Scan(&a)
		if a == 1 {
			j++
		}
	}

	if j>0 {
		fmt.Println("HARD")
	} else {
		fmt.Println("EASY")
	}
}
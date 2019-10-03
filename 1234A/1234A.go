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

	var n, m, a, s int

	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		fmt.Scan(&m)
		s = 0
		for j:=0; j<m; j++ {
			fmt.Scan(&a)
			s+=a
		}
		if (s%m==0) {
			fmt.Println(s/m)
		} else {
			fmt.Println(s/m+1)
		}

	}
}

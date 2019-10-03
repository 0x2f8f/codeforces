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

	var n,n1,n2,n3 int
	r := 0

	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		fmt.Scan(&n1, &n2, &n3)
		if ( (n1+n2+n3)>1 ) {
			r++
		}

	}
	fmt.Println(r)
}

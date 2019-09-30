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
	var s string
	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		fmt.Scan(&s)
		if (len(s)>10) {
			fmt.Printf("%s%d%s\n", string(s[0]),(len(s)-2),string(s[len(s)-1]));
		} else {
			fmt.Println(s)
		}
	}
}

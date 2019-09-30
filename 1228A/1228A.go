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

	var l, r int
	b := true

	fmt.Scan(&l, &r)

	for i:=l; i<=r; i++ {
		n:=i
		a := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		b = true
		for ; n>0; {
			k :=n%10
			a[k]++
			if (a[k]==2){
				b=false
				break;
			}
			n/=10
		}
		if (b) {
			fmt.Println(i)
			break;
		}
	}
	if (!b) {
		fmt.Println(-1)
	}

}
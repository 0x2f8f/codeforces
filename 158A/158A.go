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

	var n,k,v,r int
	kv:=0

	fmt.Scan(&n, &k)
	for i:=1; i<=n; i++ {
		fmt.Scan(&v)
		if(v>=kv && v>0) {
			r++
		}
		if (i==k) {
			kv=v
		}
		if (i>k && kv==0) {
			break;
		}
	}
	fmt.Println(r)
}

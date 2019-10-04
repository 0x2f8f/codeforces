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

	var a1, a2, k1, k2, n int
	fmt.Scan(&a1, &a2, &k1, &k2, &n)

	var s1, s2, r, ost, kol int
	if (n > k1) {
		//команда а
		r = n/k1
		if (r>a1) {
			kol = a1
		} else {
			kol = r
		}
		ost = n - kol*k1

		//переход к команде б
		if (ost > k2) {

		} else {

		}

	} else {
		kol = 0
		ost = n

		//переход к команде б
	}

}

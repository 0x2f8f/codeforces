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

	var a1, a2, k1, k2, n, max, min int
	fmt.Scan(&a1, &a2, &k1, &k2, &n)

	//ищем max
	kol1, ost := cards(a1, k1, n)
	kol2, ost := cards(a2, k2, ost)
	sum1 := kol1 + kol2

	kol3, ost := cards(a2, k2, n)
	kol4, ost := cards(a1, k1, ost)
	sum2 := kol3 + kol4

	if (sum1 > sum2) {
		max = sum1
	} else {
		max = sum2
	}

	//ищем min
	if (k1>1) {
		k1--
	}
	if (k2>1) {
		k2--
	}

	min = n - (k1*a1+k2*a2)
	if (min<0) {
		min = 0;
	}

	fmt.Println(min, max)
}

func cards(a int, k int, n int) (int, int) {
	var kol, ost int
	if (n >= k) {
		r := n/k
		if (r>a) {
			kol = a
		} else {
			kol = r
		}
		ost = n - kol*k
	} else {
		kol = 0
		ost = n
	}
	return kol, ost
}
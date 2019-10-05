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

	var a1, a2, k1, k2, n, max, min, k1tmp, k2tmp int64
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
		k1tmp = k1-1
	} else {
		k1tmp = k1
	}

	if (k2>1) {
		k2tmp = k2-1
	} else {
		k2tmp = k2
	}

	min = n - (k1tmp*a1+k2tmp*a2)
	if (min<=0) {
		min = 0
		if (k1>k2) {
			minIter(&a1, &k1, &n, &min)
			minIter(&a2, &k2, &n, &min)
		} else {
			minIter(&a2, &k2, &n, &min)
			minIter(&a1, &k1, &n, &min)
		}
	}

	fmt.Println(min, max)
}

func cards(a, k, n int64) (int64, int64) {
	var kol, ost int64
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

func minIter(a, k, n, min *int64) {
	var tmp int64
	if (*n>0) {
		for (*a > 0) {
			if (*n > 0) {
				if (*k > 1) {
					tmp = *k - 1
				} else {
					if (*k>*n) {
						tmp = *n
					} else {
						tmp =*k
					}

					*min += tmp
				}
				*n-=tmp
				*a--

			} else {
				break
			}
		}
	}
}
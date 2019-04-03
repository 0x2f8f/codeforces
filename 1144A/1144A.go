package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	_, err := os.Stat("in.txt")
	if err == nil {
		os.Stdin, err = os.OpenFile("in.txt", os.O_RDONLY, 0666)
	}

	var n int
	var a string
	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		fmt.Scan(&a)
		isDiverse(a);
	}
}

func isDiverse(a string) {
		chars :=  strings.Split(a,"")
		var m []int

		for _, v :=range chars {
			c, _ := utf8.DecodeRuneInString(v)
			in:=int(c-'0')
			m = append(m,in)
		}
		sort.Ints(m)
		i:=0
		j:=0
		for _,v:=range m {
			if i==0 {
				i=v
			} else {
				if (v==(i+1)) {
					i=v
				} else {
					j++
				}
			}
		}
		if j==0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

}
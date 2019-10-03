package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	_, err := os.Stat("in.txt")
	if err == nil {
		os.Stdin, err = os.OpenFile("in.txt", os.O_RDONLY, 0666)
	}

	ca := []string{"a","o","y","e","u","i"}

	var s string
	fmt.Scan(&s)
	chars := strings.Split(strings.ToLower(s),"")
	for _, c := range chars {
		if (!in_array(c, ca)) {
			fmt.Print(".")
			fmt.Print(c)
		}

	}
}

func in_array (a string, list []string) bool {
	for _, b := range list {
		if (b == a) {
			return true
		}
	}
	return false
}

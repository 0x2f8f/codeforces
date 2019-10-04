package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main() {
	_, err := os.Stat("in.txt")
	if err == nil {
		os.Stdin, err = os.OpenFile("in.txt", os.O_RDONLY, 0666)
	}

	var s string
	fmt.Scan(&s)

	fmt.Println(getStr2(s))
}

func in_array (a string, list []string) bool {
	for _, b := range list {
		if (b == a) {
			return true
		}
	}
	return false
}

func getStr (str string) string {
	numbers := []string{"0","1","2","3","4","5","6","7","8","9"}
	s := ""
	currentNstr := ""
	var currentN int
	statusN := true
	chars :=  strings.Split(str,"")
	if (in_array(chars[0],numbers)) {
		for _, c := range chars {
			if (in_array(c, numbers)) { //число
				currentNstr += c
			} else {
				if (statusN == true) {
					// Уже не число. Преобразуем currentNstr в число currentN и обнуляем currentNstr
					currentN, _ = strconv.Atoi(currentNstr)
					statusN = false
				}
				//работа только с буквами
				s += c
			}
		}
		return strings.Repeat(s, currentN)
	} else {
		return str
	}
}

func getStr2 (s string) string {
	if (strings.Index(s, "(")>=0) {
		modeSearch := false
		num_vloj := 0
		searchStr := ""
		result := ""
		chars := strings.Split(s, "")
		for _, c := range chars {
			if (modeSearch == true) {
				if (c == "(") {
					num_vloj++
					searchStr += c
					continue
				} else if (c == ")") {
					num_vloj--

					if (num_vloj == 0) {
						modeSearch = false
						result += getStr2(searchStr)
						searchStr = ""
					} else {
						searchStr += c
					}

				} else {
					searchStr += c
					continue
				}
			} else {
				if (c == "(") {
					modeSearch = true
					num_vloj++
					continue
				} else {
					//просто конкантерируем с результатом
					result += c
				}
			}
		}
		return getStr(result)
	} else {
		return getStr(s)
	}
}
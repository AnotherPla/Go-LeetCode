package main

import (
	"fmt"
	"strings"
)

func IsNumber(s string) bool {
	for key, value := range s {
		if value != ' ' {
			s = s[key:]
			break
		}
	}
	fmt.Printf("s:%v\n", s)
	var tmp []string
	value := strings.Split(s, " ")[0]
	fmt.Printf("value:%v\n", value)
	flag := false
	if strings.Contains(value, "e") {
		tmp = strings.Split(s, "e")
		flag = true
	} else if strings.Contains(value, "E") {
		tmp = strings.Split(s, "E")
		flag = true
	}
	isNum := func(s string) bool {
		if len(s) == 0 {
			return false
		}
		if s[0] == '+' || s[0] == '-' {
			s = s[1:]
		}
		for _, value := range s {
			switch value {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				continue
			default:
				return false
			}
		}
		return true
	}
	isPointNum := func(s string) bool {
		if s[0] == '+' || s[0] == '-' {
			s = s[1:]
		}
		var tmp []string
		fmt.Printf("s:%v\n", s)
		if strings.Contains(s, ".") {
			tmp = strings.Split(s, ".")
		} else {
			tmp = []string{s}
		}
		for _, value := range tmp {
			if !isNum(value) && value != "" {
				fmt.Printf("value: %v\n", value)
				return false
			}
		}
		return true
	}
	if flag {
		x, y := tmp[0], tmp[1]
		println(".......")
		if (isPointNum(x) || isNum(x)) && isNum(y) {
			return true
		}
	}
	//fmt.Printf("isPointNum(value): %v\n", isPointNum(value))
	//fmt.Printf("isNum(value): %v\n", isNum(value))
	return isPointNum(value) || isNum(value)

}

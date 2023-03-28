package main

import (
	"fmt"
)

func ReplaceSpace(s string) string {
	var rep []byte
	for i, val := range s {
		fmt.Printf("val: %v\n", val)
		fmt.Printf("s[i]: %v\n", s[i])
		if val == ' ' {
			rep = append(rep, '%', '2', '0')
		} else {
			rep = append(rep, s[i])
		}
	}
	fmt.Printf("rep: %s\n", rep)
	fmt.Printf("rep[0:5]: %s\n", rep[0:5])
	return string(rep)
}

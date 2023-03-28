package main

import "fmt"

func FirstUniqChar(s string) byte {
	hash := make([]int, 300)
	for i, val := range s {
		if hash[val] > 0 {
			hash[val]++
			/* if s[i] == 'n' {
				fmt.Printf("val: %c, hash[val]: %v\n", s[i], hash[val])
			} */
			fmt.Printf("s[i]: %v\n", s[i])
			fmt.Printf("val: %v\n", val)
		} else {
			hash[val] = 1
		}
	}
	for i, val := range s {
		if hash[val] == 1 {
			return s[i]

		}
	}
	return ' '
}

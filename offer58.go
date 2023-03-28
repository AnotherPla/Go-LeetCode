package main

func ReverseLeftWords(s string, n int) string {
	var char []byte
	for i := range s {
		char = append(char, s[i])
	}
	var front, last []byte
	front = char[0:n]
	last = char[n:]

	// 等价于
	/* for i := range front {
		last = append(last, front[i])
	} */
	// 或者
	/* char = append(char[n:], char[0:n]...) */
	last = append(last, front...)

	return string(last)

}

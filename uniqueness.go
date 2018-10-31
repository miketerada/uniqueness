package main

import (
	"fmt"
	"os"
	"strconv"
)

func ChopUp(s string, n int) map[string]int {
	if 0 == len(s) || n > len(s) || 0 >= n {
		return nil
	}
	m := make(map[string]int)

	for i := 0; n <= len(s); i++ {
		key := s[:n]
		m[key] = m[key] + 1
		s = s[1:]
	}
	return m
}

func main() {
	if 3 != len(os.Args) {
		fmt.Println("Enter a string along with a string length.")
		os.Exit(1)
	}
	str := os.Args[1]
	length, err := strconv.Atoi(os.Args[2])
	if nil != err {
		fmt.Println("Enter a numeric value for string length.")
		os.Exit(1)
	}

	m := ChopUp(str, length)
	fmt.Println(m)
}

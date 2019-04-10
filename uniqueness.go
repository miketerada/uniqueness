// This application is a simple exercise to extract a list of unique strings
// within a given string.
package main

import (
	"fmt"
	"os"
	"strconv"
)

// This application will return unique strings of a given length from
// a given string.
func ChopUp(s string, n int) map[string]int {
	// Check for invalid input.
	if 0 == len(s) || n > len(s) || 0 >= n {
		return nil
	}

	// Create a new map to hold value.
	m := make(map[string]int)

	// Loop through each string 'n' length and increment value in map.
	// Since we have a 'copy' of the input string we can reset 's' within
	// the loop as we remove the first character each iteration.
	for i := 0; n <= len(s); i++ {
		key := s[:n]
		//m[key] = m[key] + 1
		m[key]++
		s = s[1:]
	}
	return m
}

// Main
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

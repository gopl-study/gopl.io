package main

import "fmt"

func deduplication(strings []string) {
	var before string
	for i, s := range strings {
		if before == s {
			strings[i] = ""
		}
		before = s
	}
}

func main() {
	s := []string{"one", "two", "two", "two", "five"}
	deduplication(s)
	fmt.Printf("%#v\n", s) // []string{"one", "two", "", "", "five"}
}

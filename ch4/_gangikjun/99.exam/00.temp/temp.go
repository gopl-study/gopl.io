package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
}

func remove(slice []int, i int) []int {
	fmt.Println(slice[i:])
	fmt.Println(slice[i+1:])
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

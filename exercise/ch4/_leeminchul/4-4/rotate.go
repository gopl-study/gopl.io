package main

import "fmt"

// arr 를 k 번 왼쪽으로 rotate
func rotateJuggling(arr []int, k int) {
	arrLen := len(arr)
	JumpIndex := func(i int) int {
		if arrLen <= i {
			i -= arrLen
		}
		return i
	}

	gcdResult := gcd(k, arrLen)
	for i := 0; i < gcdResult; i++ {
		temp := arr[i]
		var nextIndex int
		for j := i; true; {
			nextIndex = JumpIndex(j + k)
			if nextIndex == i {
				arr[j] = temp
				break
			}
			arr[j] = arr[nextIndex]
			j = nextIndex
		}
	}
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotateJuggling(a[:], 2)
	fmt.Println(a) // [2 3 4 5 6 7 8 9 0 1]
}

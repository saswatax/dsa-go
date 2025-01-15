package main

import "fmt"

func towerOfHanoi(n int, a rune, b rune, c rune) {
	if n == 1 {
		fmt.Printf("Move 1 from %v to %v\n", string(a), string(c))
		return
	}
	towerOfHanoi(n-1, a, c, b)
	fmt.Printf("Move %v from %v to %v\n", n, string(a), string(c))
	towerOfHanoi(n-1, b, a, c)
}

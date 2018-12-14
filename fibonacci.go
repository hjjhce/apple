package main

import (
	"fmt"
)

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func fib2(n int) {
	x := 0
	y := 1
	fmt.Println(x)
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Println(x)
	}
}

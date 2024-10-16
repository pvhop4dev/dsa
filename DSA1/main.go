package main

import (
	"dsa/dsa1/leetcode"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(leetcode.Fibonacci(45))

	fmt.Println(time.Since(now))
}

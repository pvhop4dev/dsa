package main

import (
	"dsa/dsa1/leetcode"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(leetcode.IsPalindrome("A man, a plan, a canal: Panama"))

	fmt.Println(time.Since(now))
}

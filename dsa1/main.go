package main

import (
	"dsa/dsa1/leetcode"
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(leetcode.LongestCommonPrefix([]string{"flower","flow","flight"}))

	fmt.Println(time.Since(now))
}

// Write an algorithm to determine if a number n is "happy".

// A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

// Return True if n is a happy number, and False if not.
// Input: 19
// Output: true
// Explanation:
// 12 + 92 = 82
// 82 + 22 = 68
// 62 + 82 = 100
// 12 + 02 + 02 = 1

package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	resMap := make(map[int]bool)
	res := 0
	for {
		res = calcSqrtDigit(n)

		if res == 1 {
			return true
		}

		if _, ok := resMap[res]; ok {
			return false
		} else {
			resMap[res] = true
		}
		n = res
	}
}

func calcSqrtDigit(num int) int {
	temp := num
	result := 0
	for temp > 0 {
		d := temp % 10
		result += d * d
		temp /= 10
	}
	return result
}

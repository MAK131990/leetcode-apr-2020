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

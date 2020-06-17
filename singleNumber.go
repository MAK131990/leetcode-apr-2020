package main

import "fmt"

func main() {
	res := singleNumber([]int{1, 2, 1})
	fmt.Println(res)
}

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}

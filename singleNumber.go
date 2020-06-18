//Given a non-empty array of integers, every element appears twice except for one. Find that single one.
// Input: [4,1,2,1,2]
// Output: 4

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

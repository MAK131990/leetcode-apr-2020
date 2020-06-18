//Given an array nums, write a function to move all 0's to the end of it while
//maintaining the relative order of the non-zero elements.

// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]

package main

import "fmt"

func main() {
	a := []int{0, 1}
	moveZeroes(a)
	fmt.Println(a)
}

func moveZeroes(nums []int) {
	length := len(nums)
	count := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	fmt.Println(count, nums)
	for i := count; i < length; i++ {
		nums[i] = 0
	}
}

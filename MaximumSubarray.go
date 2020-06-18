// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.

package main

import (
	"fmt"
)

func main() {
	res := maxSubarray([]int{-2, -1})
	fmt.Println(res)
}

func maxSubarray(nums []int) int {
	maxSoFar := 0
	maxEndingHere := 0
	negCount := 0
	minNegNum := 0
	for i := 0; i < len(nums); i++ {
		maxEndingHere += nums[i]
		if nums[i] < 0 {
			if negCount == 0 {
				minNegNum = nums[i]
			}
			negCount++
			if minNegNum < nums[i] {
				minNegNum = nums[i]
			}
		}
		if maxEndingHere < 0 {
			maxEndingHere = 0
		} else {
			if maxSoFar < maxEndingHere {
				maxSoFar = maxEndingHere
			}
		}
	}
	if negCount == len(nums) {
		return minNegNum
	}
	return maxSoFar
}

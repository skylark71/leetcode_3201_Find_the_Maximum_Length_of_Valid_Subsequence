package main

import "fmt"

func main() {
	arr := []int{2, 39, 23}
	fmt.Println(maximumLength(arr))
}

func maximumLength(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	evenCount := 0
	for _, num := range nums {
		if num%2 == 0 {
			evenCount++
		}
	}

	oddCount := 0
	for _, num := range nums {
		if num%2 != 0 {
			oddCount++
		}
	}

	alt1 := 1
	alt2 := 1
	prev1 := nums[0] % 2
	prev2 := 1 - prev1

	for i := 1; i < len(nums); i++ {
		curr := nums[i] % 2
		if curr != prev1 {
			alt1++
			prev1 = curr
		}
		if curr != prev2 {
			alt2++
			prev2 = curr
		}
	}

	// Return the maximum of all cases
	maxLen := evenCount
	if oddCount > maxLen {
		maxLen = oddCount
	}
	if alt1 > maxLen {
		maxLen = alt1
	}
	if alt2 > maxLen {
		maxLen = alt2
	}

	return maxLen
}

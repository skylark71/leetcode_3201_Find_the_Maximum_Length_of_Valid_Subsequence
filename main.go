package main

import "fmt"

func main() {
	arr := []int{1, 2, 1, 1, 2, 1, 2}
	fmt.Println(maximumLength(arr))
}

func maximumLength(nums []int) int {
	evenCount, oddCount := 0, 0
	for _, n := range nums {
		if n%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	// Чередование
	maxAlt1 := alternatingLength(nums, 0) // начиная с чётного
	maxAlt2 := alternatingLength(nums, 1) // начиная с нечётного

	maxAlt := maxAlt1
	if maxAlt2 > maxAlt {
		maxAlt = maxAlt2
	}

	maxSame := evenCount
	if oddCount > maxSame {
		maxSame = oddCount
	}

	if maxAlt > maxSame {
		return maxAlt
	}
	return maxSame
}

func alternatingLength(nums []int, startParity int) int {
	count := 0
	expected := startParity

	for _, n := range nums {
		if n%2 == expected {
			count++
			expected = 1 - expected
		}
	}

	return count
}

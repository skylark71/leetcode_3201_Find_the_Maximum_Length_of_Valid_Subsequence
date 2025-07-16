# Intuition
The task requires finding the longest subsequence where the sum of every two consecutive elements has the same parity (`% 2`).
There are only two possible consistent parity values:

- If all adjacent pairs are (even, even) or (odd, odd), then the sum is always even → `(a + b) % 2 == 0`.
- If adjacent elements alternate between even and odd, then the sum is always odd → `(a + b) % 2 == 1`.

Therefore, the solution focuses on finding:
1. The longest subsequence of elements with the **same parity** (for even-sum case).
2. The longest subsequence of **alternating parity** (for odd-sum case).

---

# Approach
- Traverse the array to count how many even and odd numbers there are.
- The larger count between even and odd gives the length of the longest subsequence with the same parity (`%2 == 0` case).
- Use a helper function `alternatingLength` to compute the longest subsequence with strictly alternating parity (`%2 == 1` case), checking both possible starting parities (even or odd).
- Return the maximum length between these two strategies.

---

# Complexity
- **Time Complexity:** $O(n)$
  We make a few linear passes through the array.

- **Space Complexity:** $O(1)$
  Only constant extra space is used (no auxiliary data structures).

---

# Code
```go
func maximumLength(nums []int) int {
	evenCount, oddCount := 0, 0
	for _, n := range nums {
		if n%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	// Check alternating parity subsequences
	maxAlt1 := alternatingLength(nums, 0) // starting with even
	maxAlt2 := alternatingLength(nums, 1) // starting with odd

	maxAlt := maxAlt1
	if maxAlt2 > maxAlt {
		maxAlt = maxAlt2
	}

	// Check same-parity subsequence
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
			expected = 1 - expected // toggle parity
		}
	}

	return count
}

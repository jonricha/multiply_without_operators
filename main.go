package main

import (
	"fmt"
	"strconv"
)

func abs(in int64) int64 {
	if in < 0 {
		return -in
	}
	return in
}

func recursiveMultiply(a int64, b int64, negative bool) int64 {
	if b == 0 {
		return 0
	}
	result := int64(0)
	if b > 1 {
		result = recursiveMultiply(a, b-1, negative)
	}
	if negative {
		return result - a
	}
	return result + a
}

// RecursiveMultiply multiplies two numbers by recursively adding numbers until the final sum is achieved.
// It works, but has problems with stack overflow if given large numbers.  Also, it does O(n) additions, which could be better.
// If the division operator was allowed we could do a divide and conquer algorithm...
func RecursiveMultiply(a int64, b int64) int64 {
	negative := (a < 0) != (b < 0)
	aAbs := abs(a)
	bAbs := abs(b)

	// Optimization (optional): make sure aAbs > bAbs to have fewer additions (as order does not matter to the result)
	if aAbs < bAbs {
		// avoid allocating extra temp memory by using xor to swap
		aAbs = aAbs ^ bAbs
		bAbs = aAbs ^ bAbs
		aAbs = aAbs ^ bAbs
	}

	return recursiveMultiply(aAbs, bAbs, negative)
}

// LoopMultiply uses a simple loop to multiply. Not the greatest solution as it can be quite slow with large numbers.
func LoopMultiply(a int64, b int64) int64 {
	result := int64(0)
	negative := (a < 0) != (b < 0)
	aAbs := abs(a)
	bAbs := abs(b)

	// Optimization (optional): make sure aAbs > bAbs to have fewer additions (as order does not matter to the result)
	if aAbs < bAbs {
		// avoid allocating extra temp memory by using xor to swap
		aAbs = aAbs ^ bAbs
		bAbs = aAbs ^ bAbs
		aAbs = aAbs ^ bAbs
	}

	for ; bAbs > 0; bAbs-- {
		result += aAbs
	}

	if negative {
		return -result
	}
	return result
}

func memoMultiply(a int64, b int64, memo []int64) int64 {
	if b == 0 {
		return 0
	} else if b == 1 {
		return a
	}
	result := int64(0)
	// recursively call this for each power of 2 (starting with the biggest)
	for check := uint(62); check > 0; check-- {
		powerOf2 := int64(1 << check)
		// check if the check-th bit is set
		if (powerOf2)&b != 0 {
			memoIndex := check - 1 // minus 1 because there's no zero-th element
			// check if we have already calculated/stored this one in the memo
			if memo[memoIndex] == 0 {
				// Get the next biggest value and double it
				halfValue := memoMultiply(a, powerOf2>>1, memo)
				memo[memoIndex] = halfValue + halfValue
			}
			result += memo[memoIndex]
			if result < 0 {
				panic("overflowed")
			}
		}
	}

	return result
}

// MemoMultiply uses memoization to not repeat additions that have already been done. Drawback is extra space is required (62 uint64-s).
// Note, unlike the "*" operator, this function panics when it encounters overflow.  It could be handled other ways, but the panic was simply
// chosen for simplicity.
func MemoMultiply(a int64, b int64) int64 {
	negative := (a < 0) != (b < 0)
	aAbs := abs(a)
	bAbs := abs(b)

	// Optimization (optional): make sure aAbs > bAbs to have fewer additions (as order does not matter to the result)
	if aAbs < bAbs {
		// avoid allocating extra temp memory by using xor to swap
		aAbs = aAbs ^ bAbs
		bAbs = aAbs ^ bAbs
		aAbs = aAbs ^ bAbs
	}

	// The size of the memo is 62 to match the int64 used (minus 2 as we don't need one for the sign bit or for the zero-th bit)
	result := memoMultiply(aAbs, bAbs, make([]int64, 62))

	if negative {
		return -result
	}
	return result
}

// This main function is really only for a sanity test.  See main_test.go for a more complete test set.
func main() {
	a := int64(10)
	b := int64(11)
	result := RecursiveMultiply(a, b)
	fmt.Println("Recursive result: " + strconv.FormatInt(result, 10))
	result = LoopMultiply(a, b)
	fmt.Println("Loop result: " + strconv.FormatInt(result, 10))
	result = MemoMultiply(a, b)
	fmt.Println("Memo result: " + strconv.FormatInt(result, 10))
}

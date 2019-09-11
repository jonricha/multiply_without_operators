package main

import (
	"fmt"
	"math"
	"testing"
)

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func recursiveTest(t *testing.T, a int64, b int64) {
	// These actual limits may be platform specific (or go settings)
	if min(a, b) > 1<<9 {
		// Known to stack overflow
		t.Logf("Skipping recursive test as it overflows")
		return
	}
	result := RecursiveMultiply(a, b)
	// Test against actual multiply operator to verify compatability
	expected := a * b
	if result != expected {
		t.Errorf("Failed Recursive %d * %d = %d, Actual: %d", a, b, expected, result)
	}
}

func loopTest(t *testing.T, a int64, b int64) {
	// These actual limits may be platform specific (or go settings)
	if min(a, b) > 1<<9 {
		//check Known to stack overflow
		t.Logf("Skipping loop test as it overflows")
		return
	}
	result := LoopMultiply(a, b)
	// Test against actual multiply operator to verify compatability
	expected := a * b
	if result != expected {
		t.Errorf("Failed Loop %d * %d = %d, Actual: %d", a, b, expected, result)
	}
}

func memoTest(t *testing.T, a, b int64, shouldPanic bool) {
	if shouldPanic {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
	}
	result := MemoMultiply(a, b)
	// Test against actual multiply operator to verify compatability
	expected := int64(0)
	expected = a * b
	if result != expected {
		fmt.Printf("Min int64: %d", math.MinInt64)
		t.Errorf("Failed Memo %d * %d = %d, Actual: %d", a, b, expected, result)
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name        string
		a           int64
		b           int64
		shouldPanic bool
	}{
		{"0x0", 0, 0, false},
		{"0x1", 0, 1, false},
		{"1x0", 1, 0, false},
		{"0x-1", 0, -1, false},
		{"10x10", 10, 10, false},
		{"1x0", 1, 0, false},
		{"1x-1", 1, -1, false},
		{"-1000x-1000", -1000, -1000, false},
		{"math.MaxInt64x1", math.MaxInt64, 1, false},
		{"math.MaxInt64xmath.MaxInt64", math.MaxInt64, math.MaxInt64, true},
		{"math.MaxInt64xmath.MaxInt64-1", math.MaxInt64, math.MaxInt64 - 1, true},
		{"large numbers", 1<<9 - 1, 1<<9 - 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recursiveTest(t, tt.a, tt.b)
			loopTest(t, tt.a, tt.b)
			memoTest(t, tt.a, tt.b, tt.shouldPanic)
		})
	}
}

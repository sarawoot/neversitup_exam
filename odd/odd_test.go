package odd_test

import (
	"exam/odd"
	"testing"
)

func TestFindOdd(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Case 1", []int{7}, 7},
		{"Case 2", []int{0}, 0},
		{"Case 2", []int{1, 1, 2}, 2},
		{"Case 4", []int{0, 1, 0, 1, 0}, 0},
		{"Case 5", []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := odd.Count(tc.input)
			if result != tc.expected {
				t.Errorf("Count(%v) = %d, want %d", tc.input, result, tc.expected)
			}
		})
	}
}

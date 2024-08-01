package smile_test

import (
	"exam/smile"
	"testing"
)

func TestCount(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected int
	}{
		{"Case 1", []string{}, 0},
		{"Case 2", []string{":)", ";(", ";}", ":-D"}, 2},
		{"Case 3", []string{";D", ":-(", ":-)", ";~)"}, 3},
		{"Case 4", []string{";]", ":[", ";*", ":$", ";-D"}, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := smile.Count(tc.input)
			if result != tc.expected {
				t.Errorf("Count(%v) = %d, want %d", tc.input, result, tc.expected)
			}
		})
	}
}

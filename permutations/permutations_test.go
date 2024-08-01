package permutations_test

import (
	"exam/permutations"
	"reflect"
	"sort"
	"testing"
)

func TestPermutations(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"aabb", []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := permutations.Permutations(tc.input)

			sort.Strings(tc.expected)
			sort.Strings(result)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Permutations(%s) = %v, want %v", tc.input, tc.expected, result)
			}
		})
	}
}

package binary_search

import "testing"

func TestRank(t *testing.T) {
	tests := []struct {
		key      int
		a        []int
		expected int
	}{
		{3, []int{1, 2, 3, 4, 5}, 2},       // key in the middle
		{1, []int{1, 2, 3, 4, 5}, 0},       // key at the beginning
		{5, []int{1, 2, 3, 4, 5}, 4},       // key at the end
		{6, []int{1, 2, 3, 4, 5}, -1},      // key not present
		{3, []int{}, -1},                   // empty array
		{3, []int{3}, 0},                   // single element array where key is present
		{4, []int{3}, -1},                  // single element array where key is not present
		{3, []int{1, 2, 3, 3, 3, 4, 5}, 3}, // key with duplicates
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := Rank(test.key, test.a)
			if result != test.expected {
				t.Errorf("Rank(%d, %v) = %d; want %d", test.key, test.a, result, test.expected)
			}
		})
	}
}

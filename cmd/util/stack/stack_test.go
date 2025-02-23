package stack

import "testing"

func TestPop(t *testing.T) {
	tests := []struct {
		name        string
		init        []string // Items to push in order.
		numPop      int      // Number of pop operations.
		expected    string   // Expected result of the final pop.
		expectedLen int      // Expected number of items remaining in the stack.
	}{
		{
			name:        "SuccessPopFromNonEmptyStackWithOneElement",
			init:        []string{"a"},
			numPop:      1,
			expected:    "a",
			expectedLen: 0,
		},
		{
			name:        "SuccessPopLastElementFromStackWithMultipleItems",
			init:        []string{"a", "b", "c"},
			numPop:      1,
			expected:    "c",
			expectedLen: 2,
		},
		{
			name:        "SuccessPopMultipleTimes",
			init:        []string{"a", "b", "c", "d"},
			numPop:      3, // Pops: "d", then "c", then "b"
			expected:    "b",
			expectedLen: 1,
		},
		{
			name:        "FailPopFromEmptyStack",
			init:        []string{},
			numPop:      1,
			expected:    "",
			expectedLen: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			s := NewStack()
			for _, item := range tc.init {
				s.Push(item)
			}
			var result string
			for i := 0; i < tc.numPop; i++ {
				result = s.Pop()
			}
			if result != tc.expected {
				t.Errorf("expected final pop %q, got %q", tc.expected, result)
			}
			if len(s.items) != tc.expectedLen {
				t.Errorf("expected remaining stack length %d, got %d", tc.expectedLen, len(s.items))
			}
		})
	}
}

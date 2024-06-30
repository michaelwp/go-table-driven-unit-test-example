package main

import "testing"

type testCase struct {
	name     string
	a, b     int
	expected int
}

func TestAdd(t *testing.T) {
	testCases := []testCase{
		{"Add positive numbers", 1, 2, 3},
		{"Add negative numbers", -1, -1, -2},
		{"Add positive and negative number", -1, 2, 1},
		{"Add zeros", 0, 0, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}

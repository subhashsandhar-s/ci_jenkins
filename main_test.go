package main

import "testing"

type TestCase struct {
	Input []int
	Output int
}

func TestAdd(t *testing.T) {
	
	testcases := []TestCase{
		{
			Input: []int{2, 3},
			Output: 5,
		},
		{
			Input: []int{2, 0},
			Output: 2,
		},
		{
			Input: []int{0, 2},
			Output: 2,
		},

	}
	for _, test := range testcases {
		actual := Add(test.Input[0], test.Input[1])
		expected := test.Output
		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	}

}


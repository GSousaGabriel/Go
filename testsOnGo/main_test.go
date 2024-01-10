package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{3, 4}, 7},
		{[]int{55, 55}, 110},
		{[]int{-10, 2}, -8},
	}

	for _, v := range tests {
		result := mySum(v.data...)

		if result != v.answer {
			t.Error("Expected:", 5, "got:", result)
		}
	}
}

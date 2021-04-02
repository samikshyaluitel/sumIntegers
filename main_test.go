package main

import (
	"testing"
)

func TestADDSuccess(t *testing.T) {

	var tcs = []struct {
		num1, num2   []int
		name, result string
	}{
		{name: "equal len",
			num1:   []int{1, 2, 3},
			num2:   []int{1, 2, 3},
			result: "444",
		},
		{name: "num1 len > num2 len",
			num1:   []int{3, 2, 1, 0},
			num2:   []int{1, 0, 3},
			result: "3511",
		},
		{name: "num2 len > num1 len",
			num1:   []int{1, 2, 3},
			num2:   []int{4, 2, 1, 2, 3},
			result: "32247",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := Add(tc.num1, tc.num2)
			if tc.result != actual {
				t.Errorf("expected %s got %s", tc.result, actual)
				t.Fail()
			}
		})
	}
}

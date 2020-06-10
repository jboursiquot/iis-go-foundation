package main

import (
	"reflect"
	"testing"
)

func TestFibBetween(t *testing.T) {
	cases := map[string]struct {
		input []int
		want  []int
	}{
		"0, 1":         {input: []int{0, 1}, want: []int{}},
		"1, 2":         {input: []int{1, 2}, want: []int{2}},
		"2, 5":         {input: []int{2, 5}, want: []int{2, 3, 5}},
		"100, 300":     {input: []int{100, 300}, want: []int{144, 233}},
		"100, 100_000": {input: []int{100, 100_000}, want: []int{144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025}},
	}

	for name, tc := range cases {
		got := fibBetween(tc.input[0], tc.input[1])
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestFib(t *testing.T) {
	cases := map[string]struct {
		input int
		want  int
	}{
		"zero":        {input: 0, want: 0},
		"one":         {input: 1, want: 1},
		"two":         {input: 2, want: 1},
		"three":       {input: 3, want: 2},
		"ten":         {input: 10, want: 55},
		"twenty-five": {input: 25, want: 75025},
	}

	for name, tc := range cases {
		got := fib(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

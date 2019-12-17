package main

import (
	"testing"
)

func TestReverseText(t *testing.T) {
	// t.SkipNow()
	cases := []struct {
		scenario string
		input    string
		want     string
	}{
		{
			scenario: "one",
			input:    "Python has superior string manipulation.",
			want:     ".noitalupinam gnirts roirepus sah nohtyP",
		},
		{"two", "Reflection is never clear.", ".raelc reven si noitcelfeR"},
	}

	for _, tc := range cases {
		t.Run(tc.scenario, func(t *testing.T) {
			if tc.want != ReverseText(tc.input) {
				t.Fail()
			}
		})
	}
}

func TestReverseWords(t *testing.T) {
	cases := []struct {
		scenario string
		input    string
		delim    string
		want     string
	}{
		{
			scenario: "one",
			input:    "Python has superior string manipulation.",
			delim:    " ",
			want:     "manipulation. string superior has Python",
		},
		{"two", "This.is.cool", ".", "cool.is.This"},
	}

	for _, tc := range cases {
		t.Run(tc.scenario, func(t *testing.T) {
			if tc.want != ReverseWords(tc.input, tc.delim) {
				t.Fail()
			}
		})
	}
}

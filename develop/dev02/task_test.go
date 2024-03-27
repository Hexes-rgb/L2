package main

import "testing"

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"a10b", "", true},
		{"a1b", "ab", false},
		{"a1b2c3", "abbccc", false},
	}

	for _, tc := range testCases {
		result, err := UnpackString(tc.input)

		if tc.err && err == nil {
			t.Errorf("UnpackString(%q) expected an error, but got none", tc.input)
		} else if !tc.err && err != nil {
			t.Errorf("UnpackString(%q) unexpected error: %v", tc.input, err)
		}

		if result != tc.expected {
			t.Errorf("UnpackString(%q) = %q, want %q", tc.input, result, tc.expected)
		}
	}
}

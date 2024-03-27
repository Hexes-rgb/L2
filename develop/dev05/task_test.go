package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	tests := []struct {
		name       string
		pattern    string
		after      int
		before     int
		context    int
		count      bool
		ignoreCase bool
		invert     bool
		fixed      bool
		lineNum    bool
		want       string
	}{
		{
			name:       "simple match",
			pattern:    "hello",
			ignoreCase: false,
			invert:     false,
			fixed:      false,
			want:       "hello world\nhello Go\n",
		},
		{
			name:       "case insensitive match",
			pattern:    "HELLO",
			ignoreCase: true,
			invert:     false,
			fixed:      false,
			want:       "hello world\nhello Go\n",
		},
		{
			name:       "invert match",
			pattern:    "Go",
			ignoreCase: false,
			invert:     true,
			fixed:      false,
			want:       "hello world\n",
		},
		{
			name:       "fixed string match",
			pattern:    "Go",
			ignoreCase: false,
			invert:     false,
			fixed:      true,
			want:       "Go\n",
		},
	}

	input := "hello world\nhello Go\ngoodbye Go\nGo is awesome\nGo"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(input))
			var output bytes.Buffer
			writer := bufio.NewWriter(&output)

			grep(scanner, tt.pattern, tt.after, tt.before, tt.count, tt.ignoreCase, tt.invert, tt.fixed, tt.lineNum, writer)
			writer.Flush()

			got := output.String()
			if got != tt.want {
				t.Errorf("grep() = %q, want %q", got, tt.want)
			}
		})
	}
}

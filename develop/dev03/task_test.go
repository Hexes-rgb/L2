package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestSortUtility(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{[]string{"test.txt"}, "Audi 2015 20000\nAudi 2015 20000\nBMW 2018 35000\nFord 2010 10000\nToyota 2020 30000\n"},
		{[]string{"-k", "2", "-n", "test.txt"}, "Ford 2010 10000\nAudi 2015 20000\nAudi 2015 20000\nToyota 2020 30000\nBMW 2018 35000\n"},
		{[]string{"-r", "test.txt"}, "Toyota 2020 30000\nFord 2010 10000\nBMW 2018 35000\nAudi 2015 20000\nAudi 2015 20000\n"},
		{[]string{"-u", "test.txt"}, "Audi 2015 20000\nBMW 2018 35000\nFord 2010 10000\nToyota 2020 30000\n"},
	}

	cmd := exec.Command("go", "build", "-o", "task_test", ".")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Cannot compile sort utility: %s", err)
	}
	defer os.Remove("task_test")

	for _, test := range tests {
		cmd := exec.Command("./task_test", test.args...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Errorf("Failed to execute task_test with args %v: %s", test.args, err)
		}
		if strings.Trim(string(out), "\n") != strings.Trim(test.expected, "\n") {
			t.Errorf("task_test %v = %q, want %q", test.args, out, test.expected)
		}
	}
}

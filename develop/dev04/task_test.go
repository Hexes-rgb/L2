package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		dict     []string
		expected map[string][]string
	}{
		{
			name: "test1",
			dict: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "слово"},
			expected: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:     "test with single words",
			dict:     []string{"пятак", "слово", "тест"},
			expected: map[string][]string{},
		},
		{
			name:     "empty dictionary",
			dict:     []string{},
			expected: map[string][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAnagrams(tt.dict)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findAnagrams() = %v, want %v", result, tt.expected)
			}
		})
	}
}

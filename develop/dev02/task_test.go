package main

import "testing"

func TestUnpackString(t *testing.T) {
	// Создаем структуру для тестовых кейсов
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
		// Вызываем функцию UnpackString с входной строкой из тестового кейса
		result, err := UnpackString(tc.input)

		// Проверяем на наличие ошибки
		if tc.err && err == nil {
			t.Errorf("UnpackString(%q) expected an error, but got none", tc.input)
		} else if !tc.err && err != nil {
			t.Errorf("UnpackString(%q) unexpected error: %v", tc.input, err)
		}

		// Сравниваем полученный результат с ожидаемым
		if result != tc.expected {
			t.Errorf("UnpackString(%q) = %q, want %q", tc.input, result, tc.expected)
		}
	}
}

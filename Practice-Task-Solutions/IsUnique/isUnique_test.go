package main

import "testing"

func TestIsUnique(t *testing.T) {
	tests := []struct {
		input  string
		result bool
	}{
		{"abcdef", true},
		{"aab", false},
		{"", false},
		{"aa", false},
		{"abcd", true},
	}

	for i, test := range tests {
		if result := IsUnique(test.input); result != test.result {
			t.Errorf("Тест %d: ожидаем %t, получили %t", i+1, test.result, result)
		}
	}
}

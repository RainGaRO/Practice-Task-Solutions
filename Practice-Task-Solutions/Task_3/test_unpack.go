package main

import "testing"

func TestUnpack(t *testing.T) {
	data := map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
	}

	for s, e := range data {
		req, err := Unpack(s)
		if err != nil {
			t.Fatalf("Ошибка распаковки %s: ошибка %v", s, err)
		}
		if req != e {
			t.Fatalf("Ошибка распаквоки %s: получили %v должно быть %v", s, req, e)
		}
	}
}

func TestUnpackError(t *testing.T) {
	s := "45"
	req, err := Unpack(s)
	if req != "" {
		t.Fatalf("Ошибка распаковки %s: должная быть пустая строка", s)
	}
	if err == nil {
		t.Fatalf("Неверная распаковка %s", s)
	}
}

func TestUnpackEscape(t *testing.T) {
	data := map[string]string{
		"qwe\\4\\5": "qwe45",
		"qwe\\45":   "qwe44444",
		"qwe\\\\5":  "qwe\\\\\\\\\\",
	}

	for s, e := range data {
		req, err := Unpack(s)
		if err != nil {
			t.Fatalf("Ошибка распаковки %s: ошибка %v", s, err)
		}
		if req != e {
			t.Fatalf("Ошибка распаквоки %s: получили %v должно быть %v", s, req, e)
		}
	}
}

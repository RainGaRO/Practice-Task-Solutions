package set

import "testing"

func TestAdd(t *testing.T) {
	set := Set{}

	// Добавляем элемент в структуру
	result := set.Add("apple")

	// Проверяем, что элемент был добавлен
	expectedSize := 1
	actualSize := set.Size()
	if actualSize != expectedSize {
		t.Errorf("Ожидалось, что размер структуры будет равен %d, но было %d", expectedSize, actualSize)
	} else if !result {
		t.Error("Элемент не был добавлен в структуру")
	}
}

func TestAddInt(t *testing.T) {
	set := Set{}

	// Добавляем элемент в структуру
	result := set.Add(42)

	// Проверяем, что элемент был добавлен
	expectedSize := 1
	actualSize := set.Size()
	if actualSize != expectedSize {
		t.Errorf("Ожидалось, что размер структуры будет равен %d, но было %d", expectedSize, actualSize)
	} else if !result {
		t.Error("Элемент не был добавлен в структуру")
	}
}

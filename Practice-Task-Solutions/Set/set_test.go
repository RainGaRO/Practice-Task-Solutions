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


func TestRemove(t *testing.T) {
	// Тест 1: Удаление существующего элемента
	s := &Set{[]interface{}{1, 2, 3}}
	if !s.Remove(2) {
		t.Errorf("Элемент не был удалён")
	}
	expected := []interface{}{1, 3}
	if reflect.DeepEqual(s.Collection, expected) == false { //len(s.Collection) != len(expected)
		t.Errorf("Неправильный результат")
	}

	// Тест 2: Удаление несуществующего элемента
	s = &Set{[]interface{}{1, 2, 3}}
	if s.Remove(4) {
		t.Error("Элемент должен быть удалён")
	}
	// Тест 3: Проверка правильности удаления элемента из коллекции
	s = &Set{[]interface{}{1, 2, 3}}
	if s.Remove(2) {
		expected := []interface{}{1, 3}
		if len(s.Collection) != len(expected) {
			t.Errorf("Неправильный размер коллекции после удаления")
		} else {
			i := 0
			for _, v := range s.Collection {
				if i == 1 {
					if v != expected[i] {
						t.Errorf("Значение в коллекции не соответствует ожидаемому")
					}
				}
				i++
			}
		}
	}
}

func TestUnion(t *testing.T) {
	// Тест 1: Объединение двух непустых наборов
	s1 := &Set{[]interface{}{1, 2, 3}}
	s2 := &Set{[]interface{}{4, 5, 6}}
	union := s1.Union(s2)
	expected := []interface{}{1, 2, 3, 4, 5, 6}
	if reflect.DeepEqual(union.Collection, expected) == false {
		t.Errorf("Неправильный результат")
	}
	// Тест 2: Объединение пустого набора с непустым набором
	s1 = &Set{}
	s2 = &Set{[]interface{}{4, 5, 6}}
	union = s1.Union(s2)
	expected = []interface{}{4, 5, 6}
	if reflect.DeepEqual(union.Collection, expected) == false {
		t.Errorf("Неправильный результат")
	}

	// Тест 3: Проверка правильности объединения элементов из двух коллекций
	s1 = &Set{[]interface{}{1, 2, 3}}
	s2 = &Set{[]interface{}{3, 4, 5}}

	// Объединение коллекций
	union = s1.Union(s2)

	// Сравнение результатов
	expected = []interface{}{1, 2, 3, 3, 4, 5}
	if !reflect.DeepEqual(union.Collection, expected) {
		t.Errorf("Неправильный результат")
	}
}
func TestDifference(t *testing.T) {
	// Тест 1: Нахождение разности двух непустых наборов
	s1 := &Set{[]interface{}{1, 2, 3}}
	s2 := &Set{[]interface{}{4, 5, 6}}
	difference := s1.Difference(s2)
	expected := []interface{}{1, 2}
	if reflect.DeepEqual(difference.Collection, expected) == false {
		t.Errorf("Неправильный результат")
	}
	// Тест 2: Нахождение разности пустого набора и непустого набора
	s1 = &Set{}
	s2 = &Set{[]interface{}{4, 5, 6}}
	difference = s1.Difference(s2)
	expected = []interface{}{4, 5, 6}
	if reflect.DeepEqual(difference.Collection, expected) == false {
		t.Errorf("Неправильный результат")
	}

	// Тест 3: Проверка правильности нахождения разности элементов из двух коллекций
	s1 = &Set{[]interface{}{1, 2, 3}}
	s2 = &Set{[]interface{}{3, 4, 5}}
	difference = s1.Difference(s2)
	i := 0
	for _, v := range difference.Collection {
		if i == 0 {
			if v != s1.Collection[i] {
				t.Errorf("Значение в коллекции не соответствует ожидаемому")
			}
		} else if i == 1 {
			if !reflect.DeepEqual(v, s2.Collection[i-1]) {
				t.Errorf("Значение в коллекции не соответствует ожидаемому")
			}
		}
		i++
	}
}

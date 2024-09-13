package main

import (
	"testing"
)

func TestPointDistance(t *testing.T) {
	p1 := Point{x: 3, y: 6}
	p2 := Point{x: 5, y: 8}

	expectDist := 5

	distance := p1.Distance(&p2)

	if distance != float64(expectDist) {
		t.Error("Ожидаем:", expectDist, "получили: ", distance)
	}
}

func TestPointInRadius(t *testing.T) {
	center := Point{x: 4, y: 6}
	radius := 3.0

	p := Point{x: 0, y: 0}

	expectResult := true

	result := p.InRadius(&center, radius)

	if result != expectResult {
		t.Error("ождидаем: ", expectResult, "получили: ", result)
	}

	p = Point{x: -1, y: -1}

	expectResult = false

	result = p.InRadius(&center, radius)

	if result != expectResult {
		t.Error("ожидаем: ", expectResult, "получили: ", result)
	}
}

/*
Реализовать структуру Point (координаты X и Y в плоскости)
Реализовать методы для вычисления расстояния между двумя точками.
Реализовать метод для проверки "попадает ли точка в радиус N относительно другой точки"
Написать тесты для этого.
*/
// X и y должны быть с боььшой буквы
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Distance(points *Point) float64 {
	dx := p.x - points.x
	dy := p.y - points.y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p *Point) InRadius(center *Point, radius float64) bool {
	distance := p.Distance(center)
	return distance <= radius
}

func main() {
	p1 := Point{
		x: 2,
		y: 3,
	}
	p2 := Point{
		x: 5,
		y: 7,
	}

	fmt.Println("Расстояние мужду точками- ", p1.Distance(&p2))

	center := Point{
		x: 4,
		y: 6,
	}

	radius := 3.0

	if p1.InRadius(&center, radius) {
		fmt.Println(p1, "попадает в радиус", radius, "относительно точки", center)
	} else {
		fmt.Println(p1, "не попадает в радиус", radius, "относительно точки", center)
	}

}

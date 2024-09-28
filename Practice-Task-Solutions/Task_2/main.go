/*
Реализовать структуру "Полигон" (набор точек, образующих замкнутый многоугольник)
Реализовать структуру "Прямоугольник", "Квадрат", "Круг" на основе точек и полигона
Сделать методы для получения площади фигур
*/

package main

import (
	"fmt"
	"math"
)

// type ISquare interface {
// 	Area()
// }

func (p Polygon) Area() float64 {
	area := 0.0
	for i := 0; i < len(p.Points)-1; i++ {
		x1, y1 := p.Points[i].X, p.Points[i].Y
		x2, y2 := p.Points[i+1].X, p.Points[i+1].Y

		area += (x1*y2 - x2*y1) / 2
	}

	return area
}

func (s Square) Area() float64 {
	// fmt.Println("Площадь квадрата:", s.polygon.Area())
	return s.Polygon.Area()
}

func (r Rectangle) Area() float64 {
	// fmt.Println("Площадь прямоугольника:", r.polygon.Area())
	return r.Polygon.Area()
}

func (c Circle) Area() float64 {
	// fmt.Println("Площадь круга:", math.Pi*c.radius*c.radius)
	return math.Pi * c.Radius * c.Radius
}

func (p Polygon) Distance(point1, point2 Point) float64 {
	x1, y1 := point1.X, point1.Y
	x2, y2 := point2.X, point2.Y

	dx := x2 - x1
	dy := y2 - y1

	return math.Sqrt(dx*dx + dy*dy)
}

func (p Polygon) Contains(point Point) bool {
	var orientation int

	for i := 0; i < len(p.Points); i++ {
		if p.Distance(p.Points[i], point) <= 0 {
			orientation = -1
		} else {
			orientation = 1
		}
	}

	return orientation == 1
}

func main() {
	p1 := Point{X: 0, Y: 0}
	p2 := Point{X: 1, Y: 0}
	p3 := Point{X: 1, Y: 1}
	p4 := Point{X: 0, Y: 1}

	centerX := 0.0
	centerY := 0.0
	radius := 2.0

	polygon := Polygon{Points: []Point{p1, p2, p3, p4}}
	fmt.Println("Площадь полигона:", polygon.Area())

	square := Square{Polygon: polygon}
	fmt.Println("Площадь квадрата:", square.Area())

	rectangle := Rectangle{Polygon: Polygon{Points: []Point{p1, p2, p3, p4}}}
	fmt.Println("Площадь прямоугольника:", rectangle.Area())

	circle := Circle{Center: Point{X: centerX, Y: centerY}, Radius: radius}
	fmt.Println("Площадь круга:", circle.Area())

	/*
		polygon := Polygon{points: []Point{p1, p2, p3, p4}}
		square := Square{polygon: polygon}
		rectangle := Rectangle{polygon: Polygon{points: []Point{p1, p2, p3, p4}}}
		circle := Circle{center: Point{x: centerX, y: centerY}, radius: radius}

		squares := []ISquare{&square, &rectangle, &circle}

		for _, square := range squares {
			square.Area()
		}

		fmt.Println("Площадь полигона:", polygon.Area())
	*/

}

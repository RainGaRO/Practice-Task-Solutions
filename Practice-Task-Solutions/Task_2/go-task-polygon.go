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

type Point struct {
	x, y float64
}

type Polygon struct {
	points []Point
}

type Square struct {
	polygon Polygon
}

type Rectangle struct {
	polygon Polygon
}

type Circle struct {
	center Point
	radius float64
}

func (p Polygon) Area() float64 {
	area := 0.0
	for i := 0; i < len(p.points)-1; i++ {
		x1, y1 := p.points[i].x, p.points[i].y
		x2, y2 := p.points[i+1].x, p.points[i+1].y

		area += (x1*y2 - x2*y1) / 2
	}

	return area
}

func (s Square) Area() float64 {
	// fmt.Println("Площадь квадрата:", s.polygon.Area())
	return s.polygon.Area()
}

func (r Rectangle) Area() float64 {
	// fmt.Println("Площадь прямоугольника:", r.polygon.Area())
	return r.polygon.Area()
}

func (c Circle) Area() float64 {
	// fmt.Println("Площадь круга:", math.Pi*c.radius*c.radius)
	return math.Pi * c.radius * c.radius
}

func main() {
	p1 := Point{x: 0, y: 0}
	p2 := Point{x: 1, y: 0}
	p3 := Point{x: 1, y: 1}
	p4 := Point{x: 0, y: 1}

	centerX := 0.0
	centerY := 0.0
	radius := 2.0

	polygon := Polygon{points: []Point{p1, p2, p3, p4}}
	fmt.Println("Площадь полигона:", polygon.Area())

	square := Square{polygon: polygon}
	fmt.Println("Площадь квадрата:", square.Area())

	rectangle := Rectangle{polygon: Polygon{points: []Point{p1, p2, p3, p4}}}
	fmt.Println("Площадь прямоугольника:", rectangle.Area())

	circle := Circle{center: Point{x: centerX, y: centerY}, radius: radius}
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

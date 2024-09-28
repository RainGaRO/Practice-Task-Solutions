package polygons

type Point struct {
	X, Y float64
}

type Polygon struct {
	Points []Point
}

type Square struct {
	Polygon Polygon
}

type Rectangle struct {
	Polygon Polygon
}

type Circle struct {
	Center Point
	Radius float64
}

package formula

import (
	"log"
	"math"

	"github.com/eyEminYILDIZ/goemetry/types"
)

func example() {
	debug := log.Default()

	point1 := types.Point{X: 10, Y: 10}
	point2 := types.Point{X: 15, Y: 20}

	ang := getAngle(point1, point2)

	debug.Println(ang)
}

func GetRotatingAngle(point1 types.Point, point2 types.Point) float64 {
	angle := getAngle(point1, point2)

	// X:+ Y:+
	if point1.X <= point2.X && point1.Y <= point2.Y {
		return (180 - angle)
	}

	// X:-  Y:+
	if point1.X >= point2.X && point1.Y <= point2.Y {
		return (180 + angle)
	}

	// X:+ Y:-
	if point1.X <= point2.X && point1.Y >= point2.Y {
		return (angle)
	}

	// X:- Y:-
	return 360 - angle
}

// Arc-Tangent implementation
func getAngle(point1 types.Point, point2 types.Point) float64 {
	xDifference := math.Abs(float64(point1.X) - float64(point2.X))
	yDifference := math.Abs(float64(point1.Y) - float64(point2.Y))

	ratio := xDifference / yDifference
	angle := 180 * math.Atan(ratio) / math.Pi

	return angle
}

package formula

import (
	"math"

	"github.com/eyEminYILDIZ/goemetry/types"
)

// func example() {
// 	debug := log.Default()

// 	point1 := types.Point{X: 10, Y: 10}
// 	point2 := types.Point{X: 15, Y: 20}

// 	ang := getAngle(point1, point2)

// 	debug.Println(ang)
// }

func GetRotatingAngle(agent types.Point, target types.Point) float64 {
	angle := getAngle(agent, target)

	// X:+ Y:+
	if agent.X <= target.X && agent.Y <= target.Y {
		return (180 - angle)
	}

	// X:-  Y:+
	if agent.X >= target.X && agent.Y <= target.Y {
		return (180 + angle)
	}

	// X:+ Y:-
	if agent.X <= target.X && agent.Y >= target.Y {
		return (angle)
	}

	// X:- Y:-
	return 360 - angle
}

// Arc-Tangent implementation
func getAngle(agent types.Point, target types.Point) float64 {
	xDifference := math.Abs(float64(agent.X) - float64(target.X))
	yDifference := math.Abs(float64(agent.Y) - float64(target.Y))

	ratio := xDifference / yDifference
	angle := 180 * math.Atan(ratio) / math.Pi

	return angle
}

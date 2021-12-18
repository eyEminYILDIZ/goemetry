package formula

import (
	"math"

	"github.com/eyEminYILDIZ/goemetry/types"
)

func GetNextPoint(agent types.Point, target types.Point, size float64) types.Point {

	nextPoint := types.Point{}

	xDifference := math.Abs(float64(agent.X) - float64(target.X))
	yDifference := math.Abs(float64(agent.Y) - float64(target.Y))

	xRatio := (xDifference / yDifference)
	yRatio := (yDifference / xDifference)

	ratioUnit := size / (xRatio + yRatio)

	xSize := xRatio * ratioUnit
	ySize := yRatio * ratioUnit

	if agent.X <= target.X && agent.Y <= target.Y { // X:+ Y:+ = target is on the right and below
		xSize *= 1
		ySize *= 1
	} else if agent.X >= target.X && agent.Y <= target.Y { // X:-  Y:+ = target is on the left and below
		xSize *= -1
		ySize *= 1
	} else if agent.X <= target.X && agent.Y >= target.Y { // X:+ Y:- = target is on the right and up
		xSize *= 1
		ySize *= -1
	} else { // X:-, Y:- => target is on the left and up
		xSize *= -1
		ySize *= -1
	}

	nextPoint.X = agent.X + int32(xSize)
	nextPoint.Y = agent.Y + int32(ySize)

	return nextPoint
}

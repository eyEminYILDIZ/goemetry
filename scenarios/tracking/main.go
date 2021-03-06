package main

import (
	"fmt"
	"math"

	formula "github.com/eyEminYILDIZ/goemetry/formula"
	types "github.com/eyEminYILDIZ/goemetry/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	agentPoint := types.Point{}
	targetPoint := types.Point{}

	setInitialPoints(&agentPoint, &targetPoint)

	draw(agentPoint, targetPoint)
}

func draw(agentPoint types.Point, targetPoint types.Point) {
	screenWidth := int32(800)
	screenHeight := int32(800)

	rl.InitWindow(screenWidth, screenHeight, "Rotation Calculations")

	// NOTE: Textures MUST be loaded after Window initialization (OpenGL context is required)
	agentTexture := rl.LoadTexture("agent.png")   // Texture loading
	targetTexture := rl.LoadTexture("target.png") // Texture loading

	// NOTE: Source rectangle (part of the texture to use for drawing)
	agentSourceRec := rl.NewRectangle(0, 0, float32(agentTexture.Width), float32(agentTexture.Height))

	// NOTE: Destination rectangle (screen rectangle where drawing part of texture)
	// agentDestRec := rl.NewRectangle(100, 100, 200, 200)

	// NOTE: Origin of the texture (rotation/scale point), it's relative to destination rectangle size
	origin := rl.NewVector2(float32(agentTexture.Width), float32(agentTexture.Height))

	rl.SetTargetFPS(60)

	sensivity := int32(10)

	for !rl.WindowShouldClose() {

		if checkCollision(agentPoint, targetPoint) {
			setInitialPoints(&agentPoint, &targetPoint)
		}

		if rl.IsKeyDown(rl.KeyRight) {
			targetPoint.X += sensivity
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			targetPoint.X -= sensivity
		}
		if rl.IsKeyDown(rl.KeyUp) {
			targetPoint.Y -= sensivity
		}
		if rl.IsKeyDown(rl.KeyDown) {
			targetPoint.Y += sensivity
		}

		keepPointInTheArea(&targetPoint, screenWidth, screenHeight)

		angle := formula.GetRotatingAngle(agentPoint, targetPoint)
		rotation := float32(angle)

		nextPoint := formula.GetNextPoint(agentPoint, targetPoint, 4)
		agentPoint.X = nextPoint.X
		agentPoint.Y = nextPoint.Y

		keepPointInTheArea(&agentPoint, screenWidth, screenHeight)

		agentDestRec := rl.NewRectangle(float32(agentPoint.X), float32(agentPoint.Y), float32(agentTexture.Width*2), float32(agentTexture.Height*2))

		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		// NOTE: Using DrawTexturePro() we can easily rotate and scale the part of the texture we draw
		// sourceRec defines the part of the texture we use for drawing
		// destRec defines the rectangle where our texture part will fit (scaling it to fit)
		// origin defines the point of the texture used as reference for rotation and scaling
		// rotation defines the texture rotation (using origin as rotation point)
		rl.DrawTexturePro(agentTexture, agentSourceRec, agentDestRec, origin, rotation, rl.White)
		rl.DrawTexture(targetTexture, targetPoint.X, targetPoint.Y, rl.White)

		// rl.DrawLine(int32(destRec.X), 0, int32(destRec.X), screenHeight, rl.Gray)
		// rl.DrawLine(0, int32(destRec.Y), screenWidth, int32(destRec.Y), rl.Gray)

		str := fmt.Sprintf("Angle: %.1f", angle)
		rl.DrawText(str, agentPoint.X+50, agentPoint.Y, 20, rl.Black)

		rl.EndDrawing()
	}

	rl.UnloadTexture(agentTexture)

	rl.CloseWindow()
}

func checkCollision(agentPoint types.Point, targetPoint types.Point) bool {
	xDifference := math.Abs(float64(agentPoint.X) - float64(targetPoint.X))
	yDifference := math.Abs(float64(agentPoint.Y) - float64(targetPoint.Y))

	return xDifference < 10 && yDifference < 10
}

func setInitialPoints(agentPoint *types.Point, targetPoint *types.Point) {
	agentPoint.X = 100
	agentPoint.Y = 100
	targetPoint.X = 600
	targetPoint.Y = 600
}

func keepPointInTheArea(targetPoint *types.Point, screenWidth int32, screenHeight int32) {
	if targetPoint.X < 0 {
		targetPoint.X = 0
	}
	if targetPoint.X > (screenWidth - 100) {
		targetPoint.X = screenWidth - 100
	}
	if targetPoint.Y < 0 {
		targetPoint.Y = 0
	}
	if targetPoint.Y > (screenHeight - 100) {
		targetPoint.Y = screenHeight - 100
	}
}

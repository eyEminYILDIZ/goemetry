package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(600)
	screenHeight := int32(600)

	rl.InitWindow(screenWidth, screenHeight, "A Surface")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawLine(screenWidth/2, 0, screenWidth/2, 600, rl.Black)
		rl.DrawLine(0, screenHeight/2, 600, screenHeight/2, rl.Black)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

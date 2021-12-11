package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(600)
	screenHeight := int32(600)

	rl.InitWindow(screenWidth, screenHeight, "raylib [textures] examples - texture source and destination rectangles")

	// NOTE: Textures MUST be loaded after Window initialization (OpenGL context is required)
	agent := rl.LoadTexture("agent.png")   // Texture loading
	target := rl.LoadTexture("target.png") // Texture loading

	// frameWidth := float32(100)
	// frameHeight := float32(100)

	// NOTE: Source rectangle (part of the texture to use for drawing)
	agentSourceRec := rl.NewRectangle(0, 0, 50, 50)

	// NOTE: Destination rectangle (screen rectangle where drawing part of texture)
	// agentDestRec := rl.NewRectangle(100, 100, 200, 200)
	agentDestRec := rl.NewRectangle(100, 100, 100, 100)
	// NOTE: Origin of the texture (rotation/scale point), it's relative to destination rectangle size
	origin := rl.NewVector2(float32(50), float32(50))

	rotation := float32(0)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Update
		rotation++

		// Draw
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		// NOTE: Using DrawTexturePro() we can easily rotate and scale the part of the texture we draw
		// sourceRec defines the part of the texture we use for drawing
		// destRec defines the rectangle where our texture part will fit (scaling it to fit)
		// origin defines the point of the texture used as reference for rotation and scaling
		// rotation defines the texture rotation (using origin as rotation point)
		rl.DrawTexturePro(agent, agentSourceRec, agentDestRec, origin, rotation, rl.White)
		rl.DrawTexture(target, 200, 500, rl.White)

		// rl.DrawLine(int32(destRec.X), 0, int32(destRec.X), screenHeight, rl.Gray)
		// rl.DrawLine(0, int32(destRec.Y), screenWidth, int32(destRec.Y), rl.Gray)

		// rl.DrawText("(c) Scarfy sprite by Eiden Marsal", screenWidth-200, screenHeight-20, 10, rl.Gray)

		rl.EndDrawing()
	}

	rl.UnloadTexture(agent)

	rl.CloseWindow()
}

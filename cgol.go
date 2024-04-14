package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Raylib-Go Template")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawText("Creeper, oh man", 150, 200, 20, rl.White)

		rl.EndDrawing()
	}
}

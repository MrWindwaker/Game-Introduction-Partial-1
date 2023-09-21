package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(320, 180, "Partial 1")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.GetColor(0x333333FF))

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

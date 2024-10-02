package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	xPosition int32   = 400
	yPosition int32   = 250
	radius    float32 = 20
)

func main() {
	rl.InitWindow(800, 500, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(80)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.DrawCircle(xPosition, yPosition, radius, color.RGBA{0x34, 0xeb, 0x64, 0xff})
		if rl.IsKeyPressed(rl.KeyUp) {
			yPosition -= 10
		} else if rl.IsKeyPressed(rl.KeyDown) {
			yPosition += 10
		} else if rl.IsKeyPressed(rl.KeyRight) {
			xPosition += 10
		} else if rl.IsKeyPressed(rl.KeyLeft) {
			xPosition -= 10
		}
		rl.EndDrawing()

	}

}

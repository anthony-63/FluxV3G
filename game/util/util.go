package util

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var MainFont rl.Font

func CenterText(msg string, x int32, y int32, fontsize int32) rl.Vector2 {
	measured := rl.MeasureTextEx(MainFont, msg, float32(fontsize)/2, 0)
	return rl.Vector2{
		X: float32(x - int32(measured.X)),
		Y: float32(y - int32(measured.Y)),
	}
}

func DrawTextFromCenter(msg string, x int32, y int32, fontsize int32, color rl.Color) {
	center := CenterText(msg, x, y, fontsize)

	rl.DrawTextEx(MainFont, msg, center, float32(fontsize), 0, color)
}

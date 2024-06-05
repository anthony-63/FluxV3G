package util

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var MainFont rl.Font

func CenterText(msg string, x float32, y float32, fontsize float32) rl.Vector2 {
	measured := rl.MeasureTextEx(MainFont, msg, fontsize/2, 0)
	return rl.Vector2{
		X: x - measured.X,
		Y: y - measured.Y,
	}
}

func DrawTextFromCenter(msg string, x float32, y float32, fontsize float32, color rl.Color, font rl.Font) {
	center := CenterText(msg, x, y, fontsize)

	rl.DrawTextEx(font, msg, center, fontsize, 0, color)
}

func DrawTextMFont(msg string, x float32, y float32, fontsize float32, color rl.Color) {
	rl.DrawTextEx(MainFont, msg, rl.Vector2{X: x, Y: y}, fontsize, 0, color)
}

func All[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}

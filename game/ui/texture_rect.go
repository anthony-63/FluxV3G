package ui

import (
	"errors"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureRect struct {
	X      float32
	Y      float32
	Width  float32
	Height float32

	Texture  rl.Texture2D
	Centered bool

	loaded bool
}

func (rect *TextureRect) SetImageFromFile(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		rect.loaded = false
	}

	img := rl.LoadImage(path)
	rl.ImageResize(img, int32(rect.Width), int32(rect.Height))
	rect.Texture = rl.LoadTextureFromImage(img)
	rect.loaded = true
}

func (rect *TextureRect) Draw() {
	if !rect.loaded {
		return
	}

	position := rl.Vector2{
		X: rect.X,
		Y: rect.Y,
	}

	if rect.Centered {
		position.X -= rect.Width / 2
		position.Y -= rect.Height / 2
	}

	rl.DrawTextureEx(rect.Texture, position, 0, 1, rl.White)
}

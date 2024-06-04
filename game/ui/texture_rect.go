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

func TextureRectFromPng(path string) TextureRect {
	rect := TextureRect{}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		rect.loaded = false
	}

	return rect
}

func (rect *TextureRect) Draw() {
	if !rect.loaded {
		return
	}
}

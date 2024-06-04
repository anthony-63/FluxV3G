package ui

import rl "github.com/gen2brain/raylib-go/raylib"

type IProgressBar interface {
	Draw()
	Update(float64)
}

type ProgressBar struct {
	X          float32
	Y          float32
	Value      float32
	Max        float32
	BackColor  rl.Color
	FrontColor rl.Color
}

func (bar *ProgressBar) Draw() {

}

func (bar *ProgressBar) Update(dt float64) {

}

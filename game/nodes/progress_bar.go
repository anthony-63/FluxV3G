package nodes

import rl "github.com/gen2brain/raylib-go/raylib"

type ProgressBar struct {
	X      float32
	Y      float32
	Width  float32
	Height float32

	Value      float32
	Max        float32
	BackColor  rl.Color
	FrontColor rl.Color

	Centered bool
}

func (bar *ProgressBar) Draw() {
	x := bar.X
	y := bar.Y

	if bar.Centered {
		x -= bar.Width / 2
		y -= bar.Height / 2
	}

	rl.DrawRectanglePro(rl.Rectangle{
		X:      x,
		Y:      y,
		Width:  bar.Width,
		Height: bar.Height,
	}, rl.Vector2{
		X: 0,
		Y: 0,
	}, 0, bar.BackColor)

	rl.DrawRectanglePro(rl.Rectangle{
		X:      x,
		Y:      y,
		Width:  bar.Width * (bar.Value / bar.Max),
		Height: bar.Height,
	}, rl.Vector2{
		X: 0,
		Y: 0,
	}, 0, bar.FrontColor)
}

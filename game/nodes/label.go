package nodes

import (
	"flux/game/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Label struct {
	X         float32
	Y         float32
	Text      string
	FontSize  float32
	FontColor rl.Color
	Font      rl.Font
	Centered  bool
}

func (label *Label) Draw() {
	if label.Centered {
		util.DrawTextFromCenter(label.Text, label.X, label.Y, label.FontSize, label.FontColor, label.Font)
	} else {
		rl.DrawTextEx(label.Font, label.Text, rl.Vector2{X: label.X, Y: label.Y}, label.FontSize, 0, label.FontColor)
	}
}

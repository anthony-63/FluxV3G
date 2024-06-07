package nodes

import (
	"flux/game/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	X      float32
	Y      float32
	Width  float32
	Height float32

	Color          rl.Color
	HighlightColor rl.Color
	FocusColor     rl.Color
	PressedColor   rl.Color

	BorderColor rl.Color
	BorderWidth float32

	Text           string
	FontSize       float32
	FontColor      rl.Color
	Font           rl.Font
	TextCentered   bool
	AnchorCentered bool
	Roundness      float32

	rect        rl.Rectangle
	border_rect rl.Rectangle

	current_color rl.Color

	focused bool
	pressed bool
}

func (button *Button) ProcessInputs() {
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), button.rect) {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			button.focused = true
			button.current_color = button.FocusColor
		} else if !button.focused {
			button.current_color = button.HighlightColor
		}
	} else if !button.focused {
		button.current_color = button.Color
	} else if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		button.focused = false
		button.current_color = button.Color
	}
}

func (button *Button) Draw() {
	if button.AnchorCentered {
		button.rect = rl.Rectangle{
			X:      button.X - (button.Width / 2),
			Y:      button.Y - (button.Height / 2),
			Width:  button.Width,
			Height: button.Height,
		}
	} else {
		button.rect = rl.Rectangle{
			X:      button.X,
			Y:      button.Y,
			Width:  button.Width,
			Height: button.Height,
		}
	}

	if button.BorderWidth > 0 {
		button.border_rect = button.rect
		button.border_rect.X -= button.BorderWidth
		button.border_rect.Y -= button.BorderWidth
		button.border_rect.Width += button.BorderWidth * 2
		button.border_rect.Height += button.BorderWidth * 2
		rl.DrawRectangleRounded(button.border_rect, button.Roundness, 100, button.BorderColor)
	}

	rl.DrawRectangleRounded(button.rect, button.Roundness, 100, button.current_color)

	if button.TextCentered {
		util.DrawTextFromCenter(button.Text, button.X, button.Y, button.FontSize, button.FontColor, button.Font)
	} else {
		rl.DrawTextEx(button.Font, button.Text, rl.Vector2{X: button.X, Y: button.Y}, button.FontSize, 0, button.FontColor)
	}
}

package scenes

import (
	"flux/game/nodes"
	"flux/game/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuScene struct {
	start_button nodes.Button
}

func CreateMenuScene() *MenuScene {
	scene := MenuScene{
		start_button: nodes.Button{
			X:      300,
			Y:      300,
			Width:  200,
			Height: 40,

			Color:          rl.DarkGray,
			HighlightColor: rl.NewColor(0x5f, 0x5f, 0x5f, 0xff),
			FocusColor:     rl.NewColor(0x3f, 0x3f, 0x3f, 0xff),
			BorderColor:    rl.Gray,

			BorderWidth:    2,
			Text:           "Test button",
			FontSize:       32,
			FontColor:      rl.White,
			Font:           util.MainFont,
			TextCentered:   true,
			AnchorCentered: true,
			Roundness:      1,
		},
	}

	return &scene
}
func (scene *MenuScene) Update(dt float64) {
	scene.start_button.ProcessInputs()
}

func (scene *MenuScene) Draw() {
	scene.start_button.Draw()
	rl.ClearBackground(rl.NewColor(0x10, 0x10, 0x10, 0xff))
}

func (scene *MenuScene) GetType() int {
	return SCENE_TYPE_MENU
}

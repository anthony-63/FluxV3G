package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuScene struct {
}

func CreateMenuScene() *MenuScene {
	scene := MenuScene{}

	return &scene
}
func (scene *MenuScene) Update(dt float64) {
}

func (scene *MenuScene) Draw() {
	rl.ClearBackground(rl.NewColor(0x10, 0x10, 0x10, 0xff))
}

func (scene *MenuScene) GetType() int {
	return SCENE_TYPE_MENU
}

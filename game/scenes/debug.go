package scenes

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type DebugScene struct {
	scene_type int
}

func CreateDebugScene() *DebugScene {
	debug := DebugScene{
		scene_type: SCENE_TYPE_DEBUG,
	}

	return &debug
}

func (debug DebugScene) Update(dt float64) {

}

func (debug *DebugScene) Draw() {
	rl.DrawText(strconv.Itoa(int(rl.GetFPS())), 0, 0, 16, rl.Yellow)

}

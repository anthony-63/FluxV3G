package scenes

import (
	"flux/game/util"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type DebugScene struct {
	scene_type int

	fps_timer float64
	fps_count int32
}

func CreateDebugScene() *DebugScene {
	debug := DebugScene{
		scene_type: SCENE_TYPE_DEBUG,
	}

	return &debug
}

func (debug *DebugScene) Update(dt float64) {
	debug.fps_timer += dt

	if debug.fps_timer >= 0.5 {
		debug.fps_count = rl.GetFPS()
		debug.fps_timer = 0
	}
}

func (debug *DebugScene) Draw() {

	util.DrawTextMFont(fmt.Sprint(debug.fps_count), 0, 0, 32, rl.Yellow)

	for _, scene := range SceneList {
		if scene.GetType() == SCENE_TYPE_GAME {
			game := scene.(*GameScene)
			util.DrawTextMFont(fmt.Sprintf("T: %.02f", game.sync_manger.RealTime), 8, 32*1, 32, rl.Green)
		}
	}
}

func (scene DebugScene) GetType() int {
	return scene.scene_type
}

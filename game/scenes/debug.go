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

	game *GameScene
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
			_ = game
			util.DrawTextMFont(fmt.Sprint("Current: ", util.SelectedMapSet.Title, "[", util.SelectedMap.Name, "]"), 8, 26*1, 26, rl.Green)
			util.DrawTextMFont(fmt.Sprintf("T: %.02f", game.sync_manger.RealTime), 8, 26*2, 26, rl.Green)
		}
	}
}

func (scene DebugScene) GetType() int {
	return scene.scene_type
}

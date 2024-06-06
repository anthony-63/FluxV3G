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
		game:       nil,
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
	// rl.DrawFPS(0, 0)

	if debug.game == nil {
		for _, scene := range SceneList {
			if scene.GetType() == SCENE_TYPE_GAME {
				debug.game = scene.(*GameScene)
			}
		}
		return
	}

	util.DrawTextMFont(fmt.Sprint("Current: ", util.SelectedMapSet.Title /*, "[", util.SelectedMap.Name, "]"*/), 8, 26*1, 26, rl.Green)
	util.DrawTextMFont(fmt.Sprintf("Rendering: %d", len(debug.game.note_renderer.ToRender)), 8, 26*2, 26, rl.Green)
	util.DrawTextMFont(fmt.Sprintf("Note: %d/%d", debug.game.note_manager.StartProcess, len(debug.game.note_manager.OrderedNotes)), 8, 26*3, 26, rl.Green)
	util.DrawTextMFont(fmt.Sprintf("Time: %.02f", debug.game.sync_manger.RealTime), 8, 26*4, 26, rl.Green)
}

func (scene DebugScene) GetType() int {
	return scene.scene_type
}

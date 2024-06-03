package game

import (
	"flux/game/scenes"
	"flux/game/util"

	"github.com/rs/zerolog/log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 1280
const HEIGHT = 720
const TITLE = "FluxV3G | 0.0.1"

type FluxWindow struct {
	scene_list []scenes.IScene
}

func CreateWindow() FluxWindow {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(WIDTH, HEIGHT, TITLE)

	// font := rl.LoadFontEx("data/fonts/noto_sans.ttf", 32, []rune{}, 2840)

	log.Info().Msg("Loading fonts...")

	util.MainFont = rl.LoadFontEx("data/fonts/noto_sans.ttf", 128, nil)

	log.Info().Msg("Done loading fonts")

	return FluxWindow{
		scene_list: []scenes.IScene{
			scenes.CreateStartupScene(),
		},
	}
}

func (window *FluxWindow) RunWindow() {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		for _, scene := range window.scene_list {
			scene.Update(float64(rl.GetFrameTime()))
			scene.Draw()
		}

		rl.EndDrawing()
	}
}

func (window *FluxWindow) WindowCleanup() {
	rl.CloseWindow()
}

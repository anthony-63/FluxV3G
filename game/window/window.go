package window

import (
	"flux/game/scenes"
	"flux/game/util"

	"github.com/rs/zerolog/log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var GameWindow FluxWindow

const WIDTH = 1280
const HEIGHT = 720
const TITLE = "FluxV3-OPT"

type FluxWindow struct {
	scene_list []scenes.IScene
}

func CreateWindow() FluxWindow {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(WIDTH, HEIGHT, TITLE)

	log.Info().Msg("Loading fonts...")

	util.MainFont = rl.LoadFontEx("data/.game/fonts/noto_sans.ttf", 32, nil)
	rl.GenTextureMipmaps(&util.MainFont.Texture)

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

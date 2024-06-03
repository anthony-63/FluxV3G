package scenes

import (
	"strconv"
	"strings"

	"flux/game/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StartupScene struct {
	loadingText string
	dot_timer   float64
}

func CreateStartupScene() *StartupScene {
	return &StartupScene{
		loadingText: "Loading Flux",
	}
}

func (scene *StartupScene) Update(dt float64) {
	scene.dot_timer += dt

	if scene.dot_timer >= 0.25 {
		scene.loadingText += "."
		if strings.Count(scene.loadingText, ".") > 3 {
			scene.loadingText = strings.ReplaceAll(scene.loadingText, ".", "")
		}

		scene.dot_timer = 0
	}
}

func (scene StartupScene) Draw() {
	rl.ClearBackground(rl.NewColor(0x10, 0x10, 0x10, 0xff))

	rl.DrawText(strconv.Itoa(int(rl.GetFPS())), 0, 0, 16, rl.Yellow)

	util.DrawTextFromCenter(scene.loadingText, int32(rl.GetScreenWidth())/2, int32(rl.GetScreenHeight())/2, 50, rl.LightGray)
}

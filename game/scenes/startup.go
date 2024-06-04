package scenes

import (
	"strconv"
	"strings"

	"flux/game/loaders"
	"flux/game/ui"
	"flux/game/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ProgressStruct struct {
	at    int
	total int
	text  string
	done  bool
}

type StartupScene struct {
	dot_timer float64

	progress      ui.ProgressBar
	loading_label ui.Label
}

func CreateStartupScene() *StartupScene {
	go loaders.LoadMaps()

	return &StartupScene{
		progress: ui.ProgressBar{
			X: float32(rl.GetScreenWidth()) / 2,
			Y: float32(rl.GetScreenHeight())/2 - 20,

			BackColor:  rl.DarkBlue,
			FrontColor: rl.SkyBlue,
		},
		loading_label: ui.Label{
			Text:      "Loading Flux",
			X:         float32(rl.GetScreenWidth()) / 2,
			Y:         float32(rl.GetScreenHeight()) / 2,
			FontSize:  30,
			FontColor: rl.LightGray,
			Font:      util.MainFont,
			Centered:  true,
		},
	}
}

func (scene *StartupScene) updateProgress(c chan ProgressStruct) {
	progress := <-c
	for !progress.done {
		scene.progress.Value = float32(progress.at)
		scene.progress.Max = float32(progress.total)
		scene.loading_label.Text = progress.text

		progress = <-c
	}
}

func (scene *StartupScene) Update(dt float64) {
	scene.dot_timer += dt

	if scene.dot_timer >= 0.25 {
		scene.loading_label.Text += "."
		if strings.Count(scene.loading_label.Text, ".") > 3 {
			scene.loading_label.Text = strings.ReplaceAll(scene.loading_label.Text, ".", "")
		}

		scene.dot_timer = 0
	}

	scene.progress.Update(dt)
	scene.loading_label.Update(dt)
}

func (scene StartupScene) Draw() {
	rl.ClearBackground(rl.NewColor(0x10, 0x10, 0x10, 0xff))
	rl.DrawText(strconv.Itoa(int(rl.GetFPS())), 0, 0, 16, rl.Yellow)

	scene.progress.Draw()
	scene.loading_label.Draw()
}

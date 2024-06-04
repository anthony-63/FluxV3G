package scenes

import (
	"strings"

	"flux/game/loaders"
	"flux/game/nodes"
	"flux/game/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type StartupScene struct {
	scene_type int

	dot_timer float64

	progress        nodes.ProgressBar
	loading_label   nodes.Label
	substatus_label nodes.Label
	flux_logo       nodes.TextureRect

	finished bool
}

func CreateStartupScene() *StartupScene {
	progress_chan := make(chan util.ProgressStruct)

	go loaders.LoadMaps(progress_chan)

	scene := StartupScene{
		scene_type: SCENE_TYPE_STARTUP,

		progress: nodes.ProgressBar{
			X: float32(rl.GetScreenWidth()) / 2,
			Y: float32(rl.GetScreenHeight())/2 + 45,

			Width:  300,
			Height: 20,

			BackColor:  rl.NewColor(0x8, 0x8, 0x8, 0xff),
			FrontColor: rl.NewColor(0x4c, 0x4c, 0x4c, 0xff),

			Centered: true,
		},
		loading_label: nodes.Label{
			Text:      "Loading Flux",
			X:         float32(rl.GetScreenWidth()) / 2,
			Y:         float32(rl.GetScreenHeight()) / 2,
			FontSize:  30,
			FontColor: rl.LightGray,
			Font:      util.MainFont,
			Centered:  true,
		},
		substatus_label: nodes.Label{
			Text:      "...",
			X:         float32(rl.GetScreenWidth()) / 2,
			Y:         float32(rl.GetScreenHeight())/2 + 20,
			FontSize:  15,
			FontColor: rl.Gray,
			Font:      util.MainFont,
			Centered:  true,
		},
		flux_logo: nodes.TextureRect{
			X:        float32(rl.GetScreenWidth()) / 2,
			Y:        float32(rl.GetScreenHeight())/2 - 100,
			Width:    150,
			Height:   150,
			Centered: true,
		},

		finished: false,
	}

	scene.flux_logo.SetImageFromFile("data/.game/images/flux.png")

	go scene.updateProgress(progress_chan)

	return &scene
}

func (scene *StartupScene) updateProgress(c chan util.ProgressStruct) {
	progress := <-c
	for !progress.Done {
		scene.progress.Value = float32(progress.At)
		scene.progress.Max = float32(progress.Total)
		scene.substatus_label.Text = progress.Text

		progress = <-c
	}

	scene.finished = true
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

	if scene.finished {
		SceneList = []IScene{CreateGameScene(), CreateDebugScene()}
	}
}

func (scene StartupScene) Draw() {
	rl.ClearBackground(rl.NewColor(0x10, 0x10, 0x10, 0xff))

	scene.progress.Draw()
	scene.loading_label.Draw()
	scene.substatus_label.Draw()
	scene.flux_logo.Draw()
}

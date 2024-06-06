package scenes

import (
	"flux/game/managers"
	"flux/game/nodes"
	"flux/game/util"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rs/zerolog/log"
)

type GameScene struct {
	scene_type int

	camera nodes.Camera
	grid   nodes.Sprite3D

	sync_manger  *managers.SyncManager
	note_manager *managers.NoteManager
}

func CreateGameScene() *GameScene {
	game := GameScene{
		scene_type: SCENE_TYPE_GAME,

		camera: nodes.NewCamera(rl.Vector3{
			X: 0,
			Y: 0,
			Z: 7.5 * util.VFCONV32,
		}),
		grid: nodes.Sprite3D{
			Position: rl.Vector3Zero(),
			Rotation: rl.Vector3{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Size: rl.Vector2One(),
		},
		sync_manger:  managers.CreateSyncManager(),
		note_manager: managers.CreateNoteManager(),
	}

	game.grid.GenPlane(1, 1, "data/.game/game/grid.png")

	log.Info().Str("current_map", util.SelectedMapSet.Title).Msg("Game")
	log.Info().Str("current_difficulty", util.SelectedMap.Name).Msg("Game")

	game.sync_manger.Start(0)

	return &game
}

func (game GameScene) Update(dt float64) {
	game.sync_manger.Update(dt)
}

func (game *GameScene) Draw() {
	rl.ClearBackground(rl.Black)

	rl.BeginMode3D(game.camera.RlCamera)

	game.grid.Draw()

	rl.EndMode3D()
}

func (scene GameScene) GetType() int {
	return scene.scene_type
}

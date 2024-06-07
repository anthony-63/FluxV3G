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
	grid   *nodes.Sprite3D
	cursor *nodes.Cursor

	sync_manger   *managers.SyncManager
	note_renderer *managers.NoteRenderer
	note_manager  *managers.NoteManager
}

func CreateGameScene() *GameScene {
	game := GameScene{
		scene_type: SCENE_TYPE_GAME,

		camera: nodes.NewCamera(rl.Vector3{
			X: 0,
			Y: 0,
			Z: 7.5 * util.VFCONV32,
		}),
		grid: &nodes.Sprite3D{
			Position: rl.Vector3Zero(),
			Rotation: rl.Vector3{
				X: 90,
				Y: 0,
				Z: 0,
			},
			Size: rl.Vector2One(),
		},
		cursor: nodes.CreateCursor(),
	}

	game.sync_manger = managers.CreateSyncManager()
	game.note_renderer = managers.CreateNoteRenderer(game.sync_manger)
	game.note_manager = managers.CreateNoteManager(game.sync_manger, game.note_renderer)

	audio_player := nodes.AudioPlayerFromFile(util.SelectedMapSet.Path + "/" + util.SelectedMapSet.MusicPath)

	game.sync_manger.AudioPlayer = audio_player

	game.grid.GenPlane(1, 1, "data/.game/game/grid.png")

	log.Info().Str("current_map", util.SelectedMapSet.Title).Msg("Game")
	log.Info().Str("current_difficulty", util.SelectedMap.Name).Msg("Game")

	game.sync_manger.Start(0)

	audio_player.SetVolume(0.1)
	audio_player.Play(0)

	rl.DisableCursor()

	return &game
}

func (game GameScene) Update(dt float64) {
	go game.sync_manger.Update(dt)
	game.cursor.Update(dt, game.grid)
	game.note_manager.Update(dt)
}

func (game *GameScene) Draw() {
	rl.ClearBackground(rl.Black)

	rl.BeginMode3D(game.camera.RlCamera)

	game.note_renderer.DrawNotesSingle()
	game.grid.Draw()
	game.cursor.Draw()

	rl.EndMode3D()
}

func (scene GameScene) GetType() int {
	return scene.scene_type
}

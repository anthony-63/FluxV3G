package managers

import (
	"errors"
	"flux/game/nodes"
	"flux/game/settings"
	"flux/game/util"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rs/zerolog/log"
)

const NOTE_MESH_PATH = "data/.game/game/mesh.gltf"

type NoteRenderer struct {
	note_mesh rl.Model

	approach_time float64

	ToRender []nodes.Note

	sync_manager *SyncManager
}

func CreateNoteRenderer(sync *SyncManager) *NoteRenderer {
	renderer := NoteRenderer{
		sync_manager:  sync,
		ToRender:      []nodes.Note{},
		approach_time: settings.GSettings.Note.ApproachTime,
	}

	if _, err := os.Stat(NOTE_MESH_PATH); errors.Is(err, os.ErrNotExist) {
		log.Error().Str("path", NOTE_MESH_PATH).Msg("Failed to find note mesh")
		os.Exit(1)
	}

	log.Info().Msg("Loading note mesh...")
	renderer.note_mesh = rl.LoadModel(NOTE_MESH_PATH)

	log.Info().Msg("Done")

	return &renderer
}

func (renderer *NoteRenderer) DrawNotesSingle() {
	sync := renderer.sync_manager

	for _, note := range renderer.ToRender {
		note_time := note.CalculateTime(sync.RealTime, renderer.approach_time*sync.Speed)
		note_distance := note_time * settings.GSettings.Note.ApproachDistance

		rl.DrawModelEx(renderer.note_mesh, rl.Vector3{
			X: float32((note.X * 2) * util.VFCONV64),
			Y: float32((note.Y * 2) * util.VFCONV64),
			Z: -float32(note_distance),
		}, rl.Vector3Zero(), 0, rl.Vector3Divide(rl.NewVector3(util.VFCONV32, util.VFCONV32, util.VFCONV32), rl.Vector3One()), rl.White)
	}
}

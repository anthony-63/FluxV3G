package managers

import (
	"errors"
	"flux/game/nodes"
	"flux/game/settings"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rs/zerolog/log"
)

const NOTE_MESH_PATH = "data/.game/game/mesh.gltf"

type NoteRenderer struct {
	note_mesh rl.Mesh
	note_mat  rl.Material

	approach_time float64

	ToRender []nodes.Note

	sync_manager *SyncManager
}

func CreateNoteRenderer(sync *SyncManager) *NoteRenderer {
	renderer := NoteRenderer{
		sync_manager:  sync,
		ToRender:      []nodes.Note{},
		note_mat:      rl.LoadMaterialDefault(),
		approach_time: settings.GSettings.Note.ApproachTime,
	}

	if _, err := os.Stat(NOTE_MESH_PATH); errors.Is(err, os.ErrNotExist) {
		log.Error().Str("path", NOTE_MESH_PATH).Msg("Failed to find note mesh")
		os.Exit(1)
	}

	log.Info().Msg("Loading note mesh...")
	model := rl.LoadModel(NOTE_MESH_PATH)
	renderer.note_mesh = model.GetMeshes()[0]
	renderer.note_mat.Maps.Color = rl.White

	log.Info().Msg("Done")

	return &renderer
}

func (renderer *NoteRenderer) DrawNotesSingle() {
	sync := renderer.sync_manager

	transforms := make([]rl.Matrix, len(renderer.ToRender))
	for i, note := range renderer.ToRender {
		note_time := note.CalculateTime(sync.RealTime, renderer.approach_time*sync.Speed)
		note_distance := note_time * settings.GSettings.Note.ApproachDistance

		transforms[i] = rl.MatrixTranslate(
			float32((note.X * 2)),
			float32((note.Y * 2)),
			-float32(note_distance),
		)

		transforms[i] = rl.MatrixMultiply(transforms[i], rl.MatrixScale(1, 1, 1))

		colored_mat := renderer.note_mat
		colored_mat.Maps.Color = note.Color
		rl.DrawMesh(renderer.note_mesh, colored_mat, transforms[i])
	}
}

package nodes

import (
	"errors"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rs/zerolog/log"
)

type Sprite3D struct {
	Position rl.Vector3
	Rotation rl.Vector3
	Size     rl.Vector2

	model rl.Model

	loaded bool
}

func (sprite *Sprite3D) GenPlane(width float32, height float32, tex_path string) {
	log.Info().Msg("Generating sprite model(plane)")
	sprite.model = rl.LoadModelFromMesh(rl.GenMeshPlane(width, height, 1, 1))
	log.Info().Msg("Generated model")

	log.Info().Msg("Loading texture...")
	if _, err := os.Stat(tex_path); errors.Is(err, os.ErrNotExist) {
		sprite.loaded = false
		log.Error().Str("path", tex_path).Msg("Failed to find sprite texture.")
		os.Exit(1)
	}

	img := rl.LoadImage(tex_path)

	tex := rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	log.Info().Msg("Loaded texture")

	log.Info().Msg("Setting texture...")
	rl.SetMaterialTexture(sprite.model.Materials, rl.MapDiffuse, tex)
	log.Info().Msg("Done")
}

func (sprite *Sprite3D) Update(dt float64) {

}

func (sprite *Sprite3D) Draw() {
	rl.PushMatrix()
	rl.Rotatef(90, 1, 0, 0)
	// rl.DrawModelEx(sprite.model, sprite.Position, rl.Vector3Zero(), 0, rl.Vector3One(), rl.White)
	rl.DrawModel(sprite.model, sprite.Position, 1, rl.White)
	rl.PopMatrix()
}

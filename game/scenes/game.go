package scenes

import (
	"flux/game/nodes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameScene struct {
	camera nodes.Camera
	grid   nodes.Sprite3D
}

func CreateGameScene() *GameScene {
	game := GameScene{
		camera: nodes.NewCamera(rl.Vector3{
			X: 0,
			Y: 0,
			Z: 1.35,
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
	}

	game.grid.GenPlane(1, 1, "data/.game/game/grid.png")

	return &game
}

func (game GameScene) Update(dt float64) {

}

func (game *GameScene) Draw() {
	rl.ClearBackground(rl.Black)

	rl.BeginMode3D(game.camera.RlCamera)

	game.grid.Draw()

	rl.EndMode3D()
}

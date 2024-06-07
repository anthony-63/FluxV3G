package nodes

import rl "github.com/gen2brain/raylib-go/raylib"

type Cursor struct {
	sprite *Sprite3D

	ClampedPosition rl.Vector2
	RealPosition    rl.Vector2
}

func CreateCursor() *Cursor {
	cursor := Cursor{
		sprite: &Sprite3D{
			Position: rl.Vector3Zero(),
			Rotation: rl.Vector3{
				X: 90,
				Y: 0,
				Z: 0,
			},
			Size: rl.Vector2One(),
		},
	}

	cursor.sprite.GenPlane(0.1, 0.1, "data/.game/game/cursor.png")

	return &cursor
}

func (cursor *Cursor) Draw() {
	cursor.sprite.Draw()
}

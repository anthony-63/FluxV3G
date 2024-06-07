package nodes

import (
	"flux/game/settings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const CLAMP = ((6. - 0.525) / 2)

type Cursor struct {
	sprite *Sprite3D

	ClampedPosition rl.Vector2
	RealPosition    rl.Vector2

	last_mouse rl.Vector2
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
			Size: rl.Vector2{
				X: 0.5,
				Y: 0.5,
			},
		},
		last_mouse: rl.GetMousePosition(),
	}

	cursor.sprite.GenPlane("data/.game/game/cursor.png")

	return &cursor
}

func (cursor *Cursor) Draw() {
	cursor.sprite.Draw()
}

func (cursor *Cursor) Update(dt float64, grid *Sprite3D) {
	new_pos := rl.GetMousePosition()
	mouse_delta := rl.Vector2Subtract(new_pos, cursor.last_mouse)
	cursor.last_mouse = new_pos

	delta := rl.Vector2{X: mouse_delta.X * float32(settings.GSettings.Cursor.Sensitivity/100), Y: mouse_delta.Y * float32(settings.GSettings.Cursor.Sensitivity/100)}

	cursor.RealPosition.X += delta.X
	cursor.RealPosition.Y += delta.Y

	cursor.RealPosition.X = rl.Clamp(cursor.RealPosition.X, -CLAMP, CLAMP)
	cursor.RealPosition.Y = rl.Clamp(cursor.RealPosition.Y, -CLAMP, CLAMP)

	cursor.sprite.Position = rl.Vector3{
		X: cursor.RealPosition.X,
		Z: cursor.RealPosition.Y,
		Y: 0,
	}
}

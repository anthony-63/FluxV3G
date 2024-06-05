package nodes

import (
	"flux/game/util"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const HIT_WINDOW float64 = 0.055
const AABB float64 = (1.75 + 0.525) / 2

type Note struct {
	X float64
	Y float64
	T float64

	Hit   bool
	Index uint
	Color rl.Color
}

func MakeNote(x float64, y float64, time float64, index uint, color rl.Color) *Note {
	note := Note{
		X:     x,
		Y:     y,
		T:     time,
		Index: index,
		Color: color,
	}

	return &note
}

func (note Note) InHitWindow(current_time float64, speed float64) bool {
	return (current_time - note.T) <= HIT_WINDOW*speed
}

func (note Note) IsVisible(current_time float64, speed float64, approach_time float64, pushback bool) bool {
	if note.Hit {
		return false
	}
	if current_time > note.T && !pushback {
		return false
	}

	return note.CalculateTime(current_time, approach_time) <= 1 && note.InHitWindow(current_time, speed)
}

func (note Note) CalculateTime(current_time float64, approach_time float64) float64 {
	return (note.T - current_time) / approach_time
}

func (note Note) IsBeingHit(cursor_pos rl.Vector2) bool {
	return math.Abs((float64(cursor_pos.X)-note.X*2)*util.VFCONV64) <= AABB && math.Abs((float64(cursor_pos.Y)-note.Y*2)*util.VFCONV64) <= AABB
}

package managers

import (
	"flux/game/nodes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type NoteManager struct {
	OrderedNotes []nodes.Note

	next_note *nodes.Note
	last_note *nodes.Note

	approach_time float64
	skipped_notes uint

	NotesProcessing uint
	StartProcess    uint

	colors []rl.Color

	pushback bool
	started  bool
}

func CreateNoteManager() *NoteManager {
	manager := NoteManager{
		next_note: nil,
		last_note: nil,
	}

	return &manager
}

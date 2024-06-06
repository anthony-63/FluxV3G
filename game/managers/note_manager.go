package managers

import (
	"flux/game/nodes"
	"flux/game/settings"
	"flux/game/util"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rs/zerolog/log"
)

type NoteManager struct {
	OrderedNotes []nodes.Note

	next_note *nodes.Note
	last_note *nodes.Note

	approach_time float64
	skipped_notes uint

	NotesProcessing int
	StartProcess    int

	sync_manager *SyncManager
	renderer     *NoteRenderer

	colors []rl.Color

	pushback bool
	started  bool
}

func CreateNoteManager(sync_manager *SyncManager, renderer *NoteRenderer) *NoteManager {
	manager := NoteManager{
		next_note: nil,
		last_note: nil,

		approach_time: settings.GSettings.Note.ApproachTime,
		pushback:      settings.GSettings.Note.Pushback,

		skipped_notes: 0,
		started:       false,

		OrderedNotes:    []nodes.Note{},
		NotesProcessing: 0,
		StartProcess:    0,

		sync_manager: sync_manager,
		renderer:     renderer,

		colors: []rl.Color{
			rl.White,
		},
	}

	manager.Load()

	return &manager
}

func (manager *NoteManager) Update(dt float64) {
	manager.updateRender()
	manager.updateNotes()
}

func (manager *NoteManager) updateRender() {
	sync := manager.sync_manager

	to_render := []nodes.Note{}
	for i := manager.StartProcess; i <= len(manager.OrderedNotes); i++ {
		note := manager.OrderedNotes[i]
		if note.IsVisible(sync.RealTime, sync.Speed, manager.approach_time*sync.Speed, manager.pushback) {
			to_render = append(to_render, note)
		}

		if note.T > sync.RealTime+manager.approach_time*sync.Speed {
			break
		}
	}

	manager.NotesProcessing = len(to_render)
	manager.renderer.ToRender = to_render
}

func (manager *NoteManager) updateNotes() {
	sync := manager.sync_manager

	to_process := []nodes.Note{}
	for i := manager.StartProcess; i <= len(manager.OrderedNotes); i++ {
		note := manager.OrderedNotes[i]

		if note.CalculateTime(sync.RealTime, manager.approach_time*sync.Speed) <= 0 && !note.Hit {
			to_process = append(to_process, note)
		}

		if note.T > sync.RealTime+manager.approach_time*sync.Speed {
			break
		}
	}

	for _, note := range to_process {
		did_hitreg := false

		if false /* note.IsBeingHit(cursor_pos) */ {
			note.Hit = true
			did_hitreg = true

			// score stuff
		}

		if !note.Hit && !note.InHitWindow(sync.RealTime, sync.Speed) {
			did_hitreg = true
			note.Hit = true
		}

		if did_hitreg {
			manager.last_note = &note
			if note.Index < len(manager.OrderedNotes)-1 {
				manager.next_note = &manager.OrderedNotes[note.Index+1]
				manager.StartProcess += 1
			} else if note.Index >= len(manager.OrderedNotes)-1 {
				manager.next_note = nil
			}
		}
	}
}

func (manager *NoteManager) Load() {
	for i, ndata := range util.SelectedMap.Notes {
		if ndata.T < util.StartFrom {
			manager.skipped_notes += 1
			continue
		}

		manager.OrderedNotes = append(manager.OrderedNotes, *nodes.MakeNote(ndata.X, ndata.Y, ndata.T, i, manager.colors[i%len(manager.colors)]))
	}

	if len(manager.OrderedNotes) > 0 {
		manager.next_note = &manager.OrderedNotes[0]
	}

	log.Info().Int("notes built", len(manager.OrderedNotes)).Msg("NoteManager")
}

package settings

var GSettings Settings = Settings{
	Note: NoteSettings{
		ApproachRate:     36.8,
		ApproachDistance: 14,
		ApproachTime:     14 / 36.8,
		Pushback:         false,
	},
	Cursor: CursorSettings{
		Scale:       1,
		Sensitivity: 0.005,
	},
}

type NoteSettings struct {
	ApproachRate     float64
	ApproachDistance float64
	ApproachTime     float64
	Pushback         bool
}

type CursorSettings struct {
	Scale       float64
	Sensitivity float64
}

type Settings struct {
	Note   NoteSettings
	Cursor CursorSettings
}

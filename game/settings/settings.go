package settings

var GSettings Settings = Settings{
	Note: NoteSettings{
		ApproachRate:     36.8,
		ApproachDistance: 14,
		ApproachTime:     14 / 36.8,
	},
}

type NoteSettings struct {
	ApproachRate     float64
	ApproachDistance float64
	ApproachTime     float64
}

type Settings struct {
	Note NoteSettings
}

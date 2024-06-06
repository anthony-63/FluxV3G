package settings

import "flux/game/util"

var GSettings Settings = Settings{
	Note: NoteSettings{
		ApproachRate:     36.8 * util.VFCONV64,
		ApproachDistance: 14 * util.VFCONV64,
		ApproachTime:     (14 * util.VFCONV64) / (36.8 * util.VFCONV64),
		Pushback:         false,
	},
}

type NoteSettings struct {
	ApproachRate     float64
	ApproachDistance float64
	ApproachTime     float64
	Pushback         bool
}

type Settings struct {
	Note NoteSettings
}

package util

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rs/zerolog/log"
)

func GetAudioFormat(data []byte) string {
	if len(data) < 10 {
		return "unknown"
	}

	if data[0] == 0x52 && data[1] == 0x49 && data[2] == 0x46 && data[3] == 0x46 {
		return ".wav"
	}

	if (data[0] == 0xff && (data[1] == 0xfb || (data[1] == 0xfa && data[2] == 0x90))) || (data[0] == 0x49 && data[1] == 0x44 && data[2] == 0x33) {
		return ".mp3"
	}

	if data[0] == 0x4f && data[1] == 0x67 && data[2] == 0x67 && data[3] == 0x53 {
		return ".ogg"
	}

	return "unknown"
}

func AudioPlayerFromFile(path string) rl.Music {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load audio file")
		os.Exit(0)
	}

	log.Info().Int("data size", len(data)).Msg("AudioLoader")
	return rl.LoadMusicStreamFromMemory(".mp3", data, int32(len(data)))
}

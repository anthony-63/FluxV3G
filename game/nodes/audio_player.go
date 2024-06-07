package nodes

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/rs/zerolog/log"
)

type AudioPlayer struct {
	music  rl.Music
	stream []byte
}

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

func AudioPlayerFromFile(path string) AudioPlayer {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load audio file")
		os.Exit(0)
	}

	player := AudioPlayer{}
	player.stream = data
	player.music = rl.LoadMusicStreamFromMemory(".mp3", player.stream, int32(len(player.stream)))
	return player
}

func (player *AudioPlayer) SetVolume(norm_vol float32) {
	rl.SetMusicVolume(player.music, norm_vol)
}

func (player *AudioPlayer) Play(from float32) {
	rl.SeekMusicStream(player.music, from)
	rl.PlayMusicStream(player.music)
}

func (player *AudioPlayer) Update() {
	rl.UpdateMusicStream(player.music)
}

func (player *AudioPlayer) ShouldSync(current_time float64) bool {
	return float32(current_time)-rl.GetMusicTimePlayed(player.music) > 0.2
}

func (player *AudioPlayer) Sync(current_time float64) {
	if player.ShouldSync(current_time) {
		log.Info().Float64("to", current_time).Msg("AudioPlayerSync")
		rl.SeekMusicStream(player.music, float32(current_time))
	}
}

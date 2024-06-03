package loaders

import (
	"errors"
	"flux/content"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog/log"
)

const MAP_DIR = "data/maps"

var LoadedMaps []content.BeatmapSet

func LoadMaps() {
	map_files, err := ioutil.ReadDir(MAP_DIR)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Log().Msg("Creating map directory.")
			os.Mkdir(MAP_DIR, os.ModeDir)
			return
		} else {
			log.Error().Msg("Failed to open map folder")
			return
		}
	}

	fmt.Println("Started map loading...")
	timestamp := time.Now().UnixMilli()
	for _, file := range map_files {
		if file.IsDir() {
			if loaded, err := content.GetBeatmapSetFromFolder(filepath.Join(MAP_DIR, file.Name())); err == nil {
				LoadedMaps = append(LoadedMaps, *loaded)
			} else {
				log.Error().Msg("Failed to load map\n" + err.Error())
			}
		}

	}

	fmt.Print("Loaded ", len(map_files), " maps in ", time.Now().UnixMilli()-timestamp, "ms\n")
}

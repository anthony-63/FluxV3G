package content

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/buger/jsonparser"
)

type BeatmapSet struct {
	broken  bool
	version uint8

	hash string
	path string

	artist  string
	title   string
	mappers []string

	music_path string

	has_cover bool
	cover     []uint8

	difficulties []Beatmap
}

func GetBeatmapSetFromFolder(path string) (*BeatmapSet, error) {
	set := BeatmapSet{}

	if _, err := os.Stat(path + "/meta.json"); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("%s: meta.json doesnt exist in directory", path)
	}

	data, err := ioutil.ReadFile(path + "/meta.json")
	if err != nil {
		return nil, fmt.Errorf("%s: failed to read meta.json", path)
	}

	if version, err := jsonparser.GetInt(data, "_version"); err != nil {
		return nil, fmt.Errorf("%s: failed to get version", path)
	} else {
		set.version = uint8(version)
	}

	if set.title, err = jsonparser.GetString(data, "_title"); err != nil {
		return nil, fmt.Errorf("%s: failed to get title", path)
	}

	if set.artist, err = jsonparser.GetString(data, "_artist"); err != nil {
		set.artist = ""
	}

	_, err = jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		mp, _ := GetBeatmapFromFile(filepath.Join(path, string(value)))
		if mp != nil {
			set.difficulties = append(set.difficulties, *mp)
		}
	}, "_difficulties")

	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse difficulties(or map has none)", path)
	}

	return &set, nil
}

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
	Broken  bool
	Version uint8

	Hash string
	Path string

	Artist  string
	Title   string
	Mappers []string

	MusicPath string

	HasCover bool
	Cover    []uint8

	Difficulties []Beatmap
}

func GetBeatmapSetFromFolder(path string) (*BeatmapSet, error) {
	set := BeatmapSet{}

	if _, err := os.Stat(path + "/meta.json"); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("%s: meta.json doesnt exist in directory", path)
	}

	set.Path = path

	data, err := ioutil.ReadFile(path + "/meta.json")
	if err != nil {
		return nil, fmt.Errorf("%s: failed to read meta.json", path)
	}

	if version, err := jsonparser.GetInt(data, "_version"); err != nil {
		return nil, fmt.Errorf("%s: failed to get version", path)
	} else {
		set.Version = uint8(version)
	}

	if set.Title, err = jsonparser.GetString(data, "_title"); err != nil {
		return nil, fmt.Errorf("%s: failed to get title", path)
	}

	if set.Artist, err = jsonparser.GetString(data, "_artist"); err != nil {
		set.Artist = ""
	}

	if set.MusicPath, err = jsonparser.GetString(data, "_music"); err != nil {
		return nil, fmt.Errorf("%s: failed to get music path", path)
	}

	_, err = jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		mp, _ := GetBeatmapFromFile(filepath.Join(path, string(value)))
		if mp != nil {
			set.Difficulties = append(set.Difficulties, *mp)
		}
	}, "_difficulties")

	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse difficulties(or map has none)", path)
	}

	return &set, nil
}

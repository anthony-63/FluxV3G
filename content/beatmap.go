package content

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/buger/jsonparser"
)

type NoteData struct {
	X float64
	Y float64
	T float64
}

type Beatmap struct {
	Broken  bool
	Version uint8

	Path string
	Name string
	ID   string

	Notes []NoteData
}

func GetBeatmapFromFile(path string) (*Beatmap, error) {
	mp := Beatmap{}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("%s: difficulty doesnt exist", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to read map", path)
	}

	if version, err := jsonparser.GetInt(data, "_version"); err != nil {
		return nil, fmt.Errorf("%s: failed to get version", path)
	} else {
		mp.Version = uint8(version)
	}

	if mp.Name, err = jsonparser.GetString(data, "_name"); err != nil {
		return nil, fmt.Errorf("%s: failed to get name", path)
	}

	_, err = jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		ndata := NoteData{}
		ndata.X, _ = jsonparser.GetFloat(value, "_x")
		ndata.Y, _ = jsonparser.GetFloat(value, "_y")
		ndata.T, _ = jsonparser.GetFloat(value, "_time")
		mp.Notes = append(mp.Notes, ndata)
	}, "_notes")

	sort.Slice(mp.Notes[:], func(i, j int) bool {
		return mp.Notes[i].T < mp.Notes[j].T
	})

	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse notes(or difficulty has none)", path)
	}

	return &mp, nil
}

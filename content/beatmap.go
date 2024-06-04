package content

import (
	"errors"
	"fmt"
	"os"

	"github.com/buger/jsonparser"
)

type NoteData struct {
	x float64
	y float64
	t float64
}

type Beatmap struct {
	broken  bool
	version uint8

	path string
	name string
	id   string

	notes []NoteData
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
		mp.version = uint8(version)
	}

	if mp.name, err = jsonparser.GetString(data, "_name"); err != nil {
		return nil, fmt.Errorf("%s: failed to get name", path)
	}

	_, err = jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		ndata := NoteData{}
		ndata.x, _ = jsonparser.GetFloat(value, "_x")
		ndata.y, _ = jsonparser.GetFloat(value, "_y")
		ndata.t, _ = jsonparser.GetFloat(value, "_time")
		mp.notes = append(mp.notes, ndata)
	}, "_notes")

	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse notes(or difficulty has none)", path)
	}

	return &mp, nil
}

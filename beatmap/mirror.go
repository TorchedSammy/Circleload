package beatmap

import (
	"errors"
	"net/http"
)

var (
	ErrMapsetNotFound = errors.New("mapset not found")
	ErrBeatmapNotFound = errors.New("map not found")
)

type Mirror interface {
	GetMapset(id int) (Mapset, error)
	GetMapsetFromMap(id int) (Mapset, error)
	Search(query string) ([]Mapset, error)
	GetMapsetData(id int) (*http.Response, error)
	SetMode(mode Mode)
}

type Options struct {
	NoVideo bool
	MaxResults int
	Mode Mode
}


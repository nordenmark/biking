package parser

import (
	"github.com/tkrajina/gpxgo/gpx"
)

type Session struct {
	Points []Point `json:"points"`
}

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Parse every byte slice (containing GPX data in XML format) into a Session struct
func ParseFiles(files [][]byte) []Session {
	sessions := make([]Session, len(files))

	for i, file := range files {
		gpxFile, err := gpx.ParseBytes(file)
		if err != nil {
			continue
		}

		points := []Point{}

		for _, track := range gpxFile.Tracks {
			for _, segment := range track.Segments {
				for _, point := range segment.Points {
					points = append(points, Point{
						Lat: point.Latitude,
						Lng: point.Longitude,
					})
				}
			}
		}

		sessions[i] = Session{
			Points: points,
		}
	}

	return sessions
}

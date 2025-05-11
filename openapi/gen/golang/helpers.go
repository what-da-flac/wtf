package golang

import "strconv"

func (x PatchV1AudioFilesIdJSONRequestBody) Map() map[string]string {
	m := make(map[string]string)
	if x.Album != "" {
		m["album"] = x.Album
	}
	if x.Genre != "" {
		m["genre"] = x.Genre
	}
	if x.Performer != "" {
		m["artist"] = x.Performer
	}
	if val := x.RecordedDate; val != nil && *val > 0 {
		m["date"] = strconv.Itoa(*val)
	}
	if x.Title != "" {
		m["title"] = x.Title
	}
	if val := x.TrackNumber; val != nil && *val > 0 {
		m["track"] = strconv.Itoa(*val)
	}
	return m
}

func (x PatchV1AudioFilesIdJSONRequestBody) DBMap() map[string]any {
	m := make(map[string]any)
	if x.Album != "" {
		m["album"] = x.Album
	}
	if x.Genre != "" {
		m["genre"] = x.Genre
	}
	if x.Performer != "" {
		m["performer"] = x.Performer
	}
	if val := x.RecordedDate; val != nil && *val > 0 {
		m["recorded_date"] = strconv.Itoa(*val)
	}
	if x.Title != "" {
		m["title"] = x.Title
	}
	if val := x.TrackNumber; val != nil && *val > 0 {
		m["track_number"] = strconv.Itoa(*val)
	}
	return m
}

func (x PathName) String() string { return string(x) }

package commands

import "io"

func CmdMediaInfo(filename string) (io.Reader, error) {
	var args = []string{
		"--Output=JSON",
		filename,
	}
	buf, err := runGeneric("", "mediainfo", args...)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

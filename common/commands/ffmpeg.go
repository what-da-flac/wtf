package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/what-da-flac/wtf/common/ifaces"
)

func CmdFFMpegAudio(input, output string, bitRate int) error {
	//	ffmpeg -y -i "$input" -map 0:a -c:a aac -b:a 320k "$output"
	var args = []string{
		"-y",
		"-i",
		input,
		"-map",
		"0:a",
		"-c:a",
		"aac",
		"-b:a",
		fmt.Sprintf("%d", bitRate),
		output,
	}
	_, err := runGeneric("", "ffmpeg", args...)
	return err
}

func CmdFFMpegSetTags(identifier ifaces.Identifier, filename string, values map[string]string) error {
	/*
		Commonly Supported Tags in .m4a / MP4 via ffmpeg
			Tag Key	Description	Notes
			title	Track title	Maps to ©nam atom
			artist	Main artist	©ART
			album	Album name	©alb
			genre	Genre	©gen
			date	Release date	©day
			track	Track number	e.g. 3/10
			disk	Disc number	e.g. 1/2
			composer	Composer	©wrt
			comment	Comment	©cmt
			album_artist	Album artist	aART
			encoder	Encoding software info
			lyrics	Lyrics	May map to ©lyr
			compilation	Compilation flag	Usually 1 or 0
	*/
	// ffmpeg -i input.m4a -metadata title="Beer is good" -codec copy output.m4a
	ext := filepath.Ext(filename)
	tmpFilename := filepath.Join(os.TempDir(), identifier.UUIDv4()+ext)
	var args = []string{
		"-y",
		"-i",
		filename,
	}
	for k, v := range values {
		args = append(args, "-metadata")
		args = append(args, k+"="+v)
	}
	args = append(args, "-codec", "copy", tmpFilename)
	if _, err := runGeneric("", "ffmpeg", args...); err != nil {
		return err
	}
	// copy destination to source filename
	if err := os.Rename(tmpFilename, filename); err != nil {
		return err
	}
	return os.RemoveAll(tmpFilename)
}

package commands

import (
	"fmt"
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

package commands

func CmdFFMpegAudio(input, output string) error {
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
		"320k",
		output,
	}
	_, err := runGeneric("", "ffmpeg", args...)
	return err
}

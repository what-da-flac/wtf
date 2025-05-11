package commands

import (
	"bytes"
	"os/exec"
)

func runGeneric(directory, name string, params ...string) (*bytes.Buffer, error) {
	// Create the command with the provided name and parameters
	cmd := exec.Command(name, params...)

	// Set the working directory if specified
	if directory != "" {
		cmd.Dir = directory
	}

	// Capture output and error streams
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	// Execute the command
	err := cmd.Run()
	if err != nil {
		// Capture both stdout and stderr for better error reporting
		return nil, err
	}
	// Return the output as a buffer
	return &stdoutBuf, nil
}

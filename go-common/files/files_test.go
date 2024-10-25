package files

import (
	"embed"
)

//go:embed test-data/*
var filesFs embed.FS

//go:embed test-data/1
var filesFs1 embed.FS

//go:embed test-data/2
var filesFs2 embed.FS

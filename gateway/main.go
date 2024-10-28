package main

import (
	_ "github.com/lib/pq"
	"github.com/what-da-flac/wtf/gateway/cmd"
)

func main() {
	cmd.Execute()
}

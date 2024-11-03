package main

import (
	_ "github.com/lib/pq"
	"github.com/what-da-flac/wtf/services/gateway/cmd"
)

func main() {
	cmd.Execute()
}

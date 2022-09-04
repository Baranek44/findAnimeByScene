package main

import (
	"fmt"
	"os"

	"github.com/baranek44/findAniemByScene/engine"
)

func main() {
	engine.UploadAllCommands()

	if err := engine.Routes.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

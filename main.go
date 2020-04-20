package main

import (
	"flag"
	"fmt"

	"github.com/HoloLabInc/go-symlink-creator/internal/settings"
	"github.com/HoloLabInc/go-symlink-creator/internal/symlinker"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if (len(args)) < 1 {
		fmt.Println("Please specify settings file")
		return
	}

	path := args[0]
	s := settings.LoadSettings(path)

	for _, setting := range s.SymLinkSettings {
		symlinker.CreateLink(setting)
	}
}

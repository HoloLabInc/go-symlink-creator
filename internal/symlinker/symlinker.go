package symlinker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/HoloLabInc/go-symlink-creator/internal/settings"
)

func CreateLink(s settings.SymLinkSetting) {
	for _, dest := range s.Dest {
		for _, target := range s.Target {
			createSymlink(s.BasePath, s.Src, dest, target, s.CreateDestFolder)
			if s.IncludeMeta {
				metaTarget := strings.TrimRight(target, "/\\") + ".meta"
				createSymlink(s.BasePath, s.Src, dest, metaTarget, s.CreateDestFolder)
			}
		}
	}
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func printError(msg string) {
	fmt.Printf("\x1b[31m%s\x1b[0m", msg)
	fmt.Println()
}

func createSymlink(base string, src string, dest string, target string, createFolder bool) {
	basedir := filepath.Dir(base)

	t := filepath.Join(basedir, src, target)
	d := filepath.Join(basedir, dest, target)
	destdir := filepath.Dir(d)

	if !exists(t) {
		msg := fmt.Sprintf("Target file does not exist: %s", t)
		printError(msg)
		return
	}

	if exists(d) {
		fmt.Println("Destination file exists:", d)
		return
	}

	if !exists(destdir) {
		if createFolder {
			if err := os.MkdirAll(destdir, 0777); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			msg := fmt.Sprintf("Destination folder does not exist: %s", destdir)
			printError(msg)
			return
		}
	}

	fmt.Printf("Create symlink from %s to %s\n", t, d)
	err := os.Symlink(t, d)
	if err != nil {
		fmt.Println(err)
	}
}

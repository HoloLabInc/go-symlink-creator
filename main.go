package main

import (
	"flag"
	"fmt"

	"github.com/HoloLabInc/go-symlink-creator/internal/symlinker"

	"github.com/HoloLabInc/go-symlink-creator/internal/settings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	if (len(args)) < 1 {
		fmt.Println("Please specify settings file")
		return
	}

	path := args[0]
	/*
		// ファイルをOpenする
		filepath := args[0]
		f, err := os.Open(filepath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		// 一気に全部読み取り
		b, err := ioutil.ReadAll(f)

		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	fmt.Println("hello world")
	s := settings.LoadSettings(path)

	fmt.Printf("--- m:\n%v\n\n", s)

	for _, setting := range s.SymLinkSettings {
		symlinker.CreateLink(setting)
	}
}

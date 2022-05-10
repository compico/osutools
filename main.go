package main

import (
	"log"

	"github.com/compico/osutools/filehelper"
)

var fh filehelper.OsuFolder

func main() {
	if err := fh.InitGamePathByReg(); err != nil {
		log.Fatalln(err)
	}
	err := fh.ReadOsudbFile()
	if err != nil {
		panic(err)
	}
}

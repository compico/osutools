package main

import (
	"log"

	"github.com/compico/osutools/filehelper"
)

func main() {
	var fh filehelper.OsuFolder
	if err := fh.InitGamePathByReg(); err != nil {
		log.Fatalln(err)
	}

	if err := fh.GetAllPaths(); err != nil {
		log.Fatalln(err)
	}

	if err := fh.ReadOsudbFile(); err != nil {
		log.Fatalln(err)
	}

	// f, err := os.Create("./osu!.db.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// bjson, err := json.Marshal(fh.DataBase)
	// if err != nil {
	// 	panic(err)
	// }
	// f.Write(bjson)
}

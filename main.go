package main

import (
	"log"

	"github.com/compico/osutools/filehelper"
	"github.com/compico/osutools/internal/webserver"
	"github.com/julienschmidt/httprouter"
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
	r := httprouter.New()
	server := webserver.NewServer(":8000", r)
	if err := server.S.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/compico/osutools/filehelper"
	"github.com/compico/osutools/internal/webserver"
	"github.com/julienschmidt/httprouter"
)

var fh filehelper.OsuFolder

func main() {
	if err := fh.InitGamePathByReg(); err != nil {
		log.Fatalln(err)
	}
	if err := fh.ReadOsudbFile(); err != nil {
		log.Fatalln(err)
	}

	r := httprouter.New()
	server := webserver.NewServer(":8000", r)

	r.GET("/", appHandler)
	r.GET("/doms/interface", domInterfaceHandler)

	r.ServeFiles("/static/images/*filepath", http.Dir("./static/images/"))
	r.ServeFiles("/static/css/*filepath", http.Dir("./static/css/"))
	r.ServeFiles("/static/js/*filepath", http.Dir("./static/js/"))
	r.ServeFiles("/static/songs/*filepath", http.Dir(fh.SongsPath))

	fmt.Println("Webserver starting on http://localhost:8000/")
	if err := server.S.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

}

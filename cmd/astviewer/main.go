package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/akito0107/astviewer/statik"
	"github.com/rakyll/statik/fs"
)

var port = flag.Int("p", 8080, "listen port")

func main() {
	flag.Parse()

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	p := fmt.Sprintf(":%d", *port)

	http.ListenAndServe(p, http.FileServer(statikFS))
}

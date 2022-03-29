package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/akito0107/astviewer"
)

var port = flag.Int("p", 8080, "listen port")

func main() {
	flag.Parse()
	p := fmt.Sprintf(":%d", *port)

	log.Fatal(http.ListenAndServe(p, http.FileServer(http.FS(astviewer.PublicFS))))
}

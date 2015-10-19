package main

import (
	"flag"
	"fmt"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"net/http"
)

func main() {
	flag.Set("bind", ":3000")

	goji.Get("/hello/:name", hello)
	goji.Serve()
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

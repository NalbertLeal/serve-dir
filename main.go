package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/NalbertLeal/serve-dir/internal"
)

func buildFileServer(dontList bool, noIndex bool) *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(internal.CurrentDir))
	handler := internal.BuildMiddlewareChain(fileServer, noIndex, dontList)
	mux.Handle("/", handler)

	return mux
}

func main() {
	NoList := flag.Bool("no-list", false, "Don't list the files inside directors")
	NoIndex := flag.Bool("no-index", false, "Don't return index.html as default")

	flag.Parse()

	mux := buildFileServer(*NoList, *NoIndex)

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatalln(err)
	}
}

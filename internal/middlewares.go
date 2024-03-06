package internal

import (
	"net/http"
	"path"
	"strings"
)

func DontListFilesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CANT IGNORE index.html
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func IgnoreIndexHtmlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			dirPath := path.Join(CurrentDir, r.URL.Path)
			htmlList := CreateDirHtml(dirPath)

			w.Write([]byte(htmlList))
		} else {
			filePath := path.Join(CurrentDir, r.URL.Path)
			http.ServeFile(w, r, filePath)
		}
	})
}

func BuildMiddlewareChain(fileServer http.Handler, noIndex bool, dontList bool) http.Handler {
	handler := fileServer
	if noIndex {
		handler = IgnoreIndexHtmlMiddleware(handler)
	}
	if dontList {
		handler = DontListFilesMiddleware(handler)
	}
	return handler
}

package main

import (
	slog "go-web-example/server/log"
	"log"
	"net/http"
)

func main() {
	slog.SetStdLogger("GoWebExample")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("RequestURI: %v", r.RequestURI)
		w.Write([]byte(r.RequestURI))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

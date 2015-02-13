package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shaked/gomobiledetect"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%#v", mobiledetect.Device(r))
}

func main() {
	log.Println("Starting local server http://localhost:10001/check (cmd+click to open from terminal)")
	mux := http.NewServeMux()
	h := &Handler{}
	mux.Handle("/check", h)
	http.ListenAndServe(":10001", mobiledetect.HandlerMux(mux, nil))
}

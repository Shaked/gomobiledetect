package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shaked/gomobiledetect"
)

func handler(w http.ResponseWriter, r *http.Request) {
	detect := gomobiledetect.NewMobileDetect(r, nil)
	requestValue := r.URL.Query().Get("r")
	fmt.Fprintln(w, "isMobile?", detect.IsMobile())
	fmt.Fprintln(w, "isTablet?", detect.IsTablet())
	fmt.Fprintln(w, "is(request)?", requestValue, " ", detect.Is(requestValue))
	fmt.Fprintln(w, "Version: ", detect.Version(requestValue))
}

func main() {
	log.Println("Starting local server http://localhost:10001/check (cmd+click to open from terminal)")
	http.HandleFunc("/check", handler)
	http.ListenAndServe(":10001", nil)
}

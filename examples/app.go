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
	fmt.Fprintln(w, "isKey(request)?", requestValue, " ", detect.IsKey(gomobiledetect.IPHONE))
	fmt.Fprintln(w, "Version: ", detect.Version(requestValue))
	fmt.Fprintln(w, "VersionKey: ", detect.Version(gomobiledetect.PROP_IPHONE))
	fmt.Fprintln(w, "VersionFloat: ", detect.Version(requestValue))
	fmt.Fprintln(w, "VersionFloatKey: ", detect.Version(gomobiledetect.PROP_IPHONE))
}

func main() {
	log.Println("Starting local server http://localhost:10001/check (cmd+click to open from terminal)")
	http.HandleFunc("/check", handler)
	http.ListenAndServe(":10001", nil)
}

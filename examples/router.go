package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/Shaked/gomobiledetect"
)

//route manager
type route struct {
	re      *regexp.Regexp
	handler func(http.ResponseWriter, *http.Request, []string, *mobiledetect.MobileDetect)
}

type RouterHandler struct {
	routes []*route
	detect *mobiledetect.MobileDetect
}

func (h *RouterHandler) AddRoute(re string, handler func(http.ResponseWriter, *http.Request, []string, *mobiledetect.MobileDetect)) {
	r := &route{regexp.MustCompile(re), handler}
	h.routes = append(h.routes, r)
}

func (h *RouterHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.detect = mobiledetect.NewMobileDetect(r, nil)
	for _, route := range h.routes {
		matches := route.re.FindStringSubmatch(r.URL.String())
		if matches != nil {
			route.handler(rw, r, matches, h.detect)
			break
		}
	}
}

func homepageHandler(w http.ResponseWriter, r *http.Request, matches []string, detect *mobiledetect.MobileDetect) {
	fmt.Fprint(w, "Hello World\n")
	fmt.Fprintf(w, "Matches %+v\n", matches)
	fmt.Fprintf(w, "Is Mobile? %+v\n", detect.IsMobile())
	fmt.Fprintf(w, "Is Tablet? %+v\n", detect.IsTablet())
}

func main() {
	reHandler := new(RouterHandler)
	reHandler.AddRoute("/device/[mobile|desktop]", homepageHandler)
	http.ListenAndServe(":9999", reHandler)
}

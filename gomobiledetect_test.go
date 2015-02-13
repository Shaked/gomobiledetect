package mobiledetect

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var httpRequest = &http.Request{}

type basicMethodsStruct struct {
	httpHeaders          map[string]string
	httpHeadersForMobile bool
	isUserAgent          string
	isMobile             bool
	isTablet             bool
	customValues         []basicMethodsStructCustomValue

	handlerCalled string
}

type basicMethodsStructCustomValue struct {
	name  string
	key   int
	value bool
}

func (h *basicMethodsStruct) Mobile(w http.ResponseWriter,
	r *http.Request,
	m *MobileDetect,
) {
	h.handlerCalled = "mobile"
}
func (h *basicMethodsStruct) Tablet(w http.ResponseWriter,
	r *http.Request,
	m *MobileDetect,
) {
	h.handlerCalled = "tablet"
}
func (h *basicMethodsStruct) Desktop(w http.ResponseWriter,
	r *http.Request,
	m *MobileDetect,
) {
	h.handlerCalled = "desktop"
}

func (h *basicMethodsStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handlerCalled = Device(r)
}

func BasicMethodsData() []basicMethodsStruct {
	return []basicMethodsStruct{
		basicMethodsStruct{
			httpHeaders: map[string]string{
				"SERVER_SOFTWARE":       "Apache/2.2.15 (Linux) Whatever/4.0 PHP/5.2.13",
				"REQUEST_METHOD":        "POST",
				"HTTP_HOST":             "home.ghita.org",
				"HTTP_X_REAL_IP":        "1.2.3.4",
				"HTTP_X_FORWARDED_FOR":  "1.2.3.5",
				"HTTP_CONNECTION":       "close",
				"HTTP_USER_AGENT":       "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A523 Safari/8536.25",
				"HTTP_ACCEPT":           "text/vnd.wap.wml, application/json, text/javascript, */*; q=0.01",
				"HTTP_ACCEPT_LANGUAGE":  "en-us,en;q=0.5",
				"HTTP_ACCEPT_ENCODING":  "gzip, deflate",
				"HTTP_X_REQUESTED_WITH": "XMLHttpRequest",
				"HTTP_REFERER":          "http://mobiledetect.net",
				"HTTP_PRAGMA":           "no-cache",
				"HTTP_CACHE_CONTROL":    "no-cache",
				"REMOTE_ADDR":           "11.22.33.44",
				"REQUEST_TIME":          "01-10-2012 07:57",
			},
			httpHeadersForMobile: false,
			isMobile:             true,
			isTablet:             false,
			customValues: []basicMethodsStructCustomValue{
				basicMethodsStructCustomValue{
					name:  "iphone",
					value: true,
				},
				basicMethodsStructCustomValue{
					name:  "ios",
					value: true,
				},
				basicMethodsStructCustomValue{
					name:  "whatever",
					value: false,
				},
			},
		},
		basicMethodsStruct{
			httpHeaders: map[string]string{
				"SERVER_SOFTWARE":       "Apache/2.2.15 (Linux) Whatever/4.0 PHP/5.2.13",
				"REQUEST_METHOD":        "POST",
				"HTTP_HOST":             "",
				"HTTP_X_REAL_IP":        "1.2.3.4",
				"HTTP_X_FORWARDED_FOR":  "1.2.3.5",
				"HTTP_CONNECTION":       "close",
				"HTTP_USER_AGENT":       "Mozilla/5.0",
				"HTTP_ACCEPT":           "application/json, text/javascript, */*; q=0.01",
				"HTTP_ACCEPT_LANGUAGE":  "en-us,en;q=0.5",
				"HTTP_ACCEPT_ENCODING":  "gzip, deflate",
				"HTTP_X_REQUESTED_WITH": "XMLHttpRequest",
				"HTTP_REFERER":          "",
				"HTTP_PRAGMA":           "no-cache",
				"HTTP_CACHE_CONTROL":    "no-cache",
				"REMOTE_ADDR":           "11.22.33.44",
				"REQUEST_TIME":          "01-10-2012 07:57",
			},
			httpHeadersForMobile: true,
			isMobile:             false,
			isTablet:             false,
			customValues: []basicMethodsStructCustomValue{
				basicMethodsStructCustomValue{
					name:  "iphone",
					value: false,
				},
				basicMethodsStructCustomValue{
					name:  "ios",
					value: false,
				},
				basicMethodsStructCustomValue{
					name:  "whatever",
					value: false,
				},
			},
		},
		basicMethodsStruct{
			httpHeaders: map[string]string{
				"SERVER_SOFTWARE":       "Apache/2.2.15 (Linux) Whatever/4.0 PHP/5.2.13",
				"REQUEST_METHOD":        "POST",
				"HTTP_HOST":             "",
				"HTTP_X_REAL_IP":        "1.2.3.4",
				"HTTP_X_FORWARDED_FOR":  "1.2.3.5",
				"HTTP_CONNECTION":       "close",
				"HTTP_USER_AGENT":       "Mozilla/5.0 (iPad; CPU OS 5_1_1 like Mac OS X; en-us) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/21.0.1180.80 Mobile/9B206 Safari/7534.48.3 (6FF046A0-1BC4-4E7D-8A9D-6BF17622A123)",
				"HTTP_ACCEPT":           "application/json, text/javascript, */*; q=0.01",
				"HTTP_ACCEPT_LANGUAGE":  "en-us,en;q=0.5",
				"HTTP_ACCEPT_ENCODING":  "gzip, deflate",
				"HTTP_X_REQUESTED_WITH": "XMLHttpRequest",
				"HTTP_REFERER":          "",
				"HTTP_PRAGMA":           "no-cache",
				"HTTP_CACHE_CONTROL":    "no-cache",
				"REMOTE_ADDR":           "11.22.33.44",
				"REQUEST_TIME":          "01-10-2012 07:57",
			},
			httpHeadersForMobile: true,
			isMobile:             true,
			isTablet:             true,
			customValues: []basicMethodsStructCustomValue{
				basicMethodsStructCustomValue{
					name:  "iphone",
					value: false,
				},
				basicMethodsStructCustomValue{
					name:  "ios",
					value: true,
				},
				basicMethodsStructCustomValue{
					name:  "whatever",
					value: false,
				},
			},
		},
		basicMethodsStruct{
			httpHeaders: map[string]string{
				"SERVER_SOFTWARE":       "Apache/2.2.15 (Linux) Whatever/4.0 PHP/5.2.13",
				"REQUEST_METHOD":        "POST",
				"HTTP_HOST":             "",
				"HTTP_X_REAL_IP":        "1.2.3.4",
				"HTTP_X_FORWARDED_FOR":  "1.2.3.5",
				"HTTP_CONNECTION":       "close",
				"HTTP_USER_AGENT":       "Mozilla/5.0 (iPad; CPU OS 5_1_1 like Mac OS X; en-us) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/21.0.1180.80 Mobile/9B206 Safari/7534.48.3 (6FF046A0-1BC4-4E7D-8A9D-6BF17622A123)",
				"HTTP_ACCEPT":           "application/json, text/javascript, */*; q=0.01",
				"HTTP_ACCEPT_LANGUAGE":  "en-us,en;q=0.5",
				"HTTP_ACCEPT_ENCODING":  "gzip, deflate",
				"HTTP_X_REQUESTED_WITH": "XMLHttpRequest",
				"HTTP_REFERER":          "",
				"HTTP_PRAGMA":           "no-cache",
				"HTTP_CACHE_CONTROL":    "no-cache",
				"REMOTE_ADDR":           "11.22.33.44",
				"REQUEST_TIME":          "01-10-2012 07:57",
			},
			httpHeadersForMobile: true,
			isMobile:             true,
			isTablet:             true,
			customValues: []basicMethodsStructCustomValue{
				basicMethodsStructCustomValue{
					key:   IPHONE,
					value: false,
				},
				basicMethodsStructCustomValue{
					key:   IOS,
					value: true,
				},
				basicMethodsStructCustomValue{
					key:   9999999,
					value: false,
				},
			},
		},
	}
}

func TestBasicMethods(t *testing.T) {
	detect := NewMobileDetect(httpRequest, nil)
	for _, data := range BasicMethodsData() {
		detect.SetHttpHeaders(data.httpHeaders)

		if 16 != len(detect.httpHeaders) {
			t.Error("Http headers were not set")
		}

		if data.httpHeadersForMobile == detect.CheckHttpHeadersForMobile() {
			t.Error("Http mobile headers check failed")
		}

		detect.SetUserAgent(data.httpHeaders["HTTP_USER_AGENT"])
		if data.isUserAgent == detect.userAgent {
			t.Error("User agent was not set")
		}

		if data.isMobile != detect.IsMobile() {
			t.Error("Mobile detection failed")
		}

		if data.isTablet != detect.IsTablet() {
			t.Error("Tablet detection failed")
		}

		for _, customValue := range data.customValues {
			if customValue.value != detect.Is(customValue.name) && customValue.value != detect.Is(customValue.key) {
				t.Errorf("Is(%s) detetction failed", customValue.name)
			}
		}
	}

	notSupported := detect.Is(1.0)
	if false != notSupported {
		t.Errorf("Type is not supported.")
	}

	r, _ := http.NewRequest("GET", "/", nil)
	emptyResult := Device(r)
	if "" != emptyResult {
		t.Errorf("Result is not empty: %s", emptyResult)
	}
}

//special headers that give `quick` indication that a device is mobile
func QuickHeadersData() []map[string]string {
	headers := []map[string]string{
		map[string]string{`HTTP_ACCEPT`: `application/json; q=0.2, application/x-obml2d; q=0.8, image/gif; q=0.99, */*`},
		map[string]string{`HTTP_ACCEPT`: `text/*; q=0.1, application/vnd.rim.html`},
		map[string]string{`HTTP_ACCEPT`: `text/vnd.wap.wml`},
		map[string]string{`HTTP_ACCEPT`: `application/vnd.wap.xhtml+xml`},
		map[string]string{`HTTP_X_WAP_PROFILE`: `hello`},
		map[string]string{`HTTP_X_WAP_CLIENTID`: ``},
		map[string]string{`HTTP_WAP_CONNECTION`: ``},
		map[string]string{`HTTP_PROFILE`: ``},
		map[string]string{`HTTP_X_OPERAMINI_PHONE_UA`: ``},
		map[string]string{`HTTP_X_NOKIA_GATEWAY_ID`: ``},
		map[string]string{`HTTP_X_ORANGE_ID`: ``},
		map[string]string{`HTTP_X_VODAFONE_3GPDPCONTEXT`: ``},
		map[string]string{`HTTP_X_HUAWEI_USERID`: ``},
		map[string]string{`HTTP_UA_OS`: ``},
		map[string]string{`HTTP_X_MOBILE_GATEWAY`: ``},
		map[string]string{`HTTP_X_ATT_DEVICEID`: ``},
		map[string]string{`HTTP_UA_CPU`: `ARM`},
	}
	return headers
}

func TestQuickHeaders(t *testing.T) {
	detect := NewMobileDetect(httpRequest, nil)
	detect.PreCompileRegexRules()
	for _, httpHeaders := range QuickHeadersData() {
		detect.SetHttpHeaders(httpHeaders)
		if true != detect.CheckHttpHeadersForMobile() {
			t.Errorf("Headers %+v failed", httpHeaders)
		}
	}
}

func QuickNonMobileHeadersData() []map[string]string {
	headers := []map[string]string{
		map[string]string{`HTTP_UA_CPU`: `AMD64`},
		map[string]string{`HTTP_UA_CPU`: `X86`},
		map[string]string{`HTTP_ACCEPT`: `text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01`},
		map[string]string{`HTTP_REQUEST_METHOD`: `DELETE`},
		map[string]string{`HTTP_VIA`: `1.1 ws-proxy.stuff.co.il C0A800FA`},
	}
	return headers
}

func TestNonMobileQuickHeaders(t *testing.T) {
	detect := NewMobileDetect(httpRequest, nil)
	for _, httpHeaders := range QuickNonMobileHeadersData() {
		detect.SetHttpHeaders(httpHeaders)
		if false != detect.CheckHttpHeadersForMobile() {
			t.Errorf("Headers %+v failed", httpHeaders)
		}
	}
}

type versionDataStruct struct {
	userAgent    string
	property     interface{}
	strVersion   string
	floatVersion float64
}

func VersionData() []versionDataStruct {
	v := []versionDataStruct{
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (Linux; Android 4.0.4; ARCHOS 80G9 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
			property:     `Android`,
			strVersion:   `4.0.4`,
			floatVersion: 4.04,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (Linux; Android 4.0.4; ARCHOS 80G9 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
			property:     `Webkit`,
			strVersion:   `535.19`,
			floatVersion: 535.19,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (Linux; Android 4.0.4; ARCHOS 80G9 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
			property:     `Chrome`,
			strVersion:   `18.0.1025.166`,
			floatVersion: 18.01025166,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`,
			property:     `BlackBerry`,
			strVersion:   `6.0.0.448`,
			floatVersion: 6.00448,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`,
			property:     `Webkit`,
			strVersion:   `534.8`,
			floatVersion: 534.8,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`,
			property:     `Webkit`,
			strVersion:   `534.8`,
			floatVersion: 534.8,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`,
			property:     `Unknown property`,
			strVersion:   ``,
			floatVersion: 0.0,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`,
			property:     struct{}{},
			strVersion:   ``,
			floatVersion: 0.0,
		},
		versionDataStruct{
			userAgent:    `Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`,
			property:     PROP_ANDROID,
			strVersion:   ``,
			floatVersion: 0.0,
		},
	}
	return v
}

//todo: check if this test is testing the code or testing that the data is correct
func TestVersionExtraction(t *testing.T) {
	detect := NewMobileDetect(httpRequest, nil)

	for _, data := range VersionData() {
		userAgent := data.userAgent
		strVersion := data.strVersion
		floatVersion := data.floatVersion
		property := data.property
		detect.SetUserAgent(userAgent)
		detectedVersion := detect.Version(property)
		if strVersion != detectedVersion {
			t.Errorf("String version %s is mismatched (detectedVersion %s, property %s)", strVersion, detectedVersion, property)
		}

		detectedVersionFloat := detect.VersionFloat(property)
		if floatVersion != detectedVersionFloat {
			t.Errorf("Float version %d is mismatched (detectedVersion %d, property %s)", floatVersion, detectedVersionFloat, property)
		}
	}
}

func TestPreCompileRegexRules(t *testing.T) {
	detect := NewMobileDetect(httpRequest, nil)
	detect.PreCompileRegexRules()
	e := len(detect.rules.mobileDetectionRules())
	c := len(detect.compiledRegexRules)
	if c != e {
		t.Errorf("Compiled rules are not being cached.\n Rules: %d\n Cached: %d\n", e, c)
	}
}

func TestHandler(t *testing.T) {
	expectedResults := map[string]string{
		"mobile":  `Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A4449d Safari/9537.53`,
		"tablet":  `Mozilla/5.0 (iPad; CPU OS 5_1_1 like Mac OS X; en-us) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/21.0.1180.80 Mobile/9B206 Safari/7534.48.3 (6FF046A0-1BC4-4E7D-8A9D-6BF17622A123)`,
		"desktop": "UNKNOWN",
	}

	deviceHandler := &basicMethodsStruct{}
	h := Handler(deviceHandler, nil)

	s := httptest.NewServer(h)
	req, _ := http.NewRequest("GET", s.URL, nil)
	c := http.Client{}
	for deviceType, userAgent := range expectedResults {
		req.Header.Set("User-Agent", userAgent)
		_, err := c.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		if deviceHandler.handlerCalled != deviceType {
			t.Errorf("actual: %s instead: %s", deviceHandler.handlerCalled, deviceType)
		}
	}
}

func TestHandlerMux(t *testing.T) {
	expectedResults := map[string]string{
		"Mobile":  `Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A4449d Safari/9537.53`,
		"Tablet":  `Mozilla/5.0 (iPad; CPU OS 5_1_1 like Mac OS X; en-us) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/21.0.1180.80 Mobile/9B206 Safari/7534.48.3 (6FF046A0-1BC4-4E7D-8A9D-6BF17622A123)`,
		"Desktop": "UNKNOWN",
	}
	mux := http.NewServeMux()
	deviceHandler := &basicMethodsStruct{}
	mux.Handle("/test", deviceHandler)
	s := httptest.NewServer(HandlerMux(mux, nil))
	req, _ := http.NewRequest("GET", s.URL+"/test", nil)
	c := http.Client{}
	for deviceType, userAgent := range expectedResults {
		req.Header.Set("User-Agent", userAgent)
		_, err := c.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		if deviceHandler.handlerCalled != deviceType {
			t.Errorf("actual: %s instead: %s", deviceHandler.handlerCalled, deviceType)
		}
	}
}

func BenchmarkIsMobile(b *testing.B) {
	req, _ := http.NewRequest("GET", "URL", strings.NewReader(""))
	detect := NewMobileDetect(req, nil)
	detect.SetUserAgent(`Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`)
	for n := 0; n < b.N; n++ {
		detect.IsMobile()
	}
}

func BenchmarkIs(b *testing.B) {
	req, _ := http.NewRequest("GET", "URL", strings.NewReader(""))
	detect := NewMobileDetect(req, nil)
	detect.SetUserAgent(`Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`)
	for n := 0; n < b.N; n++ {
		detect.Is("iphone")
	}
}
func BenchmarkIsKey(b *testing.B) {
	req, _ := http.NewRequest("GET", "URL", strings.NewReader(""))
	detect := NewMobileDetect(req, nil)
	detect.SetUserAgent(`Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`)
	for n := 0; n < b.N; n++ {
		detect.IsKey(IPHONE)
	}
}

func BenchmarkVersion(b *testing.B) {
	req, _ := http.NewRequest("GET", "URL", strings.NewReader(""))
	detect := NewMobileDetect(req, nil)
	detect.SetUserAgent(`Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`)
	for n := 0; n < b.N; n++ {
		detect.Version("iphone")
	}
}
func BenchmarkVersionKey(b *testing.B) {
	req, _ := http.NewRequest("GET", "URL", strings.NewReader(""))
	detect := NewMobileDetect(req, nil)
	detect.SetUserAgent(`Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`)
	for n := 0; n < b.N; n++ {
		detect.VersionKey(PROP_IPHONE)
	}
}

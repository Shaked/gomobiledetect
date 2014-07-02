// Gomobiledetect is a lightweight Go package imported from PHP for detecting mobile devices (including tablets). It uses the User-Agent string combined with specific HTTP headers to detect the mobile environment
package gomobiledetect

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	//A frequently used regular expression to extract version #s.
	verRegex = `([\w._\+]+)`

	MOBILE_GRADE_A = "A"
	MOBILE_GRADE_B = "B"
	MOBILE_GRADE_C = "C"
)

// MobileDetect holds the structure to figure out a browser from a UserAgent string and methods necessary to make it happen
type MobileDetect struct {
	rules                *rules
	userAgent            string
	httpHeaders          map[string]string
	mobileDetectionRules map[string]string
	compiledRegexRules   map[string]*regexp.Regexp
}

// NewMobileDetect creates the MobileDetect object
func NewMobileDetect(r *http.Request, rules *rules) *MobileDetect {
	if nil == rules {
		rules = NewRules()
	}
	md := &MobileDetect{
		rules:              rules,
		userAgent:          r.UserAgent(),
		httpHeaders:        getHttpHeaders(r),
		compiledRegexRules: make(map[string]*regexp.Regexp),
	}
	return md
}

func getHttpHeaders(r *http.Request) map[string]string {
	httpHeaders := map[string]string{
		"SERVER_SOFTWARE":  r.Header.Get("SERVER_SOFTWARE"),
		"REQUEST_METHOD":   r.Method,
		"HOST":             r.Host,
		"X_REAL_IP":        r.Header.Get("X_REAL_IP"),
		"X_FORWARDED_FOR":  r.Header.Get("X_FORWARDED_FOR"),
		"CONNECTION":       r.Header.Get("CONNECTION"),
		"USER-AGENT":       r.UserAgent(),
		"ACCEPT":           r.Header.Get("ACCEPT"),
		"ACCEPT-LANGUAGE":  r.Header.Get("ACCEPT-LANGUAGE"),
		"ACCEPT-ENCODING":  r.Header.Get("ACCEPT-ENCODING"),
		"X_REQUESTED_WITH": r.Header.Get("X_REQUESTED_WITH"),
		"REFERER":          r.Referer(),
		"PRAGMA":           r.Header.Get("PRAGMA"),
		"CACHE_CONTROL":    r.Header.Get("CACHE_CONTROL"),
		"REMOTE_ADDR":      r.RemoteAddr,
		"REQUEST_TIME":     r.Header.Get("REQUEST_TIME"),
	}

	return httpHeaders
}

func (md *MobileDetect) PreCompileRegexRules() *MobileDetect {
	for _, ruleValue := range md.rules.getMobileDetectionRules() {
		md.match(ruleValue)
	}
	return md
}

func (md *MobileDetect) SetUserAgent(userAgent string) *MobileDetect {
	md.userAgent = userAgent
	return md
}

func (md *MobileDetect) SetHttpHeaders(httpHeaders map[string]string) *MobileDetect {
	md.httpHeaders = httpHeaders
	return md
}

// IsMobile is a specific case to detect only mobile browsers.
func (md *MobileDetect) IsMobile() bool {
	if md.CheckHttpHeadersForMobile() {
		return true
	}
	return md.matchDetectionRulesAgainstUA()
}

// IsMobile is a specific case to detect only mobile browsers on tablets. Do not overlap with IsMobile
func (md *MobileDetect) IsTablet() bool {
	for _, rule := range md.rules.tabletDevices {
		if md.match(rule) {
			return true
		}
	}
	return false
}

// Is compared the detected browser with a "rule"
func (md *MobileDetect) Is(key string) bool {
	return md.matchUAAgainstKey(key)
}

// Version detects the browser version returning as string
func (md *MobileDetect) Version(propertyName string) string {
	if "" != propertyName {
		properties := md.Properties()

		if _, ok := properties[propertyName]; ok {
			for _, propertyMatchString := range properties[propertyName] {
				propertyPattern := `(?is)` + strings.Replace(string(propertyMatchString), `[VER]`, verRegex, -1)

				// Escape the special character which is the delimiter.
				//propertyPattern = strings.Replace(propertyPattern, `/`, `\/`, -1)

				// Identify and extract the version.
				re := regexp.MustCompile(propertyPattern)
				match := re.FindStringSubmatch(md.userAgent)
				if len(match) > 0 {
					return match[1]
				}
			}
		}
	}
	return ""
}

// VersionFloat does the same as Version, but returns a float number good for version comparison
func (md *MobileDetect) VersionFloat(propertyName string) float64 {
	version := md.Version(propertyName)
	replacer := strings.NewReplacer(`_`, `.`, `/`, `.`)
	version = replacer.Replace(version)

	versionNumbers := strings.Split(version, `.`)

	versionNumbersLength := len(versionNumbers)
	if versionNumbersLength > 1 {
		firstNumber := versionNumbers[0]
		retVersion := make([]string, (versionNumbersLength - 1))
		for i := 1; i < versionNumbersLength; i++ {
			retVersion[(i - 1)] = strings.Replace(versionNumbers[i], `.`, ``, -1)
		}

		version = firstNumber + `.` + strings.Join(retVersion, ``)
	}
	versionFloat, err := strconv.ParseFloat(version, 64)

	if nil != err {
		return 0.0
	}
	return versionFloat
}

//Search for a certain key in the rules array.
//If the key is found the try to match the corresponding regex agains the User-Agent.
func (md *MobileDetect) matchUAAgainstKey(key string) bool {
	// Make the keys lowercase so we can match: isIphone(), isiPhone(), isiphone(), etc.
	key = strings.ToLower(key)

	//change the keys to lower case
	rules := make(map[string]string)
	for ruleKey, ruleValue := range md.rules.getMobileDetectionRules() {
		ruleKey = strings.ToLower(ruleKey)
		rules[ruleKey] = ruleValue
	}

	if rule, ok := rules[key]; ok && "" != rule {
		return md.match(rule)
	}
	return false
}

//Find a detection rule that matches the current User-agent.
func (md *MobileDetect) matchDetectionRulesAgainstUA() bool {
	for _, ruleValue := range md.rules.getMobileDetectionRules() {
		if "" != ruleValue {
			if md.match(ruleValue) {
				return true
			}
		}
	}

	return false
}

// Some detection rules are relative (not standard),because of the diversity of devices, vendors and
// their conventions in representing the User-Agent or the HTTP headers.
// This method will be used to check custom regexes against the User-Agent string.
// @todo: search in the HTTP headers too.
func (md *MobileDetect) match(rule string) bool {
	//Escape the special character which is the delimiter
	//rule = strings.Replace(rule, `\`, `\/`, -1)
	rule = `(?is)` + rule
	if _, ok := md.compiledRegexRules[rule]; !ok {
		md.compiledRegexRules[rule] = regexp.MustCompile(rule)
	}
	re := md.compiledRegexRules[rule]
	ret := re.MatchString(md.userAgent)
	return ret
}

// CheckHttpHeadersForMobile looks for mobile rules to confirm if the browser is a mobile browser
func (md *MobileDetect) CheckHttpHeadersForMobile() bool {
	for _, mobileHeader := range md.getMobileHeaders() {
		if headerString, ok := md.httpHeaders[mobileHeader]; ok {
			mobileHeaderMatches := md.getMobileHeaderMatches()
			if matches, ok := mobileHeaderMatches[mobileHeader]; ok {
				for _, match := range matches {
					if -1 != strings.Index(headerString, match) {
						return true
					}
				}
				return false
			} else {
				return true
			}
		}
	}
	return false
}

func (md *MobileDetect) getMobileHeaders() []string {
	return []string{
		"HTTP_ACCEPT",
		"HTTP_X_WAP_PROFILE",
		"HTTP_X_WAP_CLIENTID",
		"HTTP_WAP_CONNECTION",
		"HTTP_PROFILE",
		// Reported by Opera on Nokia devices (eg. C3).
		"HTTP_X_OPERAMINI_PHONE_UA",
		"HTTP_X_NOKIA_GATEWAY_ID",
		"HTTP_X_ORANGE_ID",
		"HTTP_X_VODAFONE_3GPDPCONTEXT",
		"HTTP_X_HUAWEI_USERID",
		// Reported by Windows Smartphones.
		"HTTP_UA_OS",
		// Reported by Verizon, Vodafone proxy system.
		"HTTP_X_MOBILE_GATEWAY",
		// Seend this on HTC Sensation. @ref: SensationXE_Beats_Z715e.
		"HTTP_X_ATT_DEVICEID",
		// Seen this on a HTC.
		"HTTP_UA_CPU",
	}
}

func (md *MobileDetect) getMobileHeaderMatches() map[string][]string {
	return map[string][]string{
		"HTTP_ACCEPT": []string{
			// Opera Mini; @reference: http://dev.opera.com/articles/view/opera-binary-markup-language/
			"application/x-obml2d",
			// BlackBerry devices.
			"application/vnd.rim.html",
			"text/vnd.wap.wml",
			"application/vnd.wap.xhtml+xml",
		},
		"HTTP_UA_CPU": []string{"ARM"},
	}
}

// Properties helps parsing User Agent string, extracting useful segments of text.
//VER refers to the regular expression defined in the constant self::VER.
func (md *MobileDetect) Properties() map[string][]string {
	return map[string][]string{

		// Build
		`Mobile`:   []string{`Mobile/[VER]`},
		`Build`:    []string{`Build/[VER]`},
		`Version`:  []string{`Version/[VER]`},
		`VendorID`: []string{`VendorID/[VER]`},

		// Devices
		`iPad`:   []string{`iPad.*CPU[a-z ]+[VER]`},
		`iPhone`: []string{`iPhone.*CPU[a-z ]+[VER]`},
		`iPod`:   []string{`iPod.*CPU[a-z ]+[VER]`},
		//`BlackBerry`    : array(`BlackBerry[VER]`, `BlackBerry [VER];`),
		`Kindle`: []string{`Kindle/[VER]`},

		// Browser
		`Chrome`: []string{`Chrome/[VER]`, `CriOS/[VER]`, `CrMo/[VER]`},
		`Coast`:  []string{`Coast/[VER]`},
		`Dolfin`: []string{`Dolfin/[VER]`},
		// @reference: https://developer.mozilla.org/en-US/docs/User_Agent_Strings_Reference
		`Firefox`: []string{`Firefox/[VER]`},
		`Fennec`:  []string{`Fennec/[VER]`},
		// @reference: http://msdn.microsoft.com/en-us/library/ms537503(v=vs.85).aspx
		`IE`: []string{`IEMobile/[VER];`, `IEMobile [VER]`, `MSIE [VER];`},
		// http://en.wikipedia.org/wiki/NetFront
		`NetFront`:       []string{`NetFront/[VER]`},
		`NokiaBrowser`:   []string{`NokiaBrowser/[VER]`},
		`Opera`:          []string{` OPR/[VER]`, `Opera Mini/[VER]`, `Version/[VER]`},
		`Opera Mini`:     []string{`Opera Mini/[VER]`},
		`Opera Mobi`:     []string{`Version/[VER]`},
		`UC Browser`:     []string{`UC Browser[VER]`},
		`MQQBrowser`:     []string{`MQQBrowser/[VER]`},
		`MicroMessenger`: []string{`MicroMessenger/[VER]`},
		// @note: Safari 7534.48.3 is actually Version 5.1.
		// @note: On BlackBerry the Version is overwriten by the OS.
		`Safari`:  []string{`Version/[VER]`, `Safari/[VER]`},
		`Skyfire`: []string{`Skyfire/[VER]`},
		`Tizen`:   []string{`Tizen/[VER]`},
		`Webkit`:  []string{`webkit[ /][VER]`},

		// Engine
		`Gecko`:   []string{`Gecko/[VER]`},
		`Trident`: []string{`Trident/[VER]`},
		`Presto`:  []string{`Presto/[VER]`},

		// OS
		`iOS`:        []string{` \bOS\b [VER] `},
		`Android`:    []string{`Android [VER]`},
		`BlackBerry`: []string{`BlackBerry[\w]+/[VER]`, `BlackBerry.*Version/[VER]`, `Version/[VER]`},
		`BREW`:       []string{`BREW [VER]`},
		`Java`:       []string{`Java/[VER]`},
		// @reference: http://windowsteamblog.com/windows_phone/b/wpdev/archive/2011/08/29/introducing-the-ie9-on-windows-phone-mango-user-agent-string.aspx
		// @reference: http://en.wikipedia.org/wiki/Windows_NT#Releases
		`Windows Phone OS`: []string{`Windows Phone OS [VER]`, `Windows Phone [VER]`},
		`Windows Phone`:    []string{`Windows Phone [VER]`},
		`Windows CE`:       []string{`Windows CE/[VER]`},
		// http://social.msdn.microsoft.com/Forums/en-US/windowsdeveloperpreviewgeneral/thread/6be392da-4d2f-41b4-8354-8dcee20c85cd
		`Windows NT`: []string{`Windows NT [VER]`},
		`Symbian`:    []string{`SymbianOS/[VER]`, `Symbian/[VER]`},
		`webOS`:      []string{`webOS/[VER]`, `hpwOS/[VER];`},
	}
}

// MobileGrade returns a graduation similar to jQuery's Graded Browse Support
func (md *MobileDetect) MobileGrade() string {
	isMobile := md.IsMobile()

	if md.isMobileGradeA(isMobile) {
		return MOBILE_GRADE_A
	}
	if md.isMobileGradeB() {
		return MOBILE_GRADE_B
	}
	return MOBILE_GRADE_C
}

func (md *MobileDetect) isMobileGradeA(isMobile bool) bool {
	if md.VersionFloat("iPad") >= 4.3 || md.VersionFloat("iPhone") >= 3.1 || md.VersionFloat("iPod") >= 3.1 ||
		(md.VersionFloat("Android") > 2.1 && md.Is("Webkit")) ||
		md.VersionFloat("Windows Phone OS") >= 7.0 ||
		md.Is("BlackBerry") && md.VersionFloat("BlackBerry") >= 6.0 ||
		md.match("Playbook.*Tablet") ||
		(md.VersionFloat("webOS") >= 1.4 && md.match("Palm|Pre|Pixi")) ||
		md.match("hp.*TouchPad") ||
		(md.Is("Firefox") && md.VersionFloat("Firefox") >= 12) ||
		(md.Is("Chrome") && md.Is("AndroidOS") && md.VersionFloat("Android") >= 4.0) ||
		(md.Is("Skyfire") && md.VersionFloat("Skyfire") >= 4.1 && md.Is("AndroidOS") && md.VersionFloat("Android") >= 2.3) ||
		(md.Is("Opera") && md.VersionFloat("Opera Mobi") > 11 && md.Is("AndroidOS")) ||
		md.Is("MeeGoOS") ||
		md.Is("Tizen") ||
		md.Is("Dolfin") && md.VersionFloat("Bada") >= 2.0 ||
		((md.Is("UC Browser") || md.Is("Dolfin")) && md.VersionFloat("Android") >= 2.3) ||
		(md.match("Kindle Fire") || md.Is("Kindle") && md.VersionFloat("Kindle") >= 3.0) ||
		(md.Is("AndroidOS") && md.Is("NookTablet")) ||
		(md.VersionFloat("Chrome") >= 11 && isMobile) ||
		(md.VersionFloat("Safari") >= 5.0 && isMobile) ||
		(md.VersionFloat("Firefox") >= 4.0 && isMobile) ||
		(md.VersionFloat("MSIE") >= 7.0 && isMobile) ||
		(md.VersionFloat("Opera") >= 10 && isMobile) {
		return true
	}
	return false
}
func (md *MobileDetect) isMobileGradeB() bool {
	if (md.Is("Blackberry") && md.VersionFloat("BlackBerry") >= 5 && md.VersionFloat("BlackBerry") < 6) ||
		(md.VersionFloat("Opera Mini") >= 5.0 && md.VersionFloat("Opera Mini") <= 6.5 && (md.VersionFloat("Android") >= 2.3 || md.Is("iOS"))) ||
		md.match("NokiaN8|NokiaC7|N97.*Series60|Symbian/3") ||
		(md.VersionFloat("Opera Mobi") >= 11 && md.Is("SymbianOS")) {
		return true
	}
	return false
}

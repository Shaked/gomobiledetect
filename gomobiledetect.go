package gomobiledetect

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	//A frequently used regular expression to extract version #s.
	VER = `([\w._\+]+)`
)

type MobileDetect struct {
	rules                *rules
	userAgent            string
	httpHeaders          map[string]string
	mobileDetectionRules map[string]string
}

func NewMobileDetect(rules *rules) *MobileDetect {
	md := &MobileDetect{rules: rules}
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

func (md *MobileDetect) IsMobile() bool {
	if md.CheckHttpHeadersForMobile() {
		return true
	}
	return md.matchDetectionRulesAgainstUA()
}

func (md *MobileDetect) IsTablet() bool {
	for rule := range md.rules.tabletDevices {
		if md.match(rule) {
			return true
		}
	}
	return false
}

func (md *MobileDetect) Is(key string) bool {
	return md.matchUAAgainstKey(key)
}

func (md *MobileDetect) Version(propertyName string) string {
	if "" != propertyName {
		properties := md.Properties()

		if _, ok := properties[propertyName]; ok {
			for _, propertyMatchString := range properties[propertyName] {
				propertyPattern := `(?is)` + strings.Replace(string(propertyMatchString), `[VER]`, VER, -1)

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
	for rule := range md.rules.getMobileDetectionRules() {
		if "" != rule {
			if md.match(rule) {
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
	re := regexp.MustCompile(`(?is)` + rule)
	return re.MatchString(md.userAgent)
}

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
		"HTTP_X_NOKIA_IPADDRESS",
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

//The individual segments that could exist in a User-Agent string.
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

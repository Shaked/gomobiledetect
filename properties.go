package mobiledetect

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	PROP_MOBILE = iota
	PROP_BUILD
	PROP_VERSION
	PROP_VENDORID
	PROP_IPAD
	PROP_IPHONE
	PROP_IPOD
	PROP_KINDLE
	PROP_CHROME
	PROP_COAST
	PROP_DOLFIN
	PROP_FIREFOX
	PROP_FENNEC
	PROP_IE
	PROP_NETFRONT
	PROP_NOKIABROWSER
	PROP_OPERA
	PROP_OPERA_MINI
	PROP_OPERA_MOBI
	PROP_UC_BROWSER
	PROP_MQQBROWSER
	PROP_MICROMESSENGER
	PROP_BAIDUBOXAPP
	PROP_BAIDUBROWSER
	PROP_SAFARI
	PROP_SKYFIRE
	PROP_TIZEN
	PROP_WEBKIT
	PROP_GECKO
	PROP_TRIDENT
	PROP_PRESTO
	PROP_IOS
	PROP_ANDROID
	PROP_BLACKBERRY
	PROP_BREW
	PROP_JAVA
	PROP_WINDOWS_PHONE_OS
	PROP_WINDOWS_PHONE
	PROP_WINDOWS_CE
	PROP_WINDOWS_NT
	PROP_SYMBIAN
	PROP_WEBOS
)

var (
	propertiesNameToVal = map[string]int{
		"mobile":           PROP_MOBILE,
		"build":            PROP_BUILD,
		"version":          PROP_VERSION,
		"vendorid":         PROP_VENDORID,
		"ipad":             PROP_IPAD,
		"iphone":           PROP_IPHONE,
		"ipod":             PROP_IPOD,
		"kindle":           PROP_KINDLE,
		"chrome":           PROP_CHROME,
		"coast":            PROP_COAST,
		"dolfin":           PROP_DOLFIN,
		"firefox":          PROP_FIREFOX,
		"fennec":           PROP_FENNEC,
		"ie":               PROP_IE,
		"netfront":         PROP_NETFRONT,
		"nokiabrowser":     PROP_NOKIABROWSER,
		"opera":            PROP_OPERA,
		"opera mini":       PROP_OPERA_MINI,
		"opera mobi":       PROP_OPERA_MOBI,
		"uc browser":       PROP_UC_BROWSER,
		"mqqbrowser":       PROP_MQQBROWSER,
		"micromessenger":   PROP_MICROMESSENGER,
		"baiduboxapp":      PROP_BAIDUBOXAPP,
		"baidubrowser":     PROP_BAIDUBROWSER,
		"safari":           PROP_SAFARI,
		"skyfire":          PROP_SKYFIRE,
		"tizen":            PROP_TIZEN,
		"webkit":           PROP_WEBKIT,
		"gecko":            PROP_GECKO,
		"trident":          PROP_TRIDENT,
		"presto":           PROP_PRESTO,
		"ios":              PROP_IOS,
		"android":          PROP_ANDROID,
		"blackberry":       PROP_BLACKBERRY,
		"brew":             PROP_BREW,
		"java":             PROP_JAVA,
		"windows phone os": PROP_WINDOWS_PHONE_OS,
		"windows phone":    PROP_WINDOWS_PHONE,
		"windows ce":       PROP_WINDOWS_CE,
		"windows nt":       PROP_WINDOWS_NT,
		"symbian":          PROP_SYMBIAN,
		"webos":            PROP_WEBOS,
	}

	// Properties helps parsing User Agent string, extracting useful segments of text.
	//VER refers to the regular expression defined in the constant self::VER.
	props = [...][]string{
		// Build
		//MOBILE:PROP_
		[]string{`Mobile/[VER]`},
		//PROP_BUILD:
		[]string{`Build/[VER]`},
		//PROP_VERSION:
		[]string{`Version/[VER]`},
		//PROP_VENDORID:
		[]string{`VendorID/[VER]`},
		// Devices
		//PROP_IPAD:
		[]string{`iPad.*CPU[a-z ]+[VER]`},
		//PROP_IPHONE:
		[]string{`iPhone.*CPU[a-z ]+[VER]`},
		//PROP_IPOD:
		[]string{`iPod.*CPU[a-z ]+[VER]`},
		//`BlackBerry`    : array(`BlackBerry[VER]`, `BlackBerry [VER];`),
		//PROP_KINDLE:
		[]string{`Kindle/[VER]`},
		// Browser
		//PROP_CHROME:
		[]string{`Chrome/[VER]`, `CriOS/[VER]`, `CrMo/[VER]`},
		//PROP_COAST:
		[]string{`Coast/[VER]`},
		//PROP_DOLFIN:
		[]string{`Dolfin/[VER]`},
		// @reference: https://developer.mozilla.org/en-US/docs/User_Agent_Strings_Reference
		//PROP_FIREFOX:
		[]string{`Firefox/[VER]`},
		//PROP_FENNEC:
		[]string{`Fennec/[VER]`},
		// @reference: http://msdn.microsoft.com/en-us/library/ms537503(v=vs.85).aspx
		//PROP_IE:
		[]string{`IEMobile/[VER];`, `IEMobile [VER]`, `MSIE [VER];`},
		// http://en.wikipedia.org/wiki/NetFront
		//PROP_NETFRONT:
		[]string{`NetFront/[VER]`},
		//PROP_NOKIABROWSER:
		[]string{`NokiaBrowser/[VER]`},
		//PROP_OPERA:
		[]string{` OPR/[VER]`, `Opera Mini/[VER]`, `Version/[VER]`},
		//PROP_OPERA_MINI:
		[]string{`Opera Mini/[VER]`},
		//PROP_OPERA_MOBI:
		[]string{`Version/[VER]`},
		//PROP_UC_BROWSER:
		[]string{`UC Browser[VER]`},
		//PROP_MQQBROWSER:
		[]string{`MQQBrowser/[VER]`},
		//PROP_MICROMESSENGER:
		[]string{`MicroMessenger/[VER]`},
		//PROP_BAIDUBOXAPP
		[]string{`baiduboxapp/[VER]`},
		//PROP_BAIDUBROWSER
		[]string{`baidubrowser/[VER]`},
		// @note: Safari 7534.48.3 is actually Version 5.1.
		// @note: On BlackBerry the Version is overwriten by the OS.
		//PROP_SAFARI:
		[]string{`Version/[VER]`, `Safari/[VER]`},
		//PROP_SKYFIRE:
		[]string{`Skyfire/[VER]`},
		//PROP_TIZEN:
		[]string{`Tizen/[VER]`},
		//PROP_WEBKIT:
		[]string{`webkit[ /][VER]`},
		// Engine
		//PROP_GECKO:
		[]string{`Gecko/[VER]`},
		//PROP_TRIDENT:
		[]string{`Trident/[VER]`},
		//PROP_PRESTO:
		[]string{`Presto/[VER]`},
		// OS
		//PROP_IOS:
		[]string{` \bOS\b [VER] `},
		//PROP_ANDROID:
		[]string{`Android [VER]`},
		//PROP_BLACKBERRY:
		[]string{`BlackBerry[\w]+/[VER]`, `BlackBerry.*Version/[VER]`, `Version/[VER]`},
		//PROP_BREW:
		[]string{`BREW [VER]`},
		//PROP_JAVA:
		[]string{`Java/[VER]`},
		// @reference: http://windowsteamblog.com/windows_phone/b/wpdev/archive/2011/08/29/introducing-the-ie9-on-windows-phone-mango-user-agent-string.aspx
		// @reference: http://en.wikipedia.org/wiki/Windows_NT#Releases
		//PROP_WINDOWS_PHONE_OS:
		[]string{`Windows Phone OS [VER]`, `Windows Phone [VER]`},
		//PROP_WINDOWS_PHONE:
		[]string{`Windows Phone [VER]`},
		//PROP_WINDOWS_CE:
		[]string{`Windows CE/[VER]`},
		// http://social.msdn.microsoft.com/Forums/en-US/windowsdeveloperpreviewgeneral/thread/6be392da-4d2f-41b4-8354-8dcee20c85cd
		//PROP_WINDOWS_NT:
		[]string{`Windows NT [VER]`},
		//PROP_SYMBIAN:
		[]string{`SymbianOS/[VER]`, `Symbian/[VER]`},
		//PROP_WEBOS:
		[]string{`webOS/[VER]`, `hpwOS/[VER];`},
	}
)

type properties struct {
	cache map[string]*regexp.Regexp
}

func newProperties() *properties {
	p := &properties{}
	p.cache = make(map[string]*regexp.Regexp)
	p.preCompile()
	return p
}

func (p *properties) preCompile() {
	for _, property := range props {
		for _, pattern := range property {
			p.compiledRegexByPattern(pattern)
		}
	}
}

func (p *properties) compiledRegexByPattern(propertyPattern string) *regexp.Regexp {
	re, ok := p.cache[propertyPattern]
	if !ok {
		p.cache[propertyPattern] = regexp.MustCompile(propertyPattern)
	}
	re = p.cache[propertyPattern]
	return re
}

func (p *properties) version(propertyVal int, userAgent string) string {
	if len(props) >= propertyVal {
		for _, propertyMatchString := range props[propertyVal] {
			propertyPattern := `(?is)` + strings.Replace(string(propertyMatchString), `[VER]`, verRegex, -1)

			// Escape the special character which is the delimiter.
			//propertyPattern = strings.Replace(propertyPattern, `/`, `\/`, -1)

			// Identify and extract the version.
			re := p.compiledRegexByPattern(propertyPattern)
			match := re.FindStringSubmatch(userAgent)
			if len(match) > 0 {
				return match[1]
			}
		}
	}
	return ""
}

func (p *properties) nameToKey(propertyName string) int {
	propertyName = strings.ToLower(propertyName)
	propertyVal, ok := propertiesNameToVal[propertyName]
	if !ok {
		return -1
	}
	return propertyVal
}

func (p *properties) versionByName(propertyName, userAgent string) string {
	if "" != propertyName {
		propertyVal := p.nameToKey(propertyName)
		if -1 != propertyVal {
			return p.version(propertyVal, userAgent)
		}
	}
	return ""
}

func (p *properties) versionFloatName(propertyName, userAgent string) float64 {
	propertyVal := p.nameToKey(propertyName)
	if -1 != propertyVal {
		return p.versionFloat(propertyVal, userAgent)
	}
	return 0.0
}

func (p *properties) versionFloat(propertyVal int, userAgent string) float64 {
	version := p.version(propertyVal, userAgent)
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

package mobiledetect

import (
	"fmt"
	"log"
	"runtime"
	"testing"
)

const TEST_SKIPPED = "SKIPPED"

type expectedResult struct {
	isMobile bool
	isTablet bool
	version  map[string]string
	model    string
}

var uaListTests = []struct {
	userAgent string
	er        expectedResult
}{
	//test empty version name
	{
		`Mozilla/5.0 (Linux; U; Android 3.2.1; en-us; A100 Build/HTK55D) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			map[string]string{``: ``},
			"",
		},
	},
	//Acer
	{
		`Mozilla/5.0 (Linux; U; Android 3.2.1; en-us; A100 Build/HTK55D) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1; en-us; A110 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; de-de; A200 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.3`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
				`Build`:   `IML74K`,
			},
			"A200",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; A500 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; A501 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; B1-A71 Build/JZO54K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; B1-710 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.72 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; A1-810 Build/JDQ39) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.58 Safari/537.31`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; nl-nl; A1-810 Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Acer; Allegro)`,
		expectedResult{
			true,
			false,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; A3-A10 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Safari/537.36`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.2.2`,
				`Build`:   `JDQ39`,
				`Webkit`:  `537.36`,
				`Chrome`:  `32.0.1700.99`,
			},
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; A1-811 Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; A1-830 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	//AdvanDigital
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-us; E1C Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; id-id; T3C Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	//Ainol
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.4; en-us; Ainol Novo8 Advanced Build/GRJ22) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; Novo10 Hero Build/20121115) AppleWebKit/535.19 (KHTML like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; es-es; novo9-Spark Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	//AllFine
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-gb; FINE7 GENIUS Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	//ASUS
	{
		`Mozilla/5.0 (Linux; U; Android 3.2.1; en-us; Transformer TF101 Build/HTK75) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `3.2.1`,
				`Webkit`:  `534.13`,
				`Safari`:  `4.0`,
			},
			"Transformer TF101",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; Transformer Build/JRO03L) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; ASUS Transformer Pad TF300T Build/JRO03C) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; fr-fr; Transformer Build/JZO54K; CyanogenMod-10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; asus_laptop Build/IMM76L) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			"",
		},
	},
	{
		`@see https://github.com/Shaked/gomobiledetect/issues/1 Mozilla/5.0 (Linux; U; Android 4.1.1; en-gb; PadFone 2 Build/JRO03L) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-gb; PadFone 2 Build/JRO03L) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.1`,
				`Build`:   `JRO03L`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.1; ME301T Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.94 Safari/537.36`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.2.1`,
				`Build`:   `JOP40D`,
			},
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.1; ME173X Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.94 Safari/537.36`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.2.1`,
				`Build`:   `JOP40D`,
			},
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; TF300T Build/JDQ39E) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.2.2`,
				`Build`:   `JDQ39E`,
			},
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; K00C Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; K00E Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; K00F Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; K00L Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.131 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			"",
		},
	},
	//Alcatel
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.7; en-in; MB525 Build/GWK74; CyanogenMod-7.2.0) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			"",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; it-it; ALCATEL ONE TOUCH 918D Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Android`: `2.3.5`,
				`Webkit`:  `533.1`,
				`Safari`:  `4.0`,
				`Build`:   `GRJ90`,
			},
			"ONE TOUCH 918D",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; ru-ru; ALCATEL ONE TOUCH 991 Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Android`: `2.3.6`,
				`Webkit`:  `533.1`,
				`Safari`:  `4.0`,
				`Build`:   `GRJ90`,
			},
			"ONE TOUCH 991",
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; ru-ru; ALCATEL ONE TOUCH 993D Build/ICECREAM) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Android`: `4.0.4`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
				`Build`:   `ICECREAM`,
			},
			"ONE TOUCH 993D",
		},
	},
	{
		`ALCATEL_A392G/1.0 ObigoInternetBrowser/Q05A[TF013513002719521000000013182904148]`,
		expectedResult{
			true,
			false,
			nil,
			`A392G`,
		},
	},
	{
		`ALCATEL_3020D/1.0 ObigoInternetBrowser/Q03C`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2; ALCATEL ONE TOUCH 5037A Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`ALCATEL_3020G/1.0 ObigoInternetBrowser/Q03C`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`ALCATEL_3041D/1.0 ObigoInternetBrowser/Q03C`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-us; ALCATEL ONE TOUCH 5037E Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 5037X Build/JDQ39) AppleWebKit/537.36 (KHTML like Gecko) Chrome/31.0.1650.59 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-us; ALCATEL ONE TOUCH 5037X Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ONE TOUCH 6012A Build/JDQ39) AppleWebKit/537.36 (KHTML like Gecko) Chrome/28.0.1500.94 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-nz; ONE TOUCH 6012A Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ONE TOUCH 6012D Build/JDQ39) AppleWebKit/537.36 (KHTML like Gecko) Chrome/35.0.1916.141 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; ONE TOUCH 6012D Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; En-us; ONE TOUCH 6012E Build/JDQ39) AppleWebKit/534.30 (KHTML, Like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ONE TOUCH 6012X Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.12975 YaBrowser/13.12.1599.12975 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; bg-bg; ONE TOUCH 6012X Build/JDQ39) AppleWebKit/534.30 (KHTML like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 6012X_orange Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.59 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; fr-fr; ALCATEL ONE TOUCH 6012X_orange Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.3; fr-fr; 6016E Build/JLS36C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.3; en-us; ALCATEL ONE TOUCH 6016E Build/JLS36C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; 6016X Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.3; ru-ru; 6016X Build/JLS36C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 6032A Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.170 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; ru-ru; ALCATEL ONE TOUCH 6032X Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-nz; ALCATEL ONE TOUCH 7040A Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; th-th; ALCATEL ONE TOUCH 7040D Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 7040D Build/JDQ39) AppleWebKit/537.36 (KHTML like Gecko) Chrome/28.0.1500.94 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 7040E Build/JDQ39) AppleWebKit/537.36 (KHTML like Gecko) Chrome/28.0.1500.94 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; it-it; ALCATEL ONE TOUCH 7041D Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 7041D Build/JDQ39) AppleWebKit/537.36 (KHTML like Gecko) Chrome/31.0.1650.59 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 7041X Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.138 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; pt-pt; ALCATEL ONE TOUCH 7041X Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 8020A Build/JDQ39) AppleWebKit/537.36 (KHTML, Like Gecko) Chrome/30.0.1599.92 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-us; ALCATEL ONE TOUCH 8020A Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 8020D Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; es-es; ALCATEL ONE TOUCH 8020D Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 8020E Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-us; ALCATEL ONE TOUCH 8020E Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ALCATEL ONE TOUCH 8020X Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.136 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-us; ALCATEL ONE TOUCH 8020X Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	//Allview
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; ALLVIEW P5 Build/IML74K) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us ; ALLVIEW SPEEDI Build/IMM76D) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1/UCBrowser/8.5.3.246/145/355`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; AllviewCity Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; ALLVIEWSPEED Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Amoi
	{
		`Amoi 8512/R18.0 NF-Browser/3.3`,
		expectedResult{
			true,
			false,
			nil,
			`8512`,
		},
	},
	//Amazon
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; KFTT Build/IML74K) AppleWebKit/535.19 (KHTML, like Gecko) Silk/3.4 Mobile Safari/535.19 Silk-Accelerated=true`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; en-US) AppleWebKit/528.5+ (KHTML, like Gecko, Safari/528.5+) Version/4.0 Kindle/3.0 (screen 600x800; rotate)`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Kindle`: `3.0`,
				`Webkit`: `528.5+`,
				`Safari`: `4.0`,
			},
			`Kindle`,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; KFOTE Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.3`,
				`Build`:   `IML74K`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; WFJWAE Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Apple
	{
		`iTunes/9.1.1`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`iTunes/11.0.2 (Windows; Microsoft Windows 8 x64 Business Edition (Build 9200)) AppleWebKit/536.27.1`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (iPod touch; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A4449d Safari/9537.53`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (iPhone; U; CPU like Mac OS X; en) AppleWebKit/420+ (KHTML, like Gecko) Version/3.0 Mobile/1A543 Safari/419.3`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Webkit`: `420+`,
				`Safari`: `3.0`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (iPhone; U; CPU iPhone OS 3_0 like Mac OS X; en-us) AppleWebKit/528.18 (KHTML, like Gecko) Version/4.0 Mobile/7A341 Safari/528.16`,
		expectedResult{
			true,
			false,
			map[string]string{
				`iOS`:    `3_0`,
				`Webkit`: `528.18`,
				`Safari`: `4.0`,
			},
			`iPhone`,
		},
	},
	{
		`Mozilla/5.0 (iPhone; CPU iPhone OS 5_1_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9B206 Safari/7534.48.3`,
		expectedResult{
			true,
			false,
			map[string]string{
				`iOS`:    `5_1_1`,
				`Webkit`: `534.46`,
				`Mobile`: `9B206`,
				`Safari`: `5.1`,
			},
			`iPhone`,
		},
	},
	{
		`Mozilla/5.0 (iPod; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A403 Safari/8536.25`,
		expectedResult{
			true,
			false,
			map[string]string{
				`iOS`:    `6_0`,
				`Webkit`: `536.26`,
				`Mobile`: `10A403`,
				`Safari`: `6.0`,
			},
			`iPod`,
		},
	},
	{
		`Mozilla/5.0 (iPad; CPU OS 5_1_1 like Mac OS X; en-us) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/21.0.1180.80 Mobile/9B206 Safari/7534.48.3 (6FF046A0-1BC4-4E7D-8A9D-6BF17622A123)`,
		expectedResult{
			true,
			true,
			map[string]string{
				`iOS`:    `5_1_1`,
				`Webkit`: `534.46.0`,
				`Mobile`: `9B206`,
				`Chrome`: `21.0.1180.80`,
			},
			`iPad`,
		},
	},
	{
		`Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A403 Safari/8536.25`,
		expectedResult{
			true,
			true,
			map[string]string{
				`iOS`:    `6_0`,
				`Webkit`: `536.26`,
				`Mobile`: `10A403`,
				`Safari`: `6.0`,
			},
			`iPad`,
		},
	},
	{
		`Mozilla/5.0 (iPad; U; CPU OS 4_2_1 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8C148 Safari/6533.18.5`,
		expectedResult{
			true,
			true,
			map[string]string{
				`iOS`:    `4_2_1`,
				`Webkit`: `533.17.9`,
				`Mobile`: `8C148`,
				`Safari`: `5.0.2`,
			},
			`iPad`,
		},
	},
	{
		`Mozilla/5.0 (iPad; U; CPU OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B334b Safari/531.21.10`,
		expectedResult{
			true,
			true,
			map[string]string{
				`iOS`:    `3_2`,
				`Webkit`: `531.21.10`,
				`Mobile`: `7B334b`,
				`Safari`: `4.0.4`,
			},
			`iPad`,
		},
	},
	{
		`Mozilla/5.0 (iPhone; CPU iPhone OS 6_0_1 like Mac OS X; da-dk) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/21.0.1180.82 Mobile/10A523 Safari/7534.48.3`,
		expectedResult{
			true,
			false,
			map[string]string{
				`iOS`:    `6_0_1`,
				`Webkit`: `534.46.0`,
				`Chrome`: `21.0.1180.82`,
				`Mobile`: `10A523`,
			},
			`iPhone`,
		},
	},
	{
		`Mozilla/5.0 (iPhone; CPU iPhone OS 6_0_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A523 Safari/8536.25`,
		expectedResult{
			true,
			false,
			map[string]string{
				`iOS`:    `6_0_1`,
				`Webkit`: `536.26`,
				`Safari`: `6.0`,
				`Mobile`: `10A523`,
			},
			`iPhone`,
		},
	},
	{
		`Mozilla/5.0 (iPhone; CPU iPhone OS 6_1 like Mac OS X; ru-ru) AppleWebKit/536.26 (KHTML, like Gecko) CriOS/23.0.1271.100 Mobile/10B142 Safari/8536.25`,
		expectedResult{
			true,
			false,
			map[string]string{
				`iOS`:    `6_1`,
				`Webkit`: `536.26`,
				`Chrome`: `23.0.1271.100`,
				`Mobile`: `10B142`,
			},
			`iPhone`,
		},
	},
	{
		`Mozilla/5.0 (iPhone; CPU iPhone OS 6_1 like Mac OS X; ru-ru) AppleWebKit/536.26 (KHTML, like Gecko) CriOS/23.0.1271.100 Mobile/10B142 Safari/8536.25`,
		expectedResult{
			true,
			false,
			map[string]string{
				`iOS`:    `6_1`,
				`Webkit`: `536.26`,
				`Chrome`: `23.0.1271.100`,
				`Mobile`: `10B142`,
			},
			`iPhone`,
		},
	},
	{
		`Mozilla/5.0 (iPhone; CPU iPhone OS 6_1_3 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10B329 Safari/8536.25`,
		expectedResult{
			true,
			false,
			map[string]string{
				`iOS`:    `6_1_3`,
				`Webkit`: `536.26`,
				`Safari`: `6.0`,
				`Mobile`: `10B329`,
			},
			`iPhone`,
		},
	},
	{
		`Mozilla/5.0 (iPad; CPU OS 6_1_3 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Coast/1.0.2.62956 Mobile/10B329 Safari/7534.48.3`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Coast`: `1.0.2.62956`,
			},
			``,
		},
	},
	{
		`CheckMyBus iOS mobile App 0.9.0 (iPhone; iPhone OS/7.1.1)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	//Archos
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Qilive 97R Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.92 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; Archos 50 Platinum Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.94 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; ARCHOS 80G9 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2.1; fr-fr; A101IT Build/FROYO) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Archos 101 Neon Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Archos 101 Cobalt Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; de-de; ARCHOS 80 TITANIUM Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; ARCHOS 101 Titanium Build/JRO03H) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; de-de; ARCHOS 70b TITANIUM Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; de-de; Archos 80 Xenon Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Archos 79 Xenon Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; de-de; ARCHOS 101 Titanium Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; de-de; ARCHOS 80XSK Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; de-de; ARCHOS FAMILYPAD 2 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; ARCHOS 97B TITANIUM Build/JRO03H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ARCHOS 101 XS 2 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; ARCHOS 80b PLATINUM Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Archos 70 Xenon Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; de-de; ARCHOS 97 CARBON Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; de-de; ARCHOS 97 TITANIUMHD Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Archos 90 Neon Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 1.6; de-de; Archos5 Build/Donut) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; de-de; ARCHOS GAMEPAD Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//AudioSonic
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-au; T-17B Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Blaupunkt
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; de-de; Endeavour 800NG Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//BlackBerry
	{
		`Mozilla/5.0 (BlackBerry; U; BlackBerry 9300; en) AppleWebKit/534.8+ (KHTML, like Gecko) Version/6.0.0.546 Mobile Safari/534.8+`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Webkit`:     `534.8+`,
				`BlackBerry`: `6.0.0.546`,
			},
			`BlackBerry 9300`,
		},
	},
	{
		`Mozilla/5.0 (BlackBerry; U; BlackBerry 9360; en-US) AppleWebKit/534.11+ (KHTML, like Gecko) Version/7.0.0.400 Mobile Safari/534.11+`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; he) AppleWebKit/534.8+ (KHTML, like Gecko) Version/6.0.0.723 Mobile Safari/534.8+`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (BlackBerry; U; BlackBerry 9700; en-US) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.448 Mobile Safari/534.8`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (BlackBerry; U; BlackBerry 9790; en-GB) AppleWebKit/534.11+ (KHTML, like Gecko) Version/7.1.0.714 Mobile Safari/534.11+`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Webkit`:     `534.11+`,
				`BlackBerry`: `7.1.0.714`,
			},
			`BlackBerry 9790`,
		},
	},
	{
		`Opera/9.80 (BlackBerry; Opera Mini/7.0.29990/28.2504; U; en) Presto/2.8.119 Version/11.10`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (BlackBerry; U; BlackBerry 9981; en-GB) AppleWebKit/534.11+ (KHTML, like Gecko) Version/7.1.0.342 Mobile Safari/534.11+`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (BlackBerry; U; BlackBerry 9800; en-GB) AppleWebKit/534.8+ (KHTML, like Gecko) Version/6.0.0.546 Mobile Safari/534.8+`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (BlackBerry; U; BlackBerry 9780; es) AppleWebKit/534.8  (KHTML, like Gecko) Version/6.0.0.480 Mobile Safari/534.8`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (BlackBerry; U; BlackBerry 9810; en-US) AppleWebKit/534.11  (KHTML, like Gecko) Version/7.0.0.583 Mobile Safari/534.11`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (BlackBerry; U; BlackBerry 9860; es) AppleWebKit/534.11+ (KHTML, like Gecko) Version/7.0.0.576 Mobile Safari/534.11+`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (BlackBerry; U; BlackBerry 9900; en-US) AppleWebKit/534.11  (KHTML, like Gecko) Version/7.1.0.523 Mobile Safari/534.11`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`BlackBerry8520/5.0.0.592 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/136`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`BlackBerry8520/5.0.0.1067 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/603`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`BlackBerry8520/5.0.0.1036 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/611`,
		expectedResult{
			true,
			false,
			map[string]string{
				`BlackBerry`: `5.0.0.1036`,
				`VendorID`:   `611`,
			},
			`BlackBerry8520`,
		},
	},
	{
		`Mozilla/5.0 (BlackBerry; U; BlackBerry 9220; en) AppleWebKit/534.11+ (KHTML, like Gecko) Version/7.1.0.337 Mobile Safari/534.11+`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML, like Gecko) Version/7.2.1.0 Safari/536.2+`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (BB10; Touch) AppleWebKit/537.1+ (KHTML, like Gecko) Version/10.0.0.1337 Mobile Safari/537.1+`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (BB10; Touch) AppleWebKit/537.10+ (KHTML, like Gecko) Version/10.0.9.2372 Mobile Safari/537.10+`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (BB10; Touch) /537.10+ (KHTML, like Gecko) Version/10.0.9.2372 Mobile Safari/537.10+`,
		expectedResult{
			true,
			false,
			map[string]string{`BlackBerry`: `10.0.9.2372`},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.2.1; en-us; Transformer TF101 Build/HTK75) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `3.2.1`,
				`Webkit`:  `534.13`,
				`Safari`:  `4.0`,
			},
			`Transformer TF101`,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; de-de; A200 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.3`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			`A200`,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; A500 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.3`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			`A500`,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; A501 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.3`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			`A501`,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; Transformer Build/JRO03L) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.1`,
				`Webkit`:  `535.19`,
				`Chrome`:  `18.0.1025.166`,
			},
			`Transformer`,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; ASUS Transformer Pad TF300T Build/JRO03C) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.1`,
				`Webkit`:  `535.19`,
				`Chrome`:  `18.0.1025.166`,
			},
			`Transformer Pad TF300T`,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; fr-fr; Transformer Build/JZO54K; CyanogenMod-10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.2`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
				`Build`:   `JZO54K`,
			},
			`Transformer`,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; B1-A71 Build/JZO54K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.2`,
				`Webkit`:  `535.19`,
				`Chrome`:  `18.0.1025.166`,
			},
			`B1-A71`,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Acer; Allegro)`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Windows Phone OS`: `7.5`,
				`Trident`:          `5.0`,
				`IE`:               `9.0`,
			},
			`Allegro`,
		},
	},
	//Broncho
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; es-es; Broncho N701 Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Digma
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; ru-ru; iDx10 3G Build/ICS.b02ref.20120331) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//bq
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; es-es; bq Livingstone 2 Build/1.1.7 20121018-10:33) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; es-es; bq Edison Build/1.1.10-1015 20121230-18:00) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.3; Maxwell Lite Build/v1.0.0.ICS.maxwell.20120920) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; zh-tw; bq Maxwell Plus Build/1.0.0 20120913-10:39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Casio
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; en-us; C771 Build/C771M120) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	//ChangJia
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; pt-br; TPC97113 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; TPC7102 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	// Coby @ref: http://www.cobyusa.com/?p=pcat&pcat_id=3001
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; en-us; MID7010 Build/FRF85B) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; MID7048 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; MID8042 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Concorde
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; hu-hu; ConCorde Tab T10 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; hu-hu; ConCorde tab PLAY Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Cresta
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; nl-nl; CRESTA.CTP888 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	// Cube
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; ru-ru; CUBE U9GT 2 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Danew
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; es-es; Dslide 700 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.3`,
				`Build`:   `IML74K`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			`Dslide 700`,
		},
	},
	//DanyTech
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Genius Tab Q4 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.59 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Dell
	{
		`Mozilla/5.0 (Linux; U; Android 1.6; en-gb; Dell Streak Build/Donut AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/ 525.20.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.7; hd-us; Dell Venue Build/GWK74; CyanogenMod-7.2.0) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; DELL; Venue Pro)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	//DPS
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; DPS Dream 9 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//ECS
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; it-it; TM105A Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.4`,
				`Build`:   `IMM76D`,
				`Webkit`:  `534.30`,
			},
			``,
		},
	},
	//Eboda
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; ro-ro; E-Boda Supreme Dual Core X190 Build/JRO03C) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.1`,
				`Safari`:  `4.0`,
				`Webkit`:  `534.30`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; ro-ro; E-Boda Essential A160 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.3; E-Boda Supreme X80 Dual Core Build/ICS.g12refM806A1YBD.20120925) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; ro-ro; E-boda essential smile Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; ro-ro; E-Boda Supreme X80 Dual Core Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; ro-ro; E-Boda Supreme XL200IPS Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Evolio
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-us; Evolio X7 Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; ro-ro; ARIA_Mini_wifi Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Fly
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; ru-ru; Fly IQ440; Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; ru-ru; FLY IQ256 Build/GRK39F) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	//Fujitsu
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; ja-jp; F-10D Build/V21R48A) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.3`,
				`Build`:   `V21R48A`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; ru-ru; M532 Build/IML74K) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.3`,
				`Build`:   `IML74K`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	//FX2
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; FX2 PAD7 RK Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	// Galapad @ref: http://www.galapad.net/product.html
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; zh-tw; G1 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.1`,
				`Build`:   `JRO03C`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	// GoClever
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; el-gr; GOCLEVER TAB A103 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.4; zh-tw; A7GOCLEVER Build/GRJ22) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; GOCLEVER TAB A104 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; cs-cz; GOCLEVER TAB A93.2 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; sk-sk; GOCLEVER TAB A971 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; lv-lv; GOCLEVER TAB A972BK Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; lv-lv; GOCLEVER TAB A972BK Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; fr-fr; GOCLEVER TAB A104.2 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; pt-pt; GOCLEVER TAB T76 Build/MID) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Google
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Nexus 4 Build/JDQ39) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.40 Mobile Safari/537.31 OPR/14.0.1074.54070`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Android`: `4.2.2`,
				`Build`:   `JDQ39`,
				`Webkit`:  `537.31`,
				`Opera`:   `14.0.1074.54070`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Nexus 4 Build/JDQ39) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.58 Mobile Safari/537.31`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Android`: `4.2.2`,
				`Chrome`:  `26.0.1410.58`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; Google Nexus 4 - 4.1.1 - API 16 - 768x1280 Build/JRO03S) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; Google Galaxy Nexus - 4.1.1 - API 16 - 720x1280 Build/JRO03S) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`@see https://github.com/Shaked/gomobiledetect/issues/1 Mozilla/5.0 (Linux; Android 4.1.1; Nexus 7 Build/JRO03D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`@see https://github.com/Shaked/gomobiledetect/issues/1 Mozilla/5.0 (Linux; Android 4.2; Nexus 7 Build/JOP40C) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`@see https://github.com/Shaked/gomobiledetect/issues/1 Mozilla/5.0 (Linux; Android 4.1.2; Nexus 7 Build/JZ054K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.2`,
				`Chrome`:  `18.0.1025.166`,
			},
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; cs-cz; Nexus S Build/JZO54K; CyanogenMod-10.0.0) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`@see https://github.com/Shaked/gomobiledetect/issues/1 Mozilla/5.0 (Linux; Android 4.3; Nexus 10 Build/JWR66Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.72 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`@see https://github.com/Shaked/gomobiledetect/issues/1 Mozilla/5.0 (Linux; U; Android; en_us; Nexus 7 Build/) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 NetFrontLifeBrowser/2.3 Mobile (Dragonfruit)`,
		expectedResult{
			true,
			true,
			nil,
			TEST_SKIPPED,
		},
	},
	//GU
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; vi-vn; TX-A1301 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.4`,
				`Build`:   `IMM76D`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; da-dk; Q702 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.3`,
				`Build`:   `IML74K`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	//HCL
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; U1 Build/HCL ME Tablet U1) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; U1 Build/HCL ME Tablet U1) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; Connect-3G-2.0 Build/HCL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.72 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; pt-br; X1 Build/HCL ME Tablet X1) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//HP
	{
		`Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.5; U; en-GB) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/234.83 Safari/534.6 TouchPad/1.0`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-gb; HP Slate 7 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; HP Slate 7 Build/JRO03H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.94 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; HP 8 Build/1.0.7_WW-FIR-13) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; HP Slate 10 HD Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; HP Slate 8 Pro Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Slate 21 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.138 Safari/537.36 OPR/22.0.1485.78487`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; HP SlateBook 10 x2 PC Build/4.3-17r20-03-23) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.94 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//HTC
	{
		`Mozilla/5.0 (X11; Linux x86_64; Z520m; en-ca) AppleWebKit/534.24 (KHTML, like Gecko) Chrome/11.0.696.34 Safari/534.24`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`HTC_Touch_HD_T8282 Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.11)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 1.5; en-us; ADR6200 Build/CUPCAKE) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.1; xx-xx; Desire_A8181 Build/ERE27) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.1-update1; de-de; HTC Desire 1.19.161.5 Build/ERE27) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.1-update1; en-gb; HTC Desire Build/ERE27) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; fr-fr; HTC Desire Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; en-dk; Desire_A8181 Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; xx-xx; 001HT Build/FRF91) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; xx-xx; HTCA8180/1.0 Android/2.2 release/06.23.2010 Browser/WAP 2.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2.2; de-at; HTC Desire Build/FRG83G) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2.2; en-sk; Desire_A8181 Build/FRG83G) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3; xx-xx; HTC/DesireS/1.07.163.1 Build/GRH78C) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; en-lv; HTC_DesireZ_A7272 Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.4; en-us; ADR6300 Build/GRJ22) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; en-gb; HTC/DesireS/2.10.161.3 Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; ru-ru; HTC_DesireS_S510e Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; en-us; Inspire 4G Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; de-de; HTC Explorer A310e Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; en-gb; HTC_ChaCha_A810e Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; nl-nl; HTC_DesireHD_A9191 Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; ru-ru; HTC Desire S Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; en-au; HTC Desire Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; de-de; HTC_DesireHD Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; ru-ua; HTC_WildfireS_A510e Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.7; en-us; HTC Vision Build/GRI40; ILWT-CM7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0; xx-xx; HTC_GOF_U/1.05.161.1 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; hu-hu; HTC Sensation Z710e Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; pl-pl; EVO3D_X515m Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; ru-ru; HTC_One_S Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; ru-ru; HTC_One_V Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; zh-cn; HTC_A320e Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; zh-tw; HTC Desire V Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.3; PG86100 Build/IML74K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-nl; SensationXE_Beats_Z715e Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; ADR6425LVW 4G Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.3; HTC One V Build/IML74K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; cs-ee; Sensation_Z710e Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; HTC Evo 4G Build/MIUI) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; Desire HD Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-my; HTC_One_X Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; it-it; IncredibleS_S710e Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; fr-fr; HTC_Desire_S Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; HTC One X Build/JRO03C) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; HTC Butterfly Build/JRO03C) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; EVO Build/JRO03C) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.169 Mobile Safari/537.22`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; HTCSensation Build/JRO03C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; GT-S6312 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; TITAN X310e)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Radar C110e)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; T8788)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC 7 Mozart T8698; QSD8x50)`,
		expectedResult{
			true,
			false,
			map[string]string{
				`IE`:               `9.0`,
				`Windows Phone OS`: `7.5`,
				`Trident`:          `5.0`,
			},
			`7 Mozart T8698`,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.1; zh-tw; HTC PG09410 Build/HMJ15) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; 7 HTC MOZART)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; 7 Mondrian T8788)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; 7 Mozart T8698)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; 7 Mozart)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; 7 Mozart; Orange)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; 7 Pro T7576)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; 7 Pro)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; 7 Schubert T9292)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; 7 Surround)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; 7 Trophy T8686)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; 7 Trophy)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Eternity)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Gold)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; HD2 LEO)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; HD2)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; HD7 T9292)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; HD7)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; iPad 3)`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; LEO)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Mazaa)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Mondrian)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Mozart T8698)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Mozart)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; mwp6985)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; PC40100)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; PC40200)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; PD67100)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; PI39100)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; PI86100)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Radar 4G)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Radar C110e; 1.08.164.02)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Radar C110e; 2.05.164.01)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Radar C110e; 2.05.168.02)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Radar)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Radar; Orange)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Schuber)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Schubert T9292)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Schubert)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Spark)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Surround)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; T7575)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; T8697)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; T9295)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; T9296)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Titan)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Torphy T8686)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; X310e)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC_blocked; T8788)`,
		expectedResult{
			true,
			false,
			map[string]string{
				`IE`:               `9.0`,
				`Windows Phone OS`: `7.5`,
				`Trident`:          `5.0`,
			},
			`T8788`,
		},
	},
	//Hudl
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Hudl HT7S3 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.82 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Huwaei
	{
		`Mozilla/5.0 (Linux; U; Android 2.1-update1; bg-bg; Ideos S7 Build/ERE27) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.1; en-us; Ideos S7 Build/ERE27) AppleWebKit/525.10+ (KHTML, like Gecko) Version/3.0.4 Mobile Safari/523.12.2`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; lt-lt; U8660 Build/HuaweiU8660) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.7; ru-ru; HUAWEI-U8850 Build/HuaweiU8850) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.2; pl-pl; MediaPad Build/HuaweiMediaPad) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.2; nl-nl; HUAWEI MediaPad Build/HuaweiMediaPad) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`HUAWEI_T8951_TD/1.0 Android/4.0.4 (Linux; U; Android 4.0.4; zh-cn) Release/05.31.2012 Browser/WAP2.0 (AppleWebKit/534.30) Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; ar-eg; MediaPad 7 Youth Build/HuaweiMediaPad) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; zh-cn; HW-HUAWEI_C8815/C8815V100R001C541B135; 540*960; CTC/2.0) AppleWebKit/534.30 (KHTML, like Gecko) Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; zh-cn; HW-HUAWEI_C8813D/C8813DV100R001C92B172; 480*854; CTC/2.0) AppleWebKit/534.30 (KHTML, like Gecko) Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; zh-cn; HW-HUAWEI_Y300C/Y300CV100R001C92B168; 480*800; CTC/2.0) AppleWebKit/534.30 (KHTML, like Gecko) Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	//Iconbit
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; NT-3702M Build/JRO03C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.72 Safari/537.36 OPR/16.0.1212.65583`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; ru-ru; NetTAB SPACE II Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//iJoy
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; fr-fr; Tablet Planet II-v3 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Intenso
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1.;de-de; INM8002KP Build/JR003H) AppleWebKit/534.30 (KHTML, like Gecko)Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.1.`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; TAB1004 Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//INQ
	{
		`INQ1/R3.9.12 NF-Browser/3.5`,
		expectedResult{
			true,
			false,
			nil,
			`INQ1`,
		},
	},
	//IRU
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; ru-ru; M702pro Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//JXD
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; F3000 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Karbonn
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; ST10 Build/JRO03C) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Kobo
	{
		`Mozilla/5.0 (Linux; U; Android 2.0; en-us;) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1 (Kobo Touch)`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `2.0`,
				`Webkit`:  `533.1`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	//Lenovo
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; es-es; IdeaTab_A1107 Build/MR1) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.3; IdeaTab A2107A-H Build/IML74K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.90 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-au; ThinkPad Tablet Build/ThinkPadTablet_A400_03) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`UCWEB/2.0 (Linux; U; Opera Mini/7.1.32052/30.3697; en-US; IdeaTabA1000-G) U2/1.0.0 UCBrowser/9.2.0.419 Mobile`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; IdeaTabA1000-F Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.1; Lenovo A3000-H Build/JOP40D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.117 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; IdeaTab A3000-F Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Safari/537.360`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1; zh-cn; Lenovo-A3000-H/S100) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.1 Mobile Safari/534.300`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; es-us; IdeaTab A3000-F Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.3; IdeaTab A2107A-H Build/IML74K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; de-de; IdeaTab A2107A-H Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; es-es; IdeaTabA2109A Build/JRO03R) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; nl-nl; IdeaTabA2109A Build/JRO03R) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; es-es; IdeaTab_A1107 Build/MR1) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.300`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; IdeaTab S6000-H Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.138 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-gb; IdeaTab S6000-F Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Lenovo B8000-F Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2;it-it; Lenovo B8000-F/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; it-it; Lenovo B6000-F/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.2.2 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Lenovo B6000-F Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; IdeaPadA10 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.166 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; Ideapad K1 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.7; de-de; IdeaPad A1 Build/GRK393; CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; Lenovo B8080-H Build/JLS36C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; hu-hu; Lenovo A3500-FL Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; Lenovo A7600-F Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; IdeaPadA10 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.94 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; Lenovo A5500-F Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.131 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//LG
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; en-us; LG-P509 Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1 MMS/LG-Android-MMS-V1.0/1.2`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2.2; pt-br; LG-P350f Build/FRG83G) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1 MMS/LG-Android-MMS-V1.0/1.2`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; en-us; LG-P500 Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1 MMS/LG-Android-MMS-V1.0/1.2`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; en-us; LS670 Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.4; ru-ru; LG-E510 Build/GRJ22) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1 MMS/LG-Android-MMS-V1.0/1.2`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.4; en-us; VS910 4G Build/GRJ22) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; nl-nl; LG-P700 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; ko-kr; LG-L160L Build/IML74K) AppleWebkit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; LG-F160S Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; nl-nl; LG-E610v/V10f Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; LG-E612 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; LG-F180K Build/JZO54K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; LG-V500 Build/JDQ39B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.59 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; LG-LW770 Build/IMM76I) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; LG-V510 Build/KOT49H.L004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; LG; LG E-900)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; LG; LG-C900)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; LG; LG-E900)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; LG; LG-E900; Orange)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; LG; LG-E900h)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; LG; LG-Optimus 7)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	// @ref: http://ja.wikipedia.org/wiki/L-06C
	{
		`Mozilla/5.0 (Linux; U; Android 3.0.1; ja-jp; L-06C Build/HRI66) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.0; en-us; LG-V900 Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Megafon
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; ru-ru; MegaFon V9 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; MT7A Build/JRO03C) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.58 Safari/537.31`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//MediaTek
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; fr-fr; MT8377 Build/JRO03C) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Safari/534.30/4.05d.1002.m7`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Micromax
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; Micromax A110 Build/JRO03C) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.169 Mobile Safari/537.22`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Android`: `4.1.1`,
				`Build`:   `JRO03C`,
				`Webkit`:  `537.22`,
				`Chrome`:  `25.0.1364.169`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0; xx-xx; Micromax P250(Funbook) Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0`,
				`Build`:   `IMM76D`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	//Microsoft
	// Surface tablet
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch; .NET4.0E; .NET4.0C; Tablet PC 2.0)`,
		expectedResult{
			true,
			true,
			map[string]string{
				`IE`:         `10.0`,
				`Windows NT`: `6.2`,
				`Trident`:    `6.0`,
			},
			``,
		},
	},
	// Ambiguos.
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	// Ambiguos.
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	// http://www.whatismybrowser.com/developers/unknown-user-agent-fragments
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch; ARMBJS)`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0; Touch; MASMJS)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// Thanks to Jonathan Donzallaz!
	// Firefox (nightly) in metro mode on Dell XPS 12
	{
		`Mozilla/5.0 (Windows NT 6.2; WOW64; rv:25.0) Gecko/20130626 Firefox/25.0`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// Firefox in desktop mode on Dell XPS 12
	{
		`Mozilla/5.0 (Windows NT 6.2; WOW64; rv:22.0) Gecko/20100101 Firefox/22.0`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// IE10 in metro mode on Dell XPS 12
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0; MDDCJS)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// IE10 in desktop mode on Dell XPS 12
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; WOW64; Trident/6.0; MDDCJS)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// Opera on Dell XPS 12
	{
		`Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.52 Safari/537.36 OPR/15.0.1147.130`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// Chrome on Dell XPS 12
	{
		`Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.116 Safari/537.36`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// Google search app from Windows Store
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0; Touch; MDDCJS; WebView/1.0)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// Modecom
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; pl-pl; FreeTAB 1014 IPS X4+ Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	// Motorola
	{
		`MOT-W510/08.11.05R MIB/BER2.2 Profile/MIDP-2.0 Configuration/CLDC-1.1 EGE/1.0 UP.Link/6.3.0.0.0`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2.2; zh-cn; ME722 Build/MLS2GC_2.6.0) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.4; en-us; DROIDX Build/4.5.1_57_DX8-51) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; en-us; MB855 Build/4.5.1A-1_SUN-254_13) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; es-us; MB526 Build/4.5.2-51_DFL-50) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; en-ca; MB860 Build/4.5.2A-51_OLL-17.8) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.7; en-us; MOT-XT535 Build/V1.540) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.7; ko-kr; A853 Build/SHOLS_U2_05.26.3; CyanogenMod-7.1.2) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.0; en-us; Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.1; en-us; Xoom Build/HMJ25) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; DROID RAZR 4G Build/6.7.2-180_DHD-16_M4-31) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; Xoom Build/IMM76L) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; pt-br; XT687 Build/V2.27D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Android`: `4.0.4`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			`XT687`,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; es-es; MOT-XT910 Build/6.7.2-180_SPU-19-TA-11.6) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; XT910 Build/9.8.2O-124_SPUL-17) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; XT915 Build/2_32A_2031) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.94 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; XT919 Build/2_290_2017) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.64 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; XT925 Build/9.8.2Q-50-XT925_VQLM-20) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.90 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; XT907 Build/9.8.1Q-66) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.90 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; XT901 Build/9.8.2Q-50_SLS-13) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; DROID BIONIC Build/9.8.2O-72_VZW-22) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; XT1022 Build/KXC20.82-14) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.4.4; en-in; XT1022 Build/KXC21.5-40) AppleWebKit/537.16 (KHTML, like Gecko) Version/4.0 Mobile Safari/537.16`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; XT1025 Build/KXC20.82-13) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; XT1052 Build/KLA20.16-2.16.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-gb; XT1052 Build/13.9.0Q2.X_83) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; XT1053 Build/13.9.0Q2.X_61) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.58 Mobile Safari/537.31`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-us; XT1053 Build/13.9.0Q2.X_55) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; XT1056 Build/13.9.0Q2.X-116-MX-17-6-2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.64 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; XT1031 Build/KXB20.9-1.10-1.18-1.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.4; XT1032 Build/KXB21.14-L1.40) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.4.2; de-de; XT1032 Build/KLB20.9-1.10-1.24-1.1) AppleWebKit/537.16 (KHTML, like Gecko) Version/4.0 Mobile Safari/537.16`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; XT1034 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.3; en-us; XT1034 Build/14.10.0Q3.X-84-16) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; XT1035 Build/14.10.0Q3.X-23) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.59 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.3; XT1039 Build/KXB21.14-L1.31) AppleWebKit/537.36 (KHTML like Gecko) Chrome/35.0.1916.141 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; XT919 Build/2_290_2002) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; XT919 Build/2_290_2004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.90 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; XT920 Build/2_290_2014) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; XT920 Build/2_310_2014) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; XT905 Build/7.7.1Q_GCIRD-16) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.169 Mobile Safari/537.22`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; XT908 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; XT897 Build/7.7.1Q-6_SPR-ASANTI_LE-18) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	//MSI
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; el-gr; MSI Enjoy 10 Plus Build/1.2) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	// Nabi @ref: https://www.nabitablet.com/
	{
		`Mozilla/5.0 (Linux; U; Android 2.2.1; en-us; NABI-A Build/MASTER) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	// NEC
	// @ref: http://www.n-keitai.com/n-08d/?from=mediasnet
	// @ref: http://devlog.dcm-gate.com/2012/03/medias-tab-n-06duseragnet.html
	// @ref: http://keitaiall.jp/N-08D.html aka MEDIAS TAB
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; ja-jp; N-08D Build/A5001911) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android AAA; BBB; N-06D Build/CCC) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Nook
	{
		`Mozilla/5.0 (Linux; U; Android 2.2.1; en-us; NOOK BNRV200 Build/ERD79 1.4.3) Apple WebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `2.2.1`,
				`Webkit`:  `533.1`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; NOOK BNTV400 Build/ICS) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.4`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
				`Build`:   `ICS`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; BNTV600 Build/IMM76L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.94 Safari/537.36 Hughes-PFB/CID5391275.AID1376709964`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.4`,
				`Webkit`:  `537.36`,
				`Chrome`:  `28.0.1500.94`,
				`Build`:   `IMM76L`,
			},
			``,
		},
	},
	//Nokia
	{
		`Nokia200/2.0 (12.04) Profile/MIDP-2.1 Configuration/CLDC-1.1 UCWEB/2.0 (Java; U; MIDP-2.0; en-US; nokia200) U2/1.0.0 UCBrowser/8.9.0.251 U2/1.0.0 Mobile UNTRUSTED/1.0`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Nokia6303iclassic/5.0 (06.61) Profile/MIDP-2.1 Configuration/CLDC-1.1 Mozilla/5.0 AppleWebKit/420+ (KHTML, like Gecko) Safari/420+`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`nokian73-1/UC Browser7.8.0.95/69/400 UNTRUSTED/1.0`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Nokia2760/2.0 (06.82) Profile/MIDP-2.1 Configuration/CLDC-1.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Nokia3650/1.0 SymbianOS/6.1 Series60/1.2 Profile/MIDP-1.0 Configuration/CLDC-1.0`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`NokiaN70-1/5.0737.3.0.1 Series60/2.8 Profile/MIDP-2.0 Configuration/CLDC-1.1/UC Browser7.8.0.95/27/352`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (S60V3; U; ru; NokiaN73) AppleWebKit/530.13 (KHTML, like Gecko) UCBrowser/8.6.0.199/28/444/UCWEB Mobile`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (S60V3; U; ru; NokiaC5-00.2)/UC Browser8.5.0.183/28/444/UCWEB Mobile`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (S60V3; U; ru; NokiaC5-00.2) AppleWebKit/530.13 (KHTML, like Gecko) UCBrowser/8.7.0.218/28/352/UCWEB Mobile`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Series40; NokiaC3-00/08.63; Profile/MIDP-2.1 Configuration/CLDC-1.1) Gecko/20100401 S40OviBrowser/2.2.0.0.33`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Opera/9.80 (Series 60; Opera Mini/7.0.31380/28.2725; U; es) Presto/2.8.119 Version/11.10`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Opera Mini`: `7.0.31380`,
				`Presto`:     `2.8.119`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Symbian/3; Series60/5.2 NokiaC7-00/025.007; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.37 Mobile Safari/533.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Symbian/3; Series60/5.2 NokiaX7-00/022.014; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.37 Mobile Safari/533.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Symbian/3; Series60/5.3 NokiaE6-00/111.140.0058; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/535.1 (KHTML, like Gecko) NokiaBrowser/8.3.1.4 Mobile Safari/535.1 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Symbian/3; Series60/5.3 NokiaC6-01/111.040.1511; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/535.1 (KHTML, like Gecko) NokiaBrowser/8.3.1.4 Mobile Safari/535.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Symbian/3; Series60/5.3 NokiaC6-01; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.4.2.6 Mobile Safari/533.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Symbian/3; Series60/5.3 NokiaC6-01; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.4.2.6 Mobile Safari/533.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Symbian/3; Series60/5.3 Nokia700/111.020.0308; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.4.1.14 Mobile Safari/533.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Symbian/3; Series60/5.3 NokiaN8-00/111.040.1511; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/535.1 (KHTML, like Gecko) NokiaBrowser/8.3.1.4 Mobile Safari/535.1 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Symbian/3; Series60/5.3 Nokia701/111.030.0609; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.4.2.6 Mobile Safari/533.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 Nokia6120c/3.83; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 Nokia6120ci/7.02; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 Nokia6120c/7.10; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE66-1/510.21.009; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE71-1/110.07.127; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaN95-3/20.2.011 Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE51-1/200.34.36; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaE63-1/500.21.009; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.2; U; Series60/3.1 NokiaN82/10.0.046; Profile/MIDP-2.0 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.3; Series60/3.2 NokiaE52-1/052.003; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/525 (KHTML, like Gecko) Version/3.0 BrowserNG/7.2.6.2`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.3; Series60/3.2 NokiaE52-1/@version@; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.26 Mobile Safari/533.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.3; Series60/3.2 NokiaC5-00/031.022; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/525 (KHTML, like Gecko) Version/3.0 BrowserNG/7.2.3.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.3; Series60/3.2 NokiaC5-00.2/081.003; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.32 Mobile Safari/533.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.3; U; Series60/3.2 NokiaN79-1/32.001; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.3; U; Series60/3.2 Nokia6220c-1/06.101; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/413 (KHTML, like Gecko) Safari/413`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.3; Series60/3.2 NokiaC5-00.2/071.003; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.26 Mobile Safari/533.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.3; Series60/3.2 NokiaE72-1/081.003; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.32 Mobile Safari/533.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.3; Series60/3.2 NokiaC5-00/061.005; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/525 (KHTML, like Gecko) Version/3.0 BrowserNG/7.2.6.2 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.4; Series60/5.0 NokiaX6-00/40.0.002; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.33 Mobile Safari/533.4 3gpp-gb`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.4; Series60/5.0 Nokia5800d-1/60.0.003; Profile/MIDP-2.1 Configuration/CLDC-1.1 AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.33 Mobile Safari/533.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.4; Series60/5.0 NokiaC5-03/12.0.023; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/525 (KHTML, like Gecko) Version/3.0 BrowserNG/7.2.6.9 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.4; Series60/5.0 Nokia5228/40.1.003; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/525 (KHTML, like Gecko) Version/3.0 BrowserNG/7.2.7.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.4; Series60/5.0 Nokia5230/51.0.002; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.33 Mobile Safari/533.4 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.4; Series60/5.0 Nokia5530c-2/32.0.007; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/525 (KHTML, like Gecko) Version/3.0 BrowserNG/7.2.6.9 3gpp-gba`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SymbianOS/9.4; Series60/5.0 NokiaN97-4/30.0.004; Profile/MIDP-2.1 Configuration/CLDC-1.1) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.28 3gpp-gba`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Symbian`:      `9.4`,
				`Webkit`:       `533.4`,
				`NokiaBrowser`: `7.3.1.28`,
			},
			`NokiaN97-4`,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Nokia; 7 Mozart T8698)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; 710)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Nokia; 800)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Nokia; 800C)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Nokia; 800C; Orange)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Nokia; 900)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Nokia; HD7 T9292)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; LG E-900)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 610)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 710)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 710; Orange)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 710; T-Mobile)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 710; Vodafone)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 800)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 800) UP.Link/5.1.2.6`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 800; Orange)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 800; SFR)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 800; T-Mobile)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 800; vodafone)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Nokia; Lumia 800c)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Lumia 900)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Nokia; Lumia 920)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 920)`,
		expectedResult{
			true,
			false,
			map[string]string{
				`IE`:               `10.0`,
				`Windows Phone OS`: `8.0`,
				`Trident`:          `6.0`,
			},
			`Lumia 920`,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Nokia; lumia800)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Nokia 610)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Nokia 710)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Nokia 800)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Nokia 800C)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; NOKIA; Nokia 900)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Nokia; Nokia)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Nokia; SGH-i917)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Nokia; TITAN X310e)`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Windows Phone OS`: `7.5`,
				`Trident`:          `5.0`,
			},
			`TITAN X310e`,
		},
	},
	{
		`Mozilla/5.0 (Mobile; Windows Phone 8.1; Android 4.0; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 520) like iPhone OS 7_0_3 Mac OS X AppleWebKit/537 (KHTML, like Gecko) Mobile Safari/537`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	//Odays
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; de-de; LOOX Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; de-de; LOOX Plus Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.X; de-de; XENO10 Build/ODYS XENO 10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2.2; de-de; ODYS Space Build/I700T_P7_T04_TSCL_FT_R_0_03_1010_110623) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; de-de; ODYS-EVO Build/ODYS-EVO) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.X; de-de; Xelio 10 Pro Build/ODYS_Xelio) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; NEO_QUAD10 Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.X; de-de; Xelio10Pro Build/ODYS_Xelio) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.1; en-us; ODYS-Xpress Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; de-de; XELIO7PHONETAB Build/IMM76I) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; XELIO10EXTREME Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; de-de; XELIO Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30 Mobile UCBrowser/3.2.1.441`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; XELIOPT2 Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; nl-nl; ODYS-NOON Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//OverMax
	{
		`OV-SteelCore(B) Mozilla/5.0 (iPad; CPU OS 5_0_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; pl-pl; OV-SteelCore Build/ICS.g08refem611.20121010) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//YONESTablet
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; pl-pl; BC1077 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	// Pantech @todo: Research http://www.pantech.com/
	{
		`PANTECH-C790/JAUS08312009 Browser/Obigo/Q05A Profile/MIDP-2.0 Configuration/CLDC-1.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2.1; ko-kr; SKY IM-A600S Build/FRG83) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; en-us; ADR8995 4G Build/GRI40) AppleWebKit/533.1 (KHTML like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.2.1; en-us; PantechP4100 Build/HTK75) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Philips
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; W732 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; Philips W336 Build/IMM76D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.45 Mobile Safari/537.36 OPR/15.0.1162.59192`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Philips_T3500/V1 Linux/3.4.5 Android/4.2.2 Release/03.26.2013 Browser/AppleWebKit534.30 Mobile Safari/534.30 MBBMS/2.2 System/Android 4.2.2;`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Philips W3568 Build/Philips_W3568) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; ru-ru; Philips W832 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux;U;Android 4.2.2;es-us;Philips S388 Build/JDQ39) AppleWebkit/534.30 (HTML,like Gecko) Version/4.0 Mobile Safari/534.30;`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; Philips W536 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux;U;Android 4.2.2;es-us;Philips S308 Build/JDQ39) AppleWebkit/534.30 (HTML,like Gecko) Version/4.0 Mobile Safari/534.30;`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; ru-ru; Philips-W8500 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; ru; Philips W8510 Build/JDQ39) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 UCBrowser/9.8.9.457 U3/0.8.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; ru-RU; Philips W3568 Build/Philips W3568) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.1 Mobile Safari/534.30;`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Philips S388 Build/JDQ39) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; Philips S388 Build/JDQ39) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; Build/PI3100.00.00.24) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.1; PI7100_93 Build/PI7100.C.00.00.11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.92 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//PointOfView
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; POV_TAB-PROTAB30-IPS10 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Prestigio
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-gb; PMP5297C_QUAD Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; sk-sk; PMP7100D3G Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.4`,
				`Build`:   `IMM76D`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			`PMP7100D3G`,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; sk-sk; PMP7280C3G Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 MobilSafari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//PROSCAN
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; PLT8088 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.1`,
				`Build`:   `JRO03H`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	//PyleAudio
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; PTBL92BC Build/IMM76D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.59 Mobile Safari/537.36`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.4`,
				`Build`:   `IMM76D`,
				`Webkit`:  `537.36`,
				`Chrome`:  `31.0.1650.59`,
			},
			``,
		},
	},
	//RockChip
	{
		`Mozilla/5.0 (Linux; U; Android 2.2.1; hu-hu; RK2818, Build/MASTER) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android Android 2.1-RK2818-1.0.0; zh-cn; MD701 Build/ECLAIR) AppleWebKit/530.17 (KHTML like Gecko) Version/4.0 Mobile Safari/530.17`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//RossMoor
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.1; ru-ru; RM-790 Build/JOP40D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//QMobile @ref: http://www.qmobile.com.pk/complete_range.php#
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; en-us; A2 Build/GRK39F) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	//Samsung
	{
		`MQQBrowser/4.0/Mozilla/5.0 (Linux; U; Android 3.2; zh-cn; GT-P6800 Build/HTJ85B) AppleWebKit/533.1 (KHTML, like Gecko) Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			map[string]string{
				`MQQBrowser`: `4.0`,
			},
			``,
		},
	},
	{
		`SAMSUNG-SGH-P250-ORANGE/P250BVHH8 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`SAMSUNG-GT-B2710/B271MCXKF1 SHP/VPP/R5 Dolfin/2.0 QTV/5.3 SMM-MMS/1.2.0 profile/MIDP-2.1 configuration/CLDC-1.1 OPN-B`,
		expectedResult{
			true,
			false,
			map[string]string{`Dolfin`: `2.0`},
			``,
		},
	},
	{
		`SAMSUNG-SGH-D900i/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`SAMSUNG-GT-S5233T/S5233TXEJE3 SHP/VPP/R5 Jasmine/0.8 Qtv5.3 SMM-MMS/1.2.0 profile/MIDP-2.1 configuration/CLDC-1.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SAMSUNG; SAMSUNG-GT-S5380D/S5380FXXKL3; U; Bada/2.0; ru-ru) AppleWebKit/534.20 (KHTML, like Gecko) Dolfin/3.0 Mobile HVGA SMM-MMS/1.2.0 OPN-B`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`SAMSUNG-GT-C3312/1.0 NetFront/4.2 Profile/MIDP-2.0 Configuration/CLDC-1.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 1.5; de-de; Galaxy Build/CUPCAKE) AppleWebKit/528.5 (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`SAMSUNG-GT-S3650/S3650XEII3 SHP/VPP/R5 Jasmine/1.0 Nextreaming SMM-MMS/1.2.0 profile/MIDP-2.1 configuration/CLDC-1.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`JUC (Linux; U; 2.3.6; zh-cn; GT-S5360; 240*320) UCWEB7.9.0.94/140/352`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SAMSUNG; SAMSUNG-GT-S5250/S5250XEKJ3; U; Bada/1.0; ru-ru) AppleWebKit/533.1 (KHTML, like Gecko) Dolfin/2.0 Mobile WQVGA SMM-MMS/1.2.0 NexPlayer/3.0 profile/MIDP-2.1 configuration/CLDC-1.1 OPN-B`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/7.0; SAMSUNG; SGH-i917)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (SAMSUNG; SAMSUNG-GT-S8530/S8530XXJKA; U; Bada/1.2; cs-cz) AppleWebKit/533.1 (KHTML, like Gecko) Dolfin/2.2 Mobile WVGA SMM-MMS/1.2.0 OPN-B`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 1.6; ru-ru; Galaxy Build/Donut) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.1-update1; ru-ru; GT-I5500 Build/ERE27) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; en-us; GALAXY_Tab Build/MASTER) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; ja-jp; SC-01C Build/FROYO) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	// @about FROYO: http://gizmodo.com/5543853/what-is-froyo
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; fr-fr; GT-I9000 Build/FROYO) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2.1; zh-cn; SCH-i909 Build/FROYO) UC AppleWebKit/534.31 (KHTML, like Gecko) Mobile Safari/534.31`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; ja-jp; SC-01C Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; en-gb; GT-P1000 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; el-gr; GT-I9001 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; en-ca; SGH-I896 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.4; es-us; GT-S5660L Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1 MicroMessenger/4.5.1.261`,
		expectedResult{
			true,
			false,
			map[string]string{`MicroMessenger`: `4.5.1.261`},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; en-us; GT-S5660 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; ru-ru; GT-S6102 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; pt-br; GT-S5367 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; fr-fr; GT-S5839i Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; en-gb; GT-S7500 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; en-gb; GT-S5830 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; es-us; GT-B5510L Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; pl-pl; GT-I9001-ORANGE/I9001BVKPC Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; en-us; GT-I8150 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; nl-nl; GT-I9070 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; en-gb; GT-S5360 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; es-us; GT-S6102B Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; en-us; GT-S5830i Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; ru-ru; GT-I8160 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; ru-ru; GT-S6802 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.7; ru-ru; GT-S5830 Build/GRWK74; LeWa_ROM_Cooper_12.09.21) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; ru-ru; GT-N7000 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.0.1; en-us; GT-P7100 Build/HRI83) AppleWebkit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.2; he-il; GT-P7300 Build/HTJ85B) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.2; en-gb; GT-P6200 Build/HTJ85B) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-gb; GT-I9100 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; GT-I9100G Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; nl-nl; GT-P5100 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7`,
		expectedResult{
			true,
			false,
			map[string]string{`Chrome`: `16.0.912.75`},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.3; SGH-T989 Build/IML74K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			map[string]string{`Chrome`: `18.0.1025.166`},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; ru-ru; GT-P5100 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; GT-I9300 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; SPH-D710 Build/IMM76I) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; zh-cn; GT-I9300 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-gb; GT-I9300-ORANGE/I9300BVBLG2 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; th-th; GT-I9300T Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; ru-ru; GT-I9100 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us ; GT-I9100 Build/IMM76D) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1/UCBrowser/8.4.1.204/145/355`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-gb; GT-N7000 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; th-th; GT-P6800 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; SAMSUNG-SGH-I747 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; es-es; GT-P5110 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; GT-P5110 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; zh-cn; SAMSUNG-GT-S7568_TD/1.0 Android/4.0.4 Release/07.15.2012 Browser/AppleWebKit534.30 Build/IMM76D) ApplelWebkit/534.30 (KHTML,like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			map[string]string{`Android`: `4.0.4`},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; ru-ru; GT-P3100 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; tr-tr; GT-P3105 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-ca; GT-N8010 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; ru-ru; GT-S7562 Build/IMM76I) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; GT-N7100 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; GT-N7100 Build/JZO54K) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.123 Mobile Safari/537.22 OPR/14.0.1025.52315`,
		expectedResult{
			true,
			false,
			map[string]string{
				`Build`:  `JZO54K`,
				`Webkit`: `537.22`,
				`Opera`:  `14.0.1025.52315`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; zh-hk; GT-N7105 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; ru-ru; GT-N8000 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; SGH-i747M Build/JRO03L) AppleWebKit/535.19(KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		//@see rules.go:35
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-gb; Galaxy Nexus - 4.1.1 - with Google Apps - API 16 - 720x1280 Build/JRO03S) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; GT-I8262 Build/JZO54K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; it-it; Galaxy Nexus Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; SGH-I777 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; GT-S7710 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; GT-I9082 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; SGH-T999L Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; GT-P5210 Build/JDQ39) AppleWebKit/537.36 (KHTML, Like Gecko) Chrome/27.0.1453.90 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; en-us; SAMSUNG GT-I9200 Build/JDQ39) AppleWebKit/535.19 (KHTML, like Gecko) Version/1.0 Chrome/18.0.1025.308 Mobile Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; zh-cn; SCH-I959 Build/JDQ39) AppleWebKit/535.19 (KHTML, like Gecko) Version/1.0 Chrome/18.0.1025.308 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; nl-nl; SM-T310 Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; en-us; SAMSUNG SM-P600 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Version/1.5 Chrome/28.0.1500.94 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-gb; GT-N5100 Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; en-us; SAMSUNG SM-T530NU Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Version/1.5 Chrome/28.0.1500.94 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; SM-T800 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.135 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; fr-fr; SAMSUNG SM-T800 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Version/1.5 Chrome/28.0.1500.94 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; SM-T700 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.517 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; CETUS)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; Focus I917 By TC)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; Focus i917)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; FOCUS S)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; GT-I8350)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; GT-i8700)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; GT-S7530)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; Hljchm`s Wp)",
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; I917)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; OMNIA 7)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; OMNIA7 By MWP_HS)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; OMNIA7)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; OMNIA7; Orange)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; SGH-i677)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; SGH-i917)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; SGH-i917.)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; SGH-i917R)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; SGH-i937)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG; SMG-917R)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG_blocked_blocked_blocked; OMNIA7; Orange)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SAMSUNG_blocked_blocked_blocked_blocked; OMNIA7; Orange)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; SUMSUNG; OMNIA 7)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Windows NT 6.2; ARM; Trident/7.0; Touch; rv:11.0; WPDesktop; SAMSUNG; GT-I8750) like Gecko`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Windows NT 6.2; ARM; Trident/7.0; Touch; rv:11.0; WPDesktop; GT-I8750) like Gecko`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	// simvalley
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.4; de-de; SP-80 Build/GRJ22) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	// sony
	{
		`SonyEricssonK800i/R1AA Browser/NetFront/3.3 Profile/MIDP-2.0 Configuration/CLDC-1.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.1-update1; es-ar; SonyEricssonE15a Build/2.0.1.A.0.47) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.1-update1; pt-br; SonyEricssonU20a Build/2.1.1.A.0.6) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; en-au; SonyEricssonX10i Build/3.0.1.G.0.75) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.4; ru-ru; SonyEricssonST18i Build/4.0.2.A.0.62) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.4; hr-hr; SonyEricssonST15i Build/4.0.2.A.0.62) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.4; sk-sk; SonyEricssonLT15i Build/4.0.2.A.0.62) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.7; th-th; SonyEricssonST27i Build/6.0.B.3.184) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.7; de-de; SonyEricssonST25i Build/6.0.B.3.184) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; pt-br; Xperia Tablet S Build/TID0092) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.3`,
				`Build`:   `TID0092`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.3; LT18i Build/4.1.A.0.562) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.3; Sony Tablet S Build/TISU0R0110) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; es-es; Sony Tablet S Build/TISU0143) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-gb; SonyEricssonLT18i Build/4.1.B.0.587) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; fr-ch; SonyEricssonSK17i Build/4.1.B.0.587) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; SonyEricssonLT26i Build/6.1.A.2.45) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; LT22i Build/6.1.B.0.544) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; vi-vn; SonyEricssonLT22i Build/6.1.B.0.544) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; ST23i Build/11.0.A.5.5) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; ST23i Build/11.0.A.2.10) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; LT28h Build/6.1.E.3.7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.94 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; SGPT13 Build/TJDS0170) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; ja-jp; SonySO-03E Build/10.1.E.0.265) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.2`,
				`Build`:   `10.1.E.0.265`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; LT26w Build/6.2.B.1.96) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.72 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; SGP321 Build/10.3.1.A.0.33) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.58 Safari/537.31`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.2.2`,
				`Build`:   `10.3.1.A.0.33`,
				`Webkit`:  `537.31`,
				`Chrome`:  `26.0.1410.58`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.3; zh-cn; XL39h Build/14.2.A.1.136) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; sv-se; C5503 Build/10.1.1.A.1.273) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; C5502 Build/10.1.1.A.1.310) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; zh-cn; SonyL39t Build/14.1.M.0.202) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; zh-cn; L39u Build/14.1.n.0.63) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; zh-tw; M35c Build/12.0.B.5.37) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; M35c Build/12.0.B.2.42) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.90 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; zh-CN; M35t Build/12.0.C.2.42) AppleWebKit/534.31 (KHTML, like Gecko) UCBrowser/9.3.2.349 U3/0.8.0 Mobile Safari/534.31`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; D6502 Build/17.1.A.2.69) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; D6503 Build/17.1.A.0.504) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; D6543 Build/17.1.A.2.55) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; D2004 Build/20.0.A.0.29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.94 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.3; en-gb; D2005 Build/20.0.A.1.12) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; D2104 Build/20.0.B.0.84) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; D2105 Build/20.0.B.0.74) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.170 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.3; pt-br; D2114 Build/20.0.B.0.85) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; D2302 Build/18.0.B.1.23) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.138 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.3; zh-cn; S50h Build/18.0.b.1.23) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 UCBrowser/9.6.3.413 U3/0.8.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; D2303 Build/18.0.C.1.13) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.138 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; D2305 Build/18.0.A.1.30) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.138 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; D2306 Build/18.0.C.1.7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.136 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; D5303 Build/19.0.1.A.0.207) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; D5306 Build/19.1.A.0.264) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.136 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.3; zh-CN; XM50h Build/19.0.D.0.269) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 UCBrowser/9.7.6.428 U3/0.8.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.3; zh-cn; XM50t Build/19.0.C.2.59) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; D5322 Build/19.0.D.0.253) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.131`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.3; zh-cn; M51w Build/14.2.A.1.146) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; M51w Build/14.2.A.1.146) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.136 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.1; D5102 Build/18.2.A.0.9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.1; D5103 Build/18.1.A.0.11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.92 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.1; D5106 Build/18.1.A.0.11) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.3; en-gb; C6902 Build/14.2.A.1.136) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30 GSA/3.2.17.1009776.arm`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; es-es; C6943 Build/14.1.G.2.257) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; C6943 Build/14.3.A.0.681) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; SGP412 Build/14.1.B.3.320) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Mobile Safari/537.36`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1; en-us; SonySGP321 Build/10.2.C.0.143) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; SGP351 Build/10.1.1.A.1.307) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.3; SGP341 Build/10.4.B.0.569) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; SGP511 Build/17.1.A.2.36) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.122 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; SGP512 Build/17.1.A.2.36) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.122 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; fr-ch; SGP311 Build/10.1.C.0.344) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; SGP312 Build/10.1.C.0.344) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.4.2; de-de; SGP521 Build/17.1.A.2.69) AppleWebKit/537.16 (KHTML, like Gecko) Version/4.0 Safari/537.16`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.4.2; zh-cn; SGP541 Build/17.1.A.2.36) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; SGP551 Build/17.1.A.2.72) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`SonyEricssonU5i/R2CA; Mozilla/5.0 (SymbianOS/9.4; U; Series60/5.0 Profile/MIDP-2.1 Configuration/CLDC-1.1) AppleWebKit/525 (KHTML, like Gecko) Version/3.0 Safari/525`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`SonyEricssonU5i/R2AA; Mozilla/5.0 (SymbianOS/9.4; U; Series60/5.0 Profile/MIDP-2.1 Configuration/CLDC-1.1) AppleWebKit/525 (KHTML, like Gecko) Version/3.0 Safari/525`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/4.0 (PDA; PalmOS/sony/model prmr/Revision:1.1.54 (en)) NetFront/3.0`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Opera/9.80 (Linux mips; U; InettvBrowser/2.2 (00014A;SonyDTV115;0002;0100) KDL40EX720; CC/BEL; en) Presto/2.7.61 Version/11.00`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Opera/9.80 (Linux armv7l; HbbTV/1.1.1 (; Sony; KDL32W650A; PKG3.211EUA; 2013;); ) Presto/2.12.362 Version/12.11`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	//Skk
	{
		`Mozilla/5.0 (Linux; U; Android 4.1; en-us; CYCLOPS Build/F10) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Storex
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; eZee_Tab903 Build/JRO03H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Safari/537.36`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.1`,
				`Build`:   `JRO03H`,
				`Webkit`:  `537.36`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.1; eZee'Tab785 Build/JRO03C) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.59 Safari/537.36`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.1`,
				`Build`:   `JRO03C`,
				`Webkit`:  `537.36`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.3; eZee'Tab971 Build/IML74K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.0.3`,
				`Build`:   `IML74K`,
				`Webkit`:  `535.19`,
			},
			``,
		},
	},
	//Tecno
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-us; TECNO P9 Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Teclast
	{
		`Mozilla/5.0 (Linux; Android 4.4.2; P98 3G\xE5\x85\xAB\xE6\xA0\xB8(A3HY) Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Telstra
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.7; en-au; T-Hub2 Build/TVA301TELBG3) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//texet @info: http://www.texet.ru/tablet/
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.4; ru-ru; TM-7021 Build/GB.m1ref.20120116) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `2.3.4`,
				`Safari`:  `4.0`,
				`Webkit`:  `533.1`,
			},
			`TM-7021`,
		},
	},
	//Tolino
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; tolino tab 7 Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.2.2`,
				`Safari`:  `4.0`,
				`Webkit`:  `534.30`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; tolino tab 8.9 Build/JDQ39) AppleWebKit/534.30 (KHTML like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.2.2`,
				`Safari`:  `4.0`,
				`Webkit`:  `534.30`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; tolino tab 7 Build/JDQ39) AppleWebkit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36 OPR/18.0.1290.67495`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; tolino tab 7 Build/JDQ39) AppleWebkit/537.36 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Toshiba
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; TOSHIBA; TSUNAGI)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	// @ref: http://www.toshiba.co.uk/discontinued-products/folio-100/
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; it-it; TOSHIBA_FOLIO_AND_A Build/TOSHIBA_FOLIO_AND_A) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `2.2`,
				`Webkit`:  `533.1`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	// Trekstor
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; ST70408-1 Build/JDQ39) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.58 Safari/537.31`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.2.2`,
				`Webkit`:  `537.31`,
				`Build`:   `JDQ39`,
				`Chrome`:  `26.0.1410.58`,
			},
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; engb; Build/IMM76D) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A405 Safari/7534.48.3 SurfTab_7.0`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; VT10416-2 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; de-de; ST10216-2A Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30;SurfTab_10.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},

	// Visture
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-gb; V97 HD Build/LR-97JC) Apple WebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; Visture V4 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; ru-ru; Visture V4 HD Build/Visture V4 HD) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; es-es; Visture V5 HD Build/Visture V5 HD) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; ru-ru; Visture V10 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Versus
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; VS-TOUCHPAD 9 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-gb; Versus Touchpad 9.7 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-gb; CnM-TOUCHPAD7 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30 BMID/E67A45B1AB`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-gb; CnM TouchPad 7DC Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30 TwonkyBeamBrowser/3.3.5-95 (Android 4.1.1; rockchip CnM TouchPad 7DC Build/meizhi_V2.80.wifi8723.20121225.b11c800)`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`OneBrowser/3.5/Mozilla/5.0 (Linux; U; Android 4.0.4; en-gb; TOUCHPAD 7 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-gb; TOUCHTAB Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.1`,
				`Build`:   `JRO03H`,
				`Webkit`:  `534.30`,
				`Safari`:  `4.0`,
			},
			``,
		},
	},
	// Viewsonic
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; de-de; ViewPad 10e Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; de-de; ViewPad 10e Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; it-it; ViewPad7 Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2.1; en-ca; ViewSonic VB733 Build/FRG83) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.2; en-gb; ViewPad7X Build/HTJ85B) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; pt-br; ViewPad 10S Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; VB100a Pro Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Vodafone
	{
		`Mozilla/5.0 (Linux; U; Android 3.2; hu-hu; SmartTab10-MSM8260-V02d-Dec022011-Vodafone-HU) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.0.3; SmartTabII10 Build/IML74K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; fr-fr; SmartTAB 1002 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; de-de, SmartTabII7 Build/A2107A_A404_107_055_130124_VODA) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Vonino
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; en-us; Sirius_Evo_QS Build/JDQ39) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-gb; Q8 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Wolfgang
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.1; nl-nl; AT-AS45q2 Build/JOP40D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	//Xoro
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; PAD 9720QR Build/PAD 9719QR) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; PAD720 Build/JDQ39) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//ZTE
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.1;zh-cn; ZTE V987 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30;`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.5; pt-pt; Blade Build/tejosunhsine) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; ZTE; N880e_Dawoer_Fulllock; China Telecom)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; ZTE; V965W)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; ZTE; Windows Phone - Internet 7; SFR)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	//Zync
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; en-us ; Z909 Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1/UCBrowser/8.4.1.204/145/444`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Console
	//Nintendo Wii:
	{
		`Mozilla/5.0 (Nintendo WiiU) AppleWebKit/534.52 (KHTML, like Gecko) NX/{Version No} NintendoBrowser/{Version No}.US`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	//Sony PlayStation:
	{
		`Mozilla/5.0 (PLAYSTATION 3 4.21) AppleWebKit/531.22.8 (KHTML, like Gecko)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	//Microsoft Xbox:
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; Xbox)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// WTF? Must investigate.
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Xbox)`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// Liebao Browser
	{
		`Mozilla/5.0 (X11; Linux i686) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/32.0.1700.102 Chrome/32.0.1700.102 Safari/537.36`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// Other
	{
		`Mozilla/5.0 (X11; Linux i686) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/32.0.1700.102 Chrome/32.0.1700.102 Safari/537.36`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Windows NT 6.1; WOW64; rv:24.0) Gecko/20100101 Firefox/24.0`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Windows NT 6.1; WOW64; rv:18.0) Gecko/20100101 Firefox/18.0 AlexaToolbar/psPCtGhf-2.2`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:22.0) Gecko/20100101 Firefox/22.0`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (X11; Linux ppc; rv:17.0) Gecko/20130626 Firefox/17.0 Iceweasel/17.0.7`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (X11; Linux) AppleWebKit/535.22+ Midori/0.4`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Macintosh; U; Intel Mac OS X; en-us) AppleWebKit/535+ (KHTML, like Gecko) Version/5.0 Safari/535.20+ Midori/0.4`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.20 Safari/537.36  OPR/15.0.1147.18 (Edition Next)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Windows NT 6.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.94 Safari/537.36`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Windows NT 5.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.94 Safari/537.36`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Windows NT 5.2; WOW64; rv:21.0) Gecko/20100101 Firefox/21.0`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Opera/9.80 (Windows NT 5.2; WOW64) Presto/2.12.388 Version/12.14`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Windows NT 6.1; WOW64; rv:19.0) Gecko/20100101 Firefox/19.0`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (X11; FreeBSD amd64; rv:14.0) Gecko/20100101 Firefox/14.0.1`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.2; Win64; x64; Trident/6.0; Touch; .NET4.0E; .NET4.0C; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Tablet PC 2.0; MASMJS)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0; MASMJS)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; WOW64; Trident/6.0; Touch; MASMJS)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Opera/9.80 (Windows NT 6.2; WOW64; MRA 8.0 (build 5784)) Presto/2.12.388 Version/12.11`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// IE 10
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// IE 11 @todo: Trident(.*)rv.(\d+)\.(\d+)
	{
		`Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv 11.0) like Gecko`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// TV
	{
		`Mozilla/5.0 (Unknown; Linux armv7l) AppleWebKit/537.1+ (KHTML, like Gecko) Safari/537.1+ HbbTV/1.1.1 ( ;LGE ;NetCast 4.0 ;03.20.30 ;1.0M ;)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`HbbTV/1.1.1 (;Panasonic;VIERA 2012;1.261;0071-3103 2000-0000;)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Opera/9.80 (Linux armv7l; HbbTV/1.1.1 (; Sony; KDL32W650A; PKG3.211EUA; 2013;); ) Presto/2.12.362 Version/12.11`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Opera/9.80 (Linux mips ; U; HbbTV/1.1.1 (; Philips; ; ; ; ) CE-HTML/1.0 NETTV/3.2.1; en) Presto/2.6.33 Version/10.70`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	//Generic
	// Carrefour tablet
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; CT1020W Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.94 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	// @comment: Pipo m6pro tablet
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; M6pro Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	// https://github.com/varnish/varnish-devicedetect/issues/17
	// To be researched.
	{
		`Mozilla/5.0 (Linux; Android 4.2.2; M6pro Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.141 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`MobileSafari/9537.53 CFNetwork/672.1.13 Darwin/13.1.0`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Appcelerator Titanium/3.2.2.GA (iPod touch/6.1.6; iPhone OS; en_US;)`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Opera Coast/3.0.3.78307 CFNetwork/672.1.15 Darwin/14.0.0`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`@see https://github.com/Shaked/gomobiledetect/issues/1 Mozilla/5.0 (Linux; Android 4.0.3; ALUMIUM10 Build/IML74K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.99 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.1; en-us; JY-G3 Build/JOP40D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`@see https://github.com/Shaked/gomobiledetect/issues/1 Mozilla/5.0 (Linux; U; Android 4.1.1; hu-hu; M758A Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`@see https://github.com/Shaked/gomobiledetect/issues/1 Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; EVOTAB Build/IMM76I) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Opera/9.80 (Series 60; Opera Mini/6.5.29260/29.3417; U; ru) Presto/2.8.119 Version/11.10`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Opera/9.80 (Android; Opera Mini/6.5.27452/29.3417; U; ru) Presto/2.8.119 Version/11.10`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Opera/9.80 (iPhone; Opera Mini/7.1.32694/27.1407; U; en) Presto/2.8.119 Version/11.10`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	// New Opera
	{
		`Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.17 (KHTML, like Gecko) Chrome/24.0.1312.60 Safari/537.17 OPR/14.0.1025.52315`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	// Unknown yet
	// Looks like Chromebook
	{
		`Mozilla/5.0 (X11; CrOS armv7l 4920.83.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.103 Safari/537.36`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_2 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8H7 Safari/6533.18.5`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; fr-fr; GT-I9070 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Opera/9.80 (Android 2.3.7; Linux; Opera Mobi/46154) Presto/2.11.355 Version/12.10`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; it-it; DATAM819HD_C Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; nl-nl; SGPT12 Build/TID0142) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; cm_tenderloin Build/IMM76L; CyanogenMod-9) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// Acer Iconia Tab
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; fr-fr; A210 Build/JRO03H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (iPad; CPU OS 6_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10B141 Safari/8536.25`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; en-gb; GT-I8150 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// Mercurio Browser
	{
		`Mozilla/5.0 (iPad; CPU OS 6_0_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Mercury/7.2 Mobile/10A523 Safari/8536.25`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Opera/9.80 (Android 2.3.7; Linux; Opera Tablet/46154) Presto/2.11.355 Version/12.10`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// sdk
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.2; en-us; sdk Build/MASTER) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// sdk
	{
		`Mozilla/5.0 (Linux; U; Android 4.2; en-us; sdk Build/JB_MR1) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; GT-P7510 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Opera/9.80 (X11; Linux zbov) Presto/2.11.355 Version/12.10`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// 7" Verso Android tablet
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; en-gb; TOUCHPAD 7 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (iPhone; U; CPU OS 4_2_1 like Mac OS X) AppleWebKit/532.9 (KHTML, like Gecko) Version/5.0.3 Mobile/8B5097d Safari/6531.22.7`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// sony xperia tablet s unforts
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-gb; SGPT12 Build/TID0142) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Nintendo 3DS; U; ; en) Version/1.7498.US`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// PocketBook
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.7; en-ru; PocketBook A10 3G Build/GRJ90) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	// PocketBook IQ701 (tablet)
	{
		`Mozilla/5.0 (Linux; U; Android 2.0.6_b1; ru-ru Build/ECLAIR) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// It is a tablet with calling
	{
		`Mozilla/5.0 (Linux; Android 4.0.4; z1000 Build/IMM76D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// HP touch pad running android cyanogenmod
	{
		`Mozilla/5.0 (Linux; Android 4.0.3; cm_tenderloin Build/GWK74) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// My device is tablet but its detected as phone
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.3; en-us; Android for Techvision TV1T808 Board Build/GRI40) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`BlackBerry8520/5.0.0.592 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/136`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// its a lenovo tablet 2 with windows 8 pro
	{
		`Mozilla/5.0 (Windows NT 6.2) AppleWebKit/537.17 (KHTML, like Gecko) Chrome/24.0.1312.56 Safari/537.17`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.1; ru-ru; LG-V900 Build/HMJ37) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; HTC; Windows Phone 8S by HTC)`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// MS Surface RT tablet actually!
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (PLAYSTATION 3 4.11) AppleWebKit/531.22.8 (KHTML, like Gecko)`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// Wrong detection - 7-inch tablet was detected as a phone. Android 3.2.1, native browser
	{
		`Mozilla/5.0 (Linux; U; Android 3.2; ru-ru; V9S_V1.4) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// Nope, its a Microsoft Surface tablet	 running Windows RT (8) with MSIE 10
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// Tablet!
	{
		`Mozilla/5.0 (Linux; U; Android 2.2; es-es; Broncho N701 Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// its a Microsoft surface rt (tablet)
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch; WebView/1.0)`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// Nextbook 7SE Tablet
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; Next7P12 Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// allview alldro speed tablet, android ics, opera mobile
	{
		`Opera/9.80 (X11; Linux zbov) Presto/2.11.355 Version/12.10`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// Its a surface in portrait
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 2.3.6; es-es; SAMSUNG GT-S5830/S5830BUKT2 Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 3.2.1; en-gb;HTC_Flyer_P512 Build/HTK75C) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// Am ramas la pozitia: 207
	// Android on Windows :) www.socketeq.com
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; full Android on Microsoft Windows, pad, pc, n*books Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// TV
	{
		`Opera/9.80 (Linux mips; U; InettvBrowser/2.2 (00014A;SonyDTV115;0002;0100) KDL40EX720; CC/BEL; en) Presto/2.7.61 Version/11.00`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Android; Mobile; rv:18.0) Gecko/18.0 Firefox/18.0`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// Maxthon
	{
		`Mozilla/5.0 (Windows NT 6.1) AppleWebKit/535.12 (KHTML, like Gecko) Maxthon/3.0 Chrome/18.0.966.0 Safari/535.12`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Opera/9.80 (Windows NT 5.1; U; Edition Yx; ru) Presto/2.10.289 Version/12.02`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.2; en-us; sdk Build/JB_MR1) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; Windows Phone 6.5.3.5)`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`PalmCentro/v0001 Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; PalmSource/Palm-D061; Blazer/4.5) 16;320x320`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0)`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; Microsoft; XDeviceEmulator)`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// @todo: research N880E
	{
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; MAL; N880E; China Telecom)`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Opera/9.80 (Series 60; Opera Mini/7.0.29482/28.2859; U; ru) Presto/2.8.119 Version/11.10`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Opera/9.80 (S60; SymbOS; Opera Mobi/SYB-1202242143; U; en-GB) Presto/2.10.254 Version/12.00`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-au; 97D Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Opera/9.80 (Android; Opera Mini/7.0.29952/28.2647; U; ru) Presto/2.8.119 Version/11.10`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Opera/9.80 (Android; Opera Mini/6.1.25375/28.2555; U; en) Presto/2.8.119 Version/11.10`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Opera/9.80 (Mac OS X; Opera Tablet/35779; U; en) Presto/2.10.254 Version/12.00`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Android; Tablet; rv:10.0.4) Gecko/10.0.4 Firefox/10.0.4 Fennec/10.0.4`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Mozilla/5.0 (Android; Tablet; rv:18.0) Gecko/18.0 Firefox/18.0`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Opera/9.80 (Linux armv7l; Maemo; Opera Mobi/14; U; en) Presto/2.9.201 Version/11.50`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	{
		`Opera/9.80 (Android 2.2.1; Linux; Opera Mobi/ADR-1207201819; U; en) Presto/2.10.254 Version/12.00`,
		expectedResult{
			true,
			false,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-us; sdk Build/JRO03E) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			false,
			false,
			nil,
			TEST_SKIPPED,
		},
	},
	// Blaupunkt Endeavour 1010
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; Endeavour 1010 Build/ONDA_MID) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; de-de; Tablet-PC-4 Build/ICS.g08refem618.20121102) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Tagi tablets
	{
		`Mozilla/5.0 (Linux; U; Android 4.2.2; de-de; Tagi Tab S10 Build/8089) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	//Bot
	{
		`Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`grub-client-1.5.3; (grub-client-1.5.3; Crawl your own stuff with http://grub.org)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Googlebot-Image/1.0`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`Python-urllib/2.5`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
	{
		`facebookexternalhit/1.0 (+http://www.facebook.com/externalhit_uatext.php)`,
		expectedResult{
			false,
			false,
			nil,
			``,
		},
	},
}

type uaListResult struct {
	success bool
	message string
	skipped bool
}

func TestUaList(t *testing.T) {
	t.Parallel()
	runtime.GOMAXPROCS(runtime.NumCPU())
	chn := make(chan *uaListResult, len(uaListTests))
	for idx, test := range uaListTests {
		go func(idx int, userAgent string, er expectedResult, chn chan *uaListResult) {
			result := &uaListResult{
				true,
				"",
				false,
			}

			if TEST_SKIPPED == er.model {
				result.success = false
				result.message = userAgent
				result.skipped = true
			} else {
				detect := NewMobileDetect(httpRequest, nil)
				detect.SetUserAgent(userAgent)
				isMobile := detect.IsMobile()

				if er.isMobile != isMobile {
					result.success = false
					result.message += fmt.Sprintf("%d: For userAgent %s\n expected result is mobile: %s got %s\n", idx, userAgent, er.isMobile, isMobile)
				}

				isTablet := detect.IsTablet()
				if er.isTablet != isTablet {
					result.success = false
					result.message += fmt.Sprintf("%d: For userAgent %s\n expected result is tablet: %s got %s\n", idx, userAgent, er.isTablet, isTablet)
				}

				for name, v := range er.version {
					actualVersion := detect.Version(name)
					if v != actualVersion {
						t.Errorf("[%s] expected version: %s, actual version: %s", name, v, actualVersion)
					}
				}
			}

			chn <- result
		}(idx, test.userAgent, test.er, chn)
	}
	for i := 0; i < len(uaListTests); i++ {
		result := <-chn

		if false == result.success {
			if result.skipped {
				go func(t *testing.T) {
					log.Printf("Skipped: %s", result.message) //t.Skipf doesn't work on Go1
				}(t)
			} else {
				t.Error(result.message)
			}
		}

		if result.success && "done" == result.message {
			break
		}
	}
}

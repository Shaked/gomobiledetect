package gomobiledetect

import (
	"fmt"
	"runtime"
	"testing"
)

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
		`Mozilla/5.0 (Linux; U; Android 4.1.1; en-gb; PadFone 2 Build/JRO03L) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
		expectedResult{
			true,
			false,
			nil,
			"",
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
	{
		`Mozilla/5.0 (Linux; U; Android 4.0.4; ru-ru; Fly IQ440; Build/IML74K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
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
		`Mozilla/5.0 (Linux; Android 4.1.1; Nexus 7 Build/JRO03D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166  Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.2; Nexus 7 Build/JOP40C) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (Linux; Android 4.1.2; Nexus 7 Build/JZ054K) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Safari/535.19`,
		expectedResult{
			true,
			true,
			map[string]string{
				`Android`: `4.1.2`,
				`Chrome`:  `18.0.1025.166`,
			},
			``,
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
		`Mozilla/5.0 (Linux; Android 4.3; Nexus 10 Build/JWR66Y) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.72 Safari/537.36`,
		expectedResult{
			true,
			true,
			nil,
			``,
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
		`Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; U1 Build/HCL ME Tablet U1) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30`,
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
	/**

	*
	 */
}

type uaListResult struct {
	success bool
	message string
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
			}
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
					t.Errorf("expected version: %s, actual version: %s", v, actualVersion)
				}
			}

			chn <- result
		}(idx, test.userAgent, test.er, chn)
	}
	for i := 0; i < len(uaListTests); i++ {
		result := <-chn
		if false == result.success {
			t.Error(result.message)
		}
		if result.success && "done" == result.message {
			break
		}
	}
}

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
		`Mozilla/5.0 (Linux; U; Android 2.1-update1; de-de; HTC Desire 1.19.161.5 Build/ERE27) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Mobile Safari/530.17`,
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
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; Radar C110e)`,
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
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; T8788)`,
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
		`Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0; HTC; TITAN X310e)`,
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
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0)`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
	{
		`Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)`,
		expectedResult{
			true,
			true,
			nil,
			``,
		},
	},
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
		`Mozilla/5.0 (SymbianOS/9.3; Series60/3.2 NokiaE52-1/@version@; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.26 Mobile Safari/533.4 3gpp-gba`,
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
	/**


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

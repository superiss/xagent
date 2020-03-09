package xagent

// Browsers
const (
	Chrome         = "chrome"
	Safari         = "safari"
	Firefox        = "firefox"
	Opera          = "opera"
	Edge           = "edge"
	Netscape       = "netscape"
	WebBrowser     = "web browser"
	MSOffice       = "microsoft browser"
	RuxitSynthetic = "ruxit synthetic"
	Samsung        = "samsung browser"
	UCbrowser      = "us browser"
	Yandex         = "yandex"
	AndroidBrowser = "android browser"
)

// OS
const (
	Windows      = "windows"
	WindowsPhone = "windows_phone"
	Android      = "android"
	MacOS        = "macintosh"
	ChromeOS     = "chromeos"
	IOS          = "ios"
	Blackberry   = "blackberry"
	Symbian      = "symbian"
	Linux        = "linux"
)

var (
	// os
	windowsos    []string
	macos        []string
	ios          []string
	androids     []string
	linuxos      []string
	chromeos     []string
	blackberryos []string
	symbianos    []string
	// devices
	computers []string
	mobiles   []string
	tablets   []string
	bots      []string
	// browsers
	chrome         []string
	safari         []string
	edge           []string
	firefox        []string
	webbrowser     []string
	msoffice       []string
	operabrowser   []string
	ruxitsynthetic []string
	samsung        []string
	ucbrowser      []string
	yandex         []string
	netscape       []string
)

// Unknown is returned when a result cannot be extracted
const (
	Computer = "computer"
	Mobile   = "mobile"
	Tablet   = "tablet"
	Bot      = "web bot"
	Unknown  = "unknown"
)

func init() {
	// OS
	windowsos = []string{
		"windows",
		"win64",
		"wow64",
	}
	macos = []string{
		"macintosh",
		"darwin/",
		"mac os x",
		"mac_powerpc",
		"itunes",
	}
	linuxos = []string{
		"linux",
		"android",
	}
	chromeos = []string{
		"cros",
		// "crios",
	}
	blackberryos = []string{
		"blackberry",
	}
	symbianos = []string{
		"symbian",
		"symbos",
		"(series",
		"ideatab",
	}
	// // // //
	// // // //
	// device
	computers = []string{
		"(windows;",
		"(macintosh;",
		"(windows nt",
		"ubuntu",
	}
	mobiles = []string{
		"android",
		"mobile",
		"blackberry",
		"avantgo",
		"blazer",
		"iris",
		"iphone",
	}
	tablets = []string{
		"ipad",
		"vivo",
		"tablet",
		"playbook",
	}
	bots = []string{
		"bot",
		"proxy",
		"slurp",
		"spider",
		"bing",
		"google",
		"crawl",
	}

	// // // // // //
	// // // // // //
	// browsers

	chrome = []string{
		"chrome",
		"chromium",
		"crios",
	}
	//
	safari = []string{
		"safari/",
	}
	//
	edge = []string{
		"edge/",
		"EdgA/",
		"edg/",
	}
	//
	firefox = []string{
		"firefox/",
	}
	//
	webbrowser = []string{
		"webbrowser",
		"web browser",
		"webview",
		"iemb",
	}
	msoffice = []string{
		"ms-office",
		"outlook-ios",
		"microsoft",
		"outlook",
		"msoffice",
		"msie",
		"trident",
	}
	operabrowser = []string{
		"opt/",
		"opera",
		"presto/",
		"opr/",
		"oprgx",
	}
	ruxitsynthetic = []string{
		"ruxitsynthetic",
	}
	samsung = []string{
		"samsung",
	}
	ucbrowser = []string{
		"ucbrowser",
		"uc browser",
		"uc web",
	}
	yandex = []string{
		"yabrowser",
		"yowser",
	}

	netscape = []string{
		"netscape",
	}
}

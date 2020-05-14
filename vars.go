package xagent

// Devices
const (
	deviceComputer     = "computer"
	deviceTablet       = "tablet"
	deviceMobile       = "mobile"
	deviceGameConsoles = "game consoles"
	deviceBot          = "bot"
)

// OS
const (
	osWindows      = "windows"
	osMacOS        = "macos"
	osLinux        = "linux"
	osChromeOS     = "chromeos"
	osWindowsPhone = "windowsphone"
	osAndroid      = "android"
	osIOS          = "ios"
	osBlackberry   = "blackberry"
	osSymbian      = "symbian"
)

// Browsers
const (
	browserChrome         = "chrome"
	browserSafari         = "safari"
	browserFirefox        = "firefox"
	browserOpera          = "opera"
	browserEdge           = "edge"
	browserWebBrowser     = "web browser"
	browserMSOffice       = "internet explorer"
	browserAndroidBrowser = "android browser"
	browserSamsung        = "samsung browser"
	browserUCbrowser      = "us browser"
	browserYandex         = "yandex"
)

// Unknown if user agent is not detected
const unknown = "unknown"

var (
	// os
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
		"crios",
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
		"pixel c build",
		"sm-t",
		"kfthwi",
		"lg-v",
	}
	// bots
	bots = []string{
		"bot",
		"proxy",
		"slurp",
		"spider",
		"bing",
		"google",
		"crawl",
	}
	consoles = []string{
		"nintendo",
		"playstation",
		"xbox",
	}

	// browsers
	chrome = []string{
		"chrome",
		"chromium",
		"crios",
		"crmo",
	}
	safari = []string{
		"safari/",
	}
	edge = []string{
		"edge/",
		"EdgA/",
		"edg/",
	}
	firefox = []string{
		"firefox/",
		"fxios",
	}
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
	ucbrowser = []string{
		"ucbrowser",
		"uc browser",
		"uc web",
	}
	yandex = []string{
		"yabrowser",
		"yowser",
	}
)

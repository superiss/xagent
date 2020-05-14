package xagent

import "strings"

// More at: https://deviceatlas.com/blog/list-of-user-agent-strings

// UserAgent structure
type UserAgent struct {
	Name       string
	DeviceType string
	OSType     string
}

// Parser useragent
func Parser(ua string) *UserAgent {
	// ua = toLower(ua)
	ua = strings.ToLower(ua)
	//
	device := device(ua)
	os := os(ua)
	browser := browser(ua)
	//
	switch {
	case os == osChromeOS && browser == browserChrome:
		device = deviceComputer
	case os == osWindows && browser == browserMSOffice:
		device = deviceComputer
		//
	case os == osMacOS:
		if device == deviceMobile || device == deviceTablet {
			os = osIOS
		}
		//
	case os == osWindows:
		if device == deviceMobile || device == deviceTablet {
			os = osWindowsPhone
		}
		//
	case os == osLinux:
		if device == deviceMobile || device == deviceTablet {
			os = osAndroid
		}
	}
	//
	return &UserAgent{
		Name:       browser,
		DeviceType: device,
		OSType:     os,
	}
}

// Device type
func device(ua string) (device string) {
	// check if bot
	for _, bot := range bots {
		if strings.Contains(ua, bot) {
			device = deviceBot
			return
		}
	}
	// check if computer
	for _, computer := range computers {
		if strings.Contains(ua, computer) {
			device = deviceComputer
		}
	}

	// check if is mobile
	for _, mobile := range mobiles {
		if strings.Contains(ua, mobile) {
			device = deviceMobile
		}
	}
	// if device is mobile but it is also a tablet
	if device == deviceMobile {
		// we scan it against for tablet
		for _, tablet := range tablets {
			if strings.Contains(ua, tablet) {
				device = deviceTablet
				return
			}
		}
	}
	// check if game console
	for _, console := range consoles {
		if strings.Contains(ua, console) {
			device = deviceGameConsoles
			return
		}
	}
	// if device is computor os but it is a mobile
	switch device {
	case deviceComputer:
		// we scan it against mobile
		for _, mobile := range mobiles {
			if strings.Contains(ua, mobile) {
				device = deviceMobile
				return
			}
		}
	case "":
		device = unknown
		return
	}
	//
	return device
}

// OS Operating System
func os(ua string) (os string) {
	// check if windows
	for _, y := range windowsos {
		if strings.Contains(ua, y) {
			os = osWindows
			return
		}
	}
	// check if MacOS
	for _, y := range macos {
		if strings.Contains(ua, y) {
			os = osMacOS
			return
		}
	}
	// check if linux
	for _, y := range linuxos {
		if strings.Contains(ua, y) {
			os = osLinux
			return
		}
	}
	// check if ChromeOS
	for _, y := range chromeos {
		if strings.Contains(ua, y) {
			os = osChromeOS
			return
		}
	}
	// check if BlackBerry
	for _, y := range blackberryos {
		if strings.Contains(ua, y) {
			os = osBlackberry
			return
		}
	}
	// check if Symbian
	for _, y := range symbianos {
		if strings.Contains(ua, y) {
			os = osSymbian
			return
		}
	}
	// clean up
	if os == "" {
		os = unknown
		return
	}
	//
	return os
}

// browser name
func browser(ua string) (browser string) {
	// Opera
	for _, y := range operabrowser {
		if strings.Contains(ua, y) {
			browser = browserOpera
			return
		}
	}
	// Yandex
	for _, y := range yandex {
		if strings.Contains(ua, y) {
			browser = browserYandex
			return
		}
	}
	// ucbrowser
	for _, y := range ucbrowser {
		if strings.Contains(ua, y) {
			browser = browserUCbrowser
			return
		}
	}
	// web browser
	for _, y := range webbrowser {
		if strings.Contains(ua, y) {
			browser = browserWebBrowser
			return
		}
	}
	// safari
	if strings.Contains(ua, "safari") {
		browser = browserSafari
	}
	// chrome
	for _, y := range chrome {
		if strings.Contains(ua, y) && strings.Contains(ua, "safari") {
			browser = browserChrome
		}
	}
	// Edge
	for _, y := range edge {
		if strings.Contains(ua, y) {
			browser = browserEdge
			return
		}
	}
	// MS Office
	for _, y := range msoffice {
		if strings.Contains(ua, y) {
			browser = browserMSOffice
			return
		}
	}
	// firefox
	for _, y := range firefox {
		if strings.Contains(ua, y) {
			browser = browserFirefox
		}
	}
	//
	if browser == "" {
		browser = unknown
		return
	}
	//
	return browser
}

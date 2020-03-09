package xagent

/*
Testing Mode:
*/

import "strings"

// UserAgent structure
type UserAgent struct {
	Name       string
	DeviceType string
	OSType     string
	Original   string
}

// New UserAgent define
func New(s string) *UserAgent {
	// s = strings.ToLower(s)
	// s = ToLower(s)
	return &UserAgent{
		Original: s,
	}
}

// Detect useragent
func (ua *UserAgent) Detect() *UserAgent {
	ua.device()
	ua.os()
	ua.browser()
	//
	return ua
}

// Device type
func (ua *UserAgent) device() {
	// check if computer
	for _, y := range computers {
		if strings.Contains(ua.Original, y) {
			ua.DeviceType = Computer
		}
	}
	// check if bot
	for _, y := range bots {
		if strings.Contains(ua.Original, y) {
			ua.DeviceType = Bot
		}
	}
	// check if is mobile
	for _, y := range mobiles {
		if strings.Contains(ua.Original, y) {
			ua.DeviceType = Mobile
		}
	}
	// if device is mobile but it is also a tablet
	if ua.DeviceType == Mobile {
		// we scan it against for tablet
		for _, y := range tablets {
			if strings.Contains(ua.Original, y) {
				ua.DeviceType = Tablet
			}
		}
	}
	// if device is computor os but it is a mobile
	if ua.DeviceType == Computer {
		// we scan it against mobile
		for _, y := range mobiles {
			if strings.Contains(ua.Original, y) {
				ua.DeviceType = Mobile
			}
		}
	}

	if ua.DeviceType == "" {
		ua.DeviceType = Unknown
	}
}

// OS Operating System
func (ua *UserAgent) os() {
	// check if ChromeOS
	for _, y := range chromeos {
		if strings.Contains(ua.Original, y) {
			ua.OSType = ChromeOS
		}
	}

	// check if BlackBerry
	for _, y := range blackberryos {
		if strings.Contains(ua.Original, y) {
			ua.OSType = Blackberry
		}
	}

	// check if Symbian
	for _, y := range symbianos {
		if strings.Contains(ua.Original, y) {
			ua.OSType = Symbian
		}
	}
	// check if windows
	for _, y := range windowsos {
		if strings.Contains(ua.Original, y) {
			ua.OSType = Windows
		}
	}

	// check if linux
	for _, y := range linuxos {
		if strings.Contains(ua.Original, y) {
			ua.OSType = Linux
		}
	}

	// check if MacOS
	for _, y := range macos {
		if strings.Contains(ua.Original, y) {
			ua.OSType = MacOS
		}
	}

	// clean up
	switch {
	case ua.OSType == MacOS:
		if ua.DeviceType == Mobile {
			ua.OSType = IOS
		}
	case ua.OSType == Windows:
		if ua.DeviceType == Mobile || ua.DeviceType == Tablet {
			ua.OSType = WindowsPhone
		}
	case ua.OSType == Linux:
		if ua.DeviceType == Mobile || ua.DeviceType == Tablet {
			ua.OSType = Android
		}
	case ua.OSType == "":
		ua.OSType = Unknown
	}
}

// browser name
func (ua *UserAgent) browser() {
	// // // // // // //
	// safari
	if strings.Contains(ua.Original, "safari") {
		ua.Name = Safari
	}
	// chrome
	for _, y := range chrome {
		if strings.Contains(ua.Original, y) && strings.Contains(ua.Original, "safari") {
			ua.Name = Chrome
		}
	}
	// // // // // // // //
	// // // // // // // //
	// firefox
	for _, y := range firefox {
		if strings.Contains(ua.Original, y) {
			ua.Name = Firefox
		}
	}
	// Edge
	for _, y := range edge {
		if strings.Contains(ua.Original, y) {
			ua.Name = Edge
		}
	}
	// Opera
	for _, y := range operabrowser {
		if strings.Contains(ua.Original, y) {
			ua.Name = Opera
		}
	}
	// MS Office
	for _, y := range msoffice {
		if strings.Contains(ua.Original, y) {
			ua.Name = MSOffice
		}
	}
	// Netscape
	for _, y := range netscape {
		if strings.Contains(ua.Original, y) {
			ua.Name = Netscape
		}
	}

	// Yandex
	for _, y := range yandex {
		if strings.Contains(ua.Original, y) {
			ua.Name = Yandex
		}
	}

	// ucbrowser
	for _, y := range ucbrowser {
		if strings.Contains(ua.Original, y) {
			ua.Name = UCbrowser
		}
	}

	// samsung
	for _, y := range samsung {
		if strings.Contains(ua.Original, y) {
			ua.Name = Samsung
		}
	}

	// ruxitsynthetic
	for _, y := range ruxitsynthetic {
		if strings.Contains(ua.Original, y) {
			ua.Name = RuxitSynthetic
		}
	}

	// web browser
	for _, y := range webbrowser {
		if strings.Contains(ua.Original, y) {
			ua.Name = WebBrowser
		}
	}

	if ua.Name == "" {
		ua.Name = Unknown
	}
}

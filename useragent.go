package useragent

import (
	"fmt"
	"strings"
)

// Parse gets the raw user agent and returns a usable UserAgent
func Parse(input string) *UserAgent {

	var (
		ua UserAgent
	)

	ua.parse(strings.ToLower(input))
	return &ua

}

func (u *UserAgent) parse(input string) {
	u.partsMap = partsToMap(input)
	u.input = input

	u.parseOS()
	u.parseBrowser()

}

func (u *UserAgent) parseOS() {
	switch {
	case u.includes("ipod"):
		u.deviceID = 3
		u.platformID = 7
		u.osID = 4
		u.osVersion = u.getVersionString("cpu iphone os ")
	case u.includes("iphone"):
		u.deviceID = 4
		u.platformID = 6
		u.osID = 4
		u.osVersion = u.getVersionString("cpu iphone os ")
	case u.includes("ipad"):
		u.deviceID = 3
		u.platformID = 5
		u.osID = 4
		u.osVersion = u.getVersionString("cpu os ")
	case u.includes("macintosh"):
		u.deviceID = 2
		u.platformID = 4
		u.osID = 3
		u.osVersion = u.getVersionString("intel mac os x ")
	case u.includes("android"):
		u.deviceID = 4
		u.platformID = 2
		u.osID = 7
		u.osVersion = u.getVersionString("android ")
	case u.includes("cros"):
		u.deviceID = 2
		u.platformID = 2
		u.osID = 9

		archs := []string{"x86_64", "armv7l", "armv6l", "aarch64", "i686"}
		for _, arch := range archs {
			if u.includes(arch) {
				u.osVersion = u.getVersionString(fmt.Sprintf("cros %s ", arch))
				break
			}
		}
	case u.includes("linux"), u.includes("debian"), u.includes("ubuntu"), u.includes("x11"):
		u.deviceID = 2
		u.platformID = 2
		u.osID = 2
	case u.includes("windows") && u.includes("nt"):
		u.deviceID = 2
		u.platformID = 3
		u.osID = 6
		u.osVersion = u.getVersionString("windows nt ")
	case u.includes("windows") && (u.includes("phone") || u.includes("mobile")):
		u.platformID = 9
		u.deviceID = 4
		u.osID = 5
	case u.includes("blackberry"), u.includes("bb10"), u.includes("playbook"):
		u.deviceID = 4
		u.platformID = 8
		u.osID = 8
	}

}

func (u *UserAgent) parseBrowser() {

	// every OS have different posibilities
	switch u.osID {
	case 0: // unknown

	case 1: // Bot

	case 2:
		u.parseLinux()
	case 3:
		u.parseMacOs()
	case 4:
		u.parseiOs()
	case 5: // windows phone

	case 6:
		u.parseWindows()
	case 7:
		u.parseAndroid()
	case 8: // blackberry
		u.browserID = 11
		if u.includes("version") {
			u.browserVersion = u.partsMap["version"]
		}
	case 9:
		// ChromeOs only has chrome, isnt?
		if u.includes("chrome") {
			u.browserVersion = u.partsMap["chrome"]
			u.browserID = 1
		}
	}

}

func (u *UserAgent) parseLinux() {
	switch {
	case u.includes("opr"):
		u.browserVersion = u.partsMap["opr"]
		u.browserID = 6
	case u.includes("vivaldi"):
		u.browserVersion = u.partsMap["vivaldi"]
		u.browserID = 7
	case u.includes("chrome"):
		u.browserVersion = u.partsMap["chrome"]
		u.browserID = 1
	case u.includes("firefox"):
		u.browserVersion = u.partsMap["firefox"]
		u.browserID = 3
		// case u.includes("edg"):
		// 	u.browserVersion = u.partsMap["edg"]
		// 	u.browserID = 8
	}
}

func (u *UserAgent) parseiOs() {
	switch {
	case u.includes("edgios"):
		u.browserVersion = u.partsMap["edgios"]
		u.browserID = 8
	case u.includes("safari") && u.includes("version"):
		u.browserVersion = u.partsMap["version"]
		u.browserID = 3
	case u.includes("crios"):
		u.browserVersion = u.partsMap["crios"]
		u.browserID = 1
	case u.includes("fxios"):
		u.browserVersion = u.partsMap["fxios"]
		u.browserID = 3
	case u.includes("opt"):
		u.browserVersion = u.partsMap["opt"]
		u.browserID = 6
	case u.includes("applewebkit"):
		u.browserID = 9
	}
}

func (u *UserAgent) parseMacOs() {
	switch {
	case u.includes("safari") && u.includes("version"):
		u.browserVersion = u.partsMap["version"]
		u.browserID = 3
	case u.includes("opr"):
		u.browserVersion = u.partsMap["opr"]
		u.browserID = 6
	case u.includes("vivaldi"):
		u.browserVersion = u.partsMap["vivaldi"]
		u.browserID = 7
	case u.includes("edg"):
		u.browserVersion = u.partsMap["edg"]
		u.browserID = 8
	case u.includes("brave"):
		// brave does not show its own version
		// some new versions even does not show itself as Brave, so it will treated as Chrome
		u.browserVersion = ""
		u.browserID = 10
	case u.includes("chrome"):
		u.browserVersion = u.partsMap["chrome"]
		u.browserID = 1
	case u.includes("firefox"):
		u.browserVersion = u.partsMap["firefox"]
		u.browserID = 3
	case u.includes("applewebkit"):
		u.browserID = 9
	}
}

func (u *UserAgent) parseWindows() {
	switch {
	case u.includes("msie"):
		u.browserVersion = u.getVersionString("msie ")
		u.browserID = 2
	case u.includes("edge"):
		u.browserVersion = u.partsMap["edge"]
		u.browserID = 8
	case u.includes("edg"):
		u.browserVersion = u.partsMap["edg"]
		u.browserID = 8
	case u.includes("opr"):
		u.browserVersion = u.partsMap["opr"]
		u.browserID = 6
	case u.includes("vivaldi"):
		u.browserVersion = u.partsMap["vivaldi"]
		u.browserID = 7
	case u.includes("brave"):
		// brave does not show its own version
		// some new versions even does not show itself as Brave, so it will treated as Chrome
		u.browserVersion = ""
		u.browserID = 10
	case u.includes("chrome"):
		u.browserVersion = u.partsMap["chrome"]
		u.browserID = 1
	case u.includes("firefox"):
		u.browserVersion = u.partsMap["firefox"]
		u.browserID = 3
	}
}

func (u *UserAgent) parseAndroid() {
	switch {
	case u.includes("firefox"):
		u.browserVersion = u.partsMap["firefox"]
		u.browserID = 3
	case u.includes("opr"):
		u.browserVersion = u.partsMap["opr"]
		u.browserID = 6
	case u.includes("opt"):
		u.browserVersion = u.partsMap["opt"]
		u.browserID = 6
	// case u.includes("vivaldi"):
	// 	u.browserVersion = u.partsMap["vivaldi"]
	// 	u.browserID = 7
	case u.includes("edga"):
		u.browserVersion = u.partsMap["edga"]
		u.browserID = 8
	case u.includes("brave"):
		u.browserVersion = ""
		u.browserID = 10
	case u.includes("chrome"):
		u.browserVersion = u.partsMap["chrome"]
		u.browserID = 1
	case u.includes("applewebkit"):
		u.browserID = 9
	}
}

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
	u.parseBot() // is a bot?

	if u.browserID == 0 {
		// some corner cases like spotify not setting the OS always
		u.parseOtherBrowsers()
	}
}

func (u *UserAgent) parseOS() bool {
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
	case u.includes("webos") || u.includes("webostv"):
		u.deviceID = 6
		u.platformID = 10
		u.osID = 11
	case u.includes("linux"), u.includes("debian"), u.includes("ubuntu"), u.includes("x11"):
		u.deviceID = 2
		u.platformID = 2
		u.osID = 2
	case u.includes("xbox"):
		// xbox announces itself also as windows nt, so takes preference
		u.deviceID = 7
		u.platformID = 10
		u.osID = 13
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
	case u.includes("symbianos"):
		u.deviceID = 4
		u.platformID = 8
		u.osID = 10
		u.osVersion = u.partsMap["symbianos"]
	case u.includes("symbian"):
		u.deviceID = 4
		u.platformID = 8
		u.osID = 10
		u.osVersion = u.partsMap["symbian"]
	case u.includes("playstation") && !u.includes("psp"):
		u.deviceID = 7
		u.platformID = 10
		u.osID = 12
	case u.includes("psp"):
		u.deviceID = 8
		u.platformID = 10
		u.osID = 12
	case u.includes("nintendo"):
		u.deviceID = 7
		u.platformID = 10
		u.osID = 14
	default:
		// os not detected
		return false
	}
	return true
}

func (u *UserAgent) parseBrowser() {

	// every OS have different posibilities
	switch u.osID {
	case 2:
		u.parseLinux()
	case 3:
		u.parseMacOs()
	case 4:
		u.parseiOs()
	case 5: // windows phone
		// TODO:!
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
	case 10:
		switch {
		case u.includes("nokiabrowser"):
			u.browserVersion = u.partsMap["nokiabrowser"]
			u.browserID = 14
		}
	case 11:
		// webos
		switch {
		case u.includes("chrome"):
			u.browserVersion = u.partsMap["chrome"]
			u.browserID = 1
		}
	case 12:
		// playstation only hast netfront browser
		u.browserID = 15
	case 13:
		// xbox uses only edge?
		switch {
		case u.includes("edge"):
			u.browserVersion = u.partsMap["edge"]
			u.browserID = 8
		}
	case 14:
		// nintendo uses nintendo browser
		u.browserID = 16
		if u.includes("nintendobrowser") {
			u.browserVersion = u.partsMap["nintendobrowser"]
		}
	}

}

func (u *UserAgent) parseLinux() {
	switch {
	case u.includes("maxthon"):
		u.browserVersion = u.partsMap["maxthon"]
		u.browserID = 18
	case u.includes("opr"):
		u.browserVersion = u.partsMap["opr"]
		u.browserID = 6
	case u.includes("vivaldi"):
		u.browserVersion = u.partsMap["vivaldi"]
		u.browserID = 7
	case u.includes("silk"):
		u.browserVersion = u.getVersionString("silk")
		u.browserID = 13
	case u.includes("samsungbrowser"):
		u.browserVersion = u.partsMap["samsungbrowser"]
		u.browserID = 20
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
	case u.includes("maxthon"):
		u.browserVersion = u.partsMap["maxthon"]
		u.browserID = 18
	case u.includes("edgios"):
		u.browserVersion = u.partsMap["edgios"]
		u.browserID = 8
	case u.includes("mqqbrowser"):
		// qqbrowser sets safari and version also
		u.browserVersion = u.partsMap["mqqbrowser"]
		u.browserID = 17
	case u.includes("yabrowser"):
		u.browserVersion = u.partsMap["yabrowser"]
		u.browserID = 21
	case u.includes("coc_coc_browser"):
		u.browserVersion = u.partsMap["coc_coc_browser"]
		u.browserID = 22
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
	case u.includes("maxthon"):
		u.browserVersion = u.partsMap["maxthon"]
		u.browserID = 18
	case u.includes("silk"):
		u.browserVersion = u.getVersionString("silk")
		u.browserID = 13
	case u.includes("coc_coc_browser"):
		u.browserVersion = u.partsMap["coc_coc_browser"]
		u.browserID = 22
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
	case u.includes("qqbrowser"):
		u.browserVersion = u.partsMap["qqbrowser"]
		u.browserID = 17
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
	case u.includes("maxthon"):
		u.browserVersion = u.partsMap["maxthon"]
		u.browserID = 18
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
	case u.includes("qqbrowser"):
		u.browserVersion = u.partsMap["qqbrowser"]
		u.browserID = 17
	case u.includes("yabrowser"):
		u.browserVersion = u.partsMap["yabrowser"]
		u.browserID = 21
	case u.includes("coc_coc_browser"):
		u.browserVersion = u.partsMap["coc_coc_browser"]
		u.browserID = 22
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
	case u.includes("baidu.sogo.uc.ucbrowser"):
		u.browserVersion = u.partsMap["baidu.sogo.uc.ucbrowser"]
		u.browserID = 12
	case u.includes("ucbrowser"):
		u.browserVersion = u.partsMap["ucbrowser"]
		u.browserID = 12
	case u.includes("ucmini"):
		u.browserVersion = u.partsMap["ucmini"]
		u.browserID = 12
	case u.includes("silk"):
		u.browserVersion = u.partsMap["silk"]
		u.browserID = 13
	case u.includes("nokiabrowser"):
		u.browserVersion = u.partsMap["nokiabrowser"]
		u.browserID = 14
	case u.includes("mqqbrowser"):
		u.browserVersion = u.partsMap["mqqbrowser"]
		u.browserID = 17
	case u.includes("samsungbrowser"):
		u.browserVersion = u.partsMap["samsungbrowser"]
		u.browserID = 20
	case u.includes("coc_coc_browser"):
		u.browserVersion = u.partsMap["coc_coc_browser"]
		u.browserID = 22
	case u.includes("chrome"):
		u.browserVersion = u.partsMap["chrome"]
		u.browserID = 1
	case u.includes("applewebkit"):
		u.browserID = 9
	}
}

func (u *UserAgent) parseOtherBrowsers() {
	switch {
	case u.includes("spotify"):
		u.browserVersion = u.partsMap["spotify"]
		u.browserID = 19

		switch {
		case u.includes("android"):
			u.deviceID = 4
			u.platformID = 2
			u.osID = 7
		default:
			if strings.Contains(u.input, "iphone") {
				u.deviceID = 4
				u.platformID = 6
				u.osID = 4
				if u.includes("ios") {
					u.osVersion = u.partsMap["ios"]
				}
			}
		}

	}
}

func (u *UserAgent) parseBot() bool {
	switch {
	case u.includes("applebot"):
		u.botID = 2
		u.botVersion = u.partsMap["applebot"]
		u.bot = true
	case u.includes("baiduspider"):
		u.botID = 3
		u.botVersion = u.partsMap["baiduspider"]
		u.bot = true
	case u.includes("baiduspider+"):
		u.botID = 3
		u.bot = true
		u.botVersion = u.partsMap["baiduspider+"]
	case u.includes("baiduspider-render"):
		u.botID = 3
		u.bot = true
		u.botVersion = u.partsMap["baiduspider-render"]
	}

	if u.bot && u.deviceID == 0 {
		u.deviceID = 1
	}
	if u.bot && u.platformID == 0 {
		u.platformID = 1
	}
	if u.bot && u.osID == 0 {
		u.osID = 1
	}

	return u.bot
}

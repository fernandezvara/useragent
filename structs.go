package useragent

type UserAgent struct {
	deviceID       int
	platformID     int
	browserID      int
	browserVersion string
	osID           int
	osVersion      string
	bot            bool
	partsMap       map[string]string
	input          string
}

func (u UserAgent) Device() device {
	return device{id: u.deviceID}
}

func (u UserAgent) Platform() platform {
	return platform{id: u.platformID}
}

func (u UserAgent) Browser() browser {
	return browser{id: u.browserID, version: u.browserVersion, bot: u.bot}
}

func (u UserAgent) OS() os {
	return os{id: u.osID, version: u.osVersion}
}

func (u UserAgent) Mobile() bool {
	return u.deviceID == 3 || u.deviceID == 4 || u.deviceID == 5
}

func (u UserAgent) Bot() bool {
	return u.bot
}

var Devices = []string{"Unknown", "Bot", "Computer", "Tablet", "Phone", "Wearable", "TV", "Console"}
var Platforms = []string{"Unknown", "Bot", "Linux", "Windows", "Mac", "iPad", "iPhone", "iPod", "Blackberry", "WindowsPhone", "Playstation", "Xbox", "Nintendo"}
var Browsers = []string{"Unknown", "Chrome", "IE", "Safari", "Firefox", "Android", "Opera", "Vivaldi", "Edge", "Webkit Based Browser", "Brave", "Blackberry", "UC", "Silk", "Nokia", "NetFront", "QQ", "Maxthon", "SogouExplorer", "Spotify", "Nintendo", "Samsung", "Yandex", "CocCoc"}
var Bots = []string{"Bot", "AppleBot", "BaiduBot", "BingBot", "DuckDuckGoBot", "FacebookBot", "GoogleBot", "LinkedInBot", "MsnBot", "PingdomBot", "TwitterBot", "YandexBot", "CocCocBot", "YahooBot"}
var OSs = []string{"Unknown", "Bot", "Linux", "MacOSX", "iOS", "WindowsPhone", "Windows", "Android", "Blackberry", "ChromeOS", "Kindle", "WebOS", "Playstation", "Xbox", "Nintendo"}

type device struct {
	id int
}

func (d device) ID() int {
	return d.id
}

func (d device) String() string {
	return Devices[d.id]
}

type platform struct {
	id int
}

func (p platform) ID() int {
	return p.id
}

func (p platform) String() string {
	return Platforms[p.id]
}

type browser struct {
	id      int
	version string
	bot     bool
}

func (b browser) ID() int {
	return b.id
}

func (b browser) String() string {
	if b.bot {
		return Bots[b.id]
	}
	return Browsers[b.id]
}

func (b browser) Version() string {
	return b.version
}

type os struct {
	id      int
	version string
}

func (o os) ID() int {
	return o.id
}

func (o os) String() string {
	return OSs[o.id]
}

func (o os) Version() string {
	return o.version
}

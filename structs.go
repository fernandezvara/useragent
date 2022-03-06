package useragent

type UserAgent struct {
	deviceID       int
	platformID     int
	browserID      int
	browserVersion string
	botID          int
	botVersion     string
	osID           int
	osVersion      string
	bot            bool
	partsMap       map[string]string
	input          string
}

var Devices = []string{
	"Unknown",
	"Bot",
	"Computer",
	"Tablet",
	"Phone",
	"Wearable",
	"TV",
	"Console",
	"Portable Console",
}

var Platforms = []string{
	"Unknown",
	"Bot",
	"Linux",
	"Windows",
	"Mac",
	"iPad",
	"iPhone",
	"iPod",
	"Blackberry",
	"WindowsPhone",
	"Playstation",
	"Xbox",
	"Nintendo",
	"TV",
}

var Browsers = []string{
	"Unknown",              // 0
	"Chrome",               // 1
	"IE",                   // 2
	"Safari",               // 3
	"Firefox",              // 4
	"Android",              // 5
	"Opera",                // 6
	"Vivaldi",              // 7
	"Edge",                 // 8
	"Webkit Based Browser", // 9
	"Brave",                // 10
	"Blackberry",           // 11
	"UC",                   // 12
	"Silk",                 // 13
	"NokiaBrowser",         // 14
	"NetFront",             // 15
	"Nintendo",             // 16
	"QQ",                   // 17
	"Maxthon",              // 18
	"Spotify",              // 19
	"Samsung",              // 20
	"Yandex",               // 21
	"CocCoc",               // 22
}

var Bots = []string{
	"",              // 0
	"Bot",           // 1
	"AppleBot",      // 2
	"BaiduBot",      // 3
	"BingBot",       // 4
	"DuckDuckGoBot", // 5
	"FacebookBot",   // 6
	"GoogleBot",     // 7
	"LinkedInBot",   // 8
	"MsnBot",        // 9
	"PingdomBot",    // 10
	"TwitterBot",    // 11
	"YandexBot",     // 12
	"CocCocBot",     // 13
	"YahooBot",      // 14
}

var OSs = []string{
	"Unknown",
	"Bot",
	"Linux",
	"MacOSX",
	"iOS",
	"WindowsPhone",
	"Windows",
	"Android",
	"Blackberry",
	"ChromeOS",
	"Symbian",
	"WebOS",
	"Playstation",
	"Xbox",
	"Nintendo",
}

func (u UserAgent) Device() device {
	return device{id: u.deviceID}
}

func (u UserAgent) Platform() platform {
	return platform{id: u.platformID}
}

func (u UserAgent) Browser() browser {
	return browser{id: u.browserID, version: u.browserVersion, bot: u.bot, mobile: u.IsMobile()}
}

func (u UserAgent) Bot() bot {
	return bot{id: u.botID, version: u.botVersion, bot: u.bot}
}

func (u UserAgent) OS() os {
	return os{id: u.osID, version: u.osVersion}
}

func (u UserAgent) IsMobile() bool {
	return u.deviceID == 3 || u.deviceID == 4 || u.deviceID == 5 || u.deviceID == 8
}

func (u UserAgent) IsBot() bool {
	return u.bot
}

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
	mobile  bool
}

func (b browser) ID() int {
	return b.id
}

func (b browser) String() string {
	return Browsers[b.id]
}

func (b browser) Version() string {
	return b.version
}

func (b browser) IsBot() bool {
	return b.bot
}

func (b browser) IsMobile() bool {
	return b.mobile
}

type bot struct {
	id      int
	version string
	bot     bool
}

func (b bot) ID() int {
	return b.id
}

func (b bot) String() string {
	return Bots[b.id]
}

func (b bot) Version() string {
	return b.version
}

func (b bot) IsBot() bool {
	return b.bot
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

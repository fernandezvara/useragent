# User-Agent parser

[![Coverage Status](https://coveralls.io/repos/github/fernandezvara/useragent/badge.svg?branch=main)](https://coveralls.io/github/fernandezvara/useragent?branch=main)

This package parses the browser **User-Agent** received getting (with the best effort) this information (when possible):

- Platform
- Device
- OS name and version.
- Browser name and version.
- Bot name and version.

# Installation

```bash
go get github.com/fernandezvara/useragent
```

# Sample usage

### Parse **user agent** string.

```go
package main

import (
	"fmt"

	"github.com/fernandezvara/useragent"
)

func main() {

	example := useragent.Parse("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4280.88 Safari/537.36")

	fmt.Println(example.Device().ID())        // 2
	fmt.Println(example.Device().String())    // Computer
	fmt.Println(example.Platform().ID())      // 3
	fmt.Println(example.Platform().String())  // Windows
	fmt.Println(example.Browser().ID())       // 1
	fmt.Println(example.Browser().String())   // Chrome
	fmt.Println(example.Browser().Version())  // 99.0.4280.88
	fmt.Println(example.Browser().IsBot())    // false
	fmt.Println(example.Browser().IsMobile()) // false
	fmt.Println(example.Bot().ID())           // 0
	fmt.Println(example.Bot().String())       // ""
	fmt.Println(example.Bot().Version())      // ""
	fmt.Println(example.Bot().IsBot())        // false
	fmt.Println(example.OS().ID())            // 6
	fmt.Println(example.OS().String())        // Windows
	fmt.Println(example.OS().Version())       // 10.0
	fmt.Println(example.IsMobile())           // false
	fmt.Println(example.IsBot())              // false

}

```

### From IDs

```go
package main

import (
	"fmt"

	"github.com/fernandezvara/useragent"
)

func main() {

	var deviceID, platformID, browserID, osID, botID int = 4, 2, 1, 7, 0
	var browserVersion, osVersion, botVersion string = "99.0.4844.48", "10", ""

	example2 := useragent.ParseIDs(deviceID, platformID, browserID, osID, botID, browserVersion, osVersion, botVersion)

	fmt.Println(example2.Device().ID())        // 4
	fmt.Println(example2.Device().String())    // Phone
	fmt.Println(example2.Platform().ID())      // 2
	fmt.Println(example2.Platform().String())  // Linux
	fmt.Println(example2.Browser().ID())       // 1
	fmt.Println(example2.Browser().String())   // Chrome
	fmt.Println(example2.Browser().Version())  // 99.0.4844.48
	fmt.Println(example2.Browser().IsBot())    // false
	fmt.Println(example2.Browser().IsMobile()) // true
	fmt.Println(example2.Bot().ID())           // 0
	fmt.Println(example2.Bot().String())       // ""
	fmt.Println(example2.Bot().Version())      // ""
	fmt.Println(example2.Bot().IsBot())        // false
	fmt.Println(example2.OS().ID())            // 7
	fmt.Println(example2.OS().String())        // Android
	fmt.Println(example2.OS().Version())       // 10
	fmt.Println(example2.IsMobile())           // true
	fmt.Println(example2.IsBot())              // false

}
```

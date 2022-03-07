# useragent

User Agent parser

This package parses the browser **User Agent** received getting (with best effort) this information (when possible):

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

package main

import "fmt"

type Page struct {
	Url      string
	LoadTime int
}

var mockResults = map[string][]Page{
	"https://golang.org/": {
		Page{"https://golang.org/pkg/", 1000},
		Page{"https://golang.org/cmd/", 900},
	},
	"https://golang.org/pkg/": {
		Page{"https://golang.org/", 500},
		Page{"https://golang.org/cmd/", 900},
		Page{"https://golang.org/pkg/fmt/", 1500},
		Page{"https://golang.org/pkg/os/", 1400},
		Page{"https://golang.org/pkg/strconv/", 2000},
		Page{"https://golang.org/pkg/crypto/", 1200},
		Page{"https://golang.org/pkg/image/", 1900},
	},
	"https://golang.org/pkg/fmt/": {
		Page{"https://golang.org/", 500},
		Page{"https://golang.org/pkg/", 1000},
	},
	"https://golang.org/pkg/os/": {
		Page{"https://golang.org/", 500},
		Page{"https://golang.org/pkg/", 1000},
	},
	"https://golang.org/pkg/image/": {
		Page{"https://golang.org/pkg/image/Alpha", 1400},
		Page{"https://golang.org/pkg/image/Alpha16", 1600},
		Page{"https://golang.org/pkg/image/CMYK", 1300},
		Page{"https://golang.org/pkg/image/Config", 1500},
	},
}

// depth tells us when to stop searching for new links
func Crawl(page Page, depth int) {
	if depth <= 0 {
		return
	}

}

func main() {
	fmt.Println("Hello world")
}

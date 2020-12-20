package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// For more info: https://pkg.go.dev/github.com/go-rod/rod/lib/launcher
	u := launcher.New().MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	page := browser.MustPage("http://example.com")

	fmt.Println(page.MustInfo().Title)
}

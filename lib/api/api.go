package main

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

func init() {
	launcher.NewBrowser().MustGet()
}

func GetPage(url string) *rod.Page {
	browser := rod.New().Timeout(time.Minute).MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser)

	page.MustNavigate("https://www.hltv.org/matches")

	return page
}

package api

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
)

func GetPage(url string) *rod.Page {
	browser := rod.New().Timeout(time.Minute).MustConnect()
	page := stealth.MustPage(browser)
	page.MustNavigate(url)

	return page
}

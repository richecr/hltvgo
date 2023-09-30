package main

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

func init() {
	launcher.NewBrowser().MustGet()
}

func main() {
	browser := rod.New().Timeout(time.Minute).MustConnect()
	defer browser.MustClose()

	// You can also use stealth.JS directly without rod
	fmt.Printf("js: %x\n\n", md5.Sum([]byte(stealth.JS)))

	page := stealth.MustPage(browser)

	page.MustNavigate("https://www.hltv.org/matches")

	printReport(page)
}

func printReport(page *rod.Page) {
	el := page.MustElement("div.upcomingMatchesAll")
	matches := el.MustElements("div.upcomingMatch")
	fmt.Println(matches)
	for _, row := range matches {
		fmt.Println(row.MustElement("a").MustText())
		fmt.Println("-------")
	}
}

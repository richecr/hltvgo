package main

import (
	"github.com/go-rod/rod"
)

func main() {
	page := GetPage("https://www.hltv.org/matches")
}

func printReport(page *rod.Page) {
	// el := page.MustElement("div.upcomingMatchesAll")
	// matches := el.MustElements("div.upcomingMatch")
	// fmt.Println(matches)
	// for _, row := range matches {
	// 	fmt.Println(row.MustElement("a").MustText())
	// 	fmt.Println("-------")
	// }
}

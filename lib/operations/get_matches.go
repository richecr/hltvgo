package operations

import (
	"strings"

	"github.com/go-rod/rod"
	"github.com/richecr/hltv-go/lib/api"
	"github.com/richecr/hltv-go/lib/entity"
)

const (
	matchesURL = "https://www.hltv.org/matches"
)

func GetMatches() ([]entity.Match, error) {
	page := api.GetPage(matchesURL)
	defer page.MustClose()

	el := page.MustElement("div.upcomingMatchesAll")
	divMatches := el.MustElements("div.upcomingMatch")

	matches := make(chan []entity.Match)
	go GetMatch(divMatches[:len(divMatches)/2], matches)
	go GetMatch(divMatches[len(divMatches)/2:], matches)

	ms1, ms2 := <-matches, <-matches
	return append(ms1, ms2...), nil
}

func GetMatch(divMatches rod.Elements, matches chan []entity.Match) {
	var partial []entity.Match
	for _, row := range divMatches {
		principal := row.MustElement("a.match.a-reset")
		matchInfoEmpty, _ := principal.Element(".matchInfoEmpty")
		if matchInfoEmpty == nil {
			tags := strings.Split(row.MustElement("a").MustText(), "\n")
			team1 := GetTeam(row, tags[2])
			team2 := GetTeam(row, tags[3])
			id := strings.Split(*principal.MustAttribute("href"), "/")[2]
			match := entity.NewMatch(id, tags[4], "", *team1, *team2)
			partial = append(partial, *match)
		}
	}
	matches <- partial
}

func GetTeam(row *rod.Element, name string) *entity.Team {
	id := row.MustAttribute("team1")
	return entity.NewTeam(*id, name)
}

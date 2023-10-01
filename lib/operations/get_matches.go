package operations

import (
	"strconv"
	"strings"
	"time"

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

	element := page.MustElement("div.upcomingMatchesAll")
	divMatches := append(
		element.MustElements("div.upcomingMatch"),
		page.MustElements("div.liveMatch-container")...,
	)

	matches := make(chan []entity.Match, len(divMatches))
	defer close(matches)

	go GetMatch(divMatches[:len(divMatches)/2], matches)
	go GetMatch(divMatches[len(divMatches)/2:], matches)

	return append(<-matches, <-matches...), nil
}

func GetMatch(divMatches rod.Elements, matches chan []entity.Match) {
	var partial []entity.Match
	for _, row := range divMatches {
		principal := row.MustElement("a.match.a-reset")
		matchInfoEmpty, _ := principal.Element(".matchInfoEmpty")
		if matchInfoEmpty == nil {
			tags := strings.Split(row.MustElement("a").MustText(), "\n")
			var team1_name, team2_name, event string = tags[2], tags[3], tags[4]
			live := false
			var date time.Time
			if tags[0] == "LIVE" {
				team2_name = tags[4]
				event = tags[6]
				live = true
			} else {
				date = ConvertStringUnixToDate(
					*row.MustElement(".matchTime").MustAttribute("data-unix"),
				)
			}
			team1 := GetTeam(row, team1_name)
			team2 := GetTeam(row, team2_name)
			id := strings.Split(*principal.MustAttribute("href"), "/")[2]
			match := entity.NewMatch(id, event, date, live, *team1, *team2)
			partial = append(partial, *match)
		}
	}
	matches <- partial
}

func GetTeam(row *rod.Element, name string) *entity.Team {
	id := row.MustAttribute("team1")
	return entity.NewTeam(*id, name)
}

func ConvertStringUnixToDate(dateUnix string) time.Time {
	i, err := strconv.ParseInt(dateUnix, 10, 64)
	if err != nil {
		panic(err)
	}

	return time.Unix(i, 0)
}

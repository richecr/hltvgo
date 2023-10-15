package operations

import (
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/richecr/hltvgo/lib/api"
	"github.com/richecr/hltvgo/lib/entity"
)

const (
	matchesURL = "https://www.hltv.org/matches"
)

func GetMatches() ([]entity.Match, error) {
	page := api.GetPage(matchesURL)
	defer page.MustClose()

	events := GetEvents(page)
	element := page.MustElement("div.upcomingMatchesAll")
	divMatches := append(
		element.MustElements("div.upcomingMatch"),
		page.MustElements("div.liveMatch-container")...,
	)

	matches := make(chan []entity.Match, len(divMatches))
	defer close(matches)

	go GetMatch(divMatches[:len(divMatches)/2], matches, events)
	go GetMatch(divMatches[len(divMatches)/2:], matches, events)

	return append(<-matches, <-matches...), nil
}

func GetMatch(divMatches rod.Elements, matches chan []entity.Match, events []entity.Event) {
	var partial []entity.Match
	for _, row := range divMatches {
		main := row.MustElement("a.match.a-reset")
		matchInfoEmpty, _ := main.Element(".matchInfoEmpty")
		if matchInfoEmpty == nil {
			tags := strings.Split(row.MustElement("a").MustText(), "\n")
			var team1_name, team2_name, event_name string = tags[2], tags[3], tags[4]
			live := false
			var date time.Time
			if tags[0] == "LIVE" {
				team2_name = tags[4]
				event_name = tags[6]
				live = true
			} else {
				date = ConvertStringUnixToDate(
					*row.MustElement(".matchTime").MustAttribute("data-unix"),
				)
			}
			team1 := GetTeam(row, team1_name)
			team2 := GetTeam(row, team2_name)
			event := FindEventByName(event_name, events)
			id := strings.Split(*main.MustAttribute("href"), "/")[2]
			match := entity.NewMatch(id, event, date, live, *team1, *team2)
			partial = append(partial, *match)
		}
	}
	matches <- partial
}

func GetEvents(page *rod.Page) []entity.Event {
	var events []entity.Event

	div := page.MustElement("div.filter-custom-content")
	tagsEvents := append(
		div.MustElements("a.filter-button-link.event-row"),
		div.MustElements("a.filter-button-link")...,
	)

	for _, tag := range tagsEvents {
		href_text := strings.Split(*tag.MustAttribute("href"), "=")
		id := href_text[len(href_text)-1]
		var el_name *rod.Element
		el_name, err := tag.Element("div.event-name")
		if err != nil {
			el_name = tag.MustElement("div.featured-event-tooltip-content")
		}
		events = append(events, *entity.NewEvent(id, el_name.MustText()))
	}
	return events
}

func GetTeam(row *rod.Element, name string) *entity.Team {
	id := row.MustAttribute("team1")
	return entity.NewTeam(*id, name)
}

func FindEventByName(name string, events []entity.Event) entity.Event {
	for _, event := range events {
		if event.Name == name {
			return event
		}
	}
	return *entity.NewEvent("0", "not found")
}

func ConvertStringUnixToDate(dateUnix string) time.Time {
	i, err := strconv.ParseInt(dateUnix, 10, 64)
	if err != nil {
		panic(err)
	}

	return time.Unix(i, 0)
}

package entity

import (
	"time"
)

type Match struct {
	Id       string
	Team1    Team
	Team2    Team
	Event    Event
	DateHour time.Time
	Live     bool
}

func NewMatch(id string, event Event, dateHour time.Time, live bool, team1, team2 Team) *Match {
	return &Match{
		Id:       id,
		Team1:    team1,
		Team2:    team2,
		Event:    event,
		DateHour: dateHour,
		Live:     live,
	}
}

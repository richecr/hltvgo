package entity

type Match struct {
	Id       string
	Team1    Team
	Team2    Team
	Event    string
	DateHour string
}

func NewMatch(id, event, dateHour string, team1, team2 Team) *Match {
	return &Match{
		Id:       id,
		Team1:    team1,
		Team2:    team2,
		Event:    event,
		DateHour: dateHour,
	}
}

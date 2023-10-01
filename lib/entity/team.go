package entity

type Team struct {
	Id   string
	Name string
}

func NewTeam(id, name string) *Team {
	return &Team{
		Id:   id,
		Name: name,
	}
}

package entity

type Event struct {
	Id   string
	Name string
}

func NewEvent(id, name string) *Event {
	return &Event{
		Id:   id,
		Name: name,
	}
}

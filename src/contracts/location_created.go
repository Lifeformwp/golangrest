package contracts

import "github.com/user/rest-api/lib/persistence"

type LocationCreatedEvent struct {
	ID		string				`json:"id"`
	Name	string				`json:"name"`
	Address	string				`json:"address"`
	Country	string				`json:"country"`
	Halls	[]persistence.Hall	`json:"halls"`
}

func (c *LocationCreatedEvent) EventName() string {
	return "locationCreated"
}

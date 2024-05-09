package domain

type CardKind int
type CardStatus int

const (
	CardKindCorporate CardKind = iota
	CardKindPersonal
)

const (
	CardStatusOpen CardStatus = iota
	CardtStatusClosed
)

func (d CardKind) String() string {
	return [...]string{"CORPORATE", "PERSONAL"}[d]
}

func (d CardStatus) String() string {
	return [...]string{"OPEN", "CLOSED"}[d]
}

type Card struct {
	RefId   int        `json:"refId"`
	Limit   float64    `json:"limit"`
	Balance float64    `json:"balance"`
	Kind    CardKind   `json:"kind"`
	Status  CardStatus `json:"status"`
}

type CardRepository interface {
	SaveCard(card Card) error
	FindAll() ([]*Card, error)
}

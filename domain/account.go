package domain

type AccountKind int
type AccountStatus int

const (
	AccountKindGold AccountKind = iota
	AccountKindPlatinum
	AccountKindSilver
)

const (
	AccountStatusOpen AccountStatus = iota
	AccountStatusClosed
)

func (d AccountKind) String() string {
	return [...]string{"GOLD", "PLATINUM", "SILVER"}[d]
}

func (d AccountStatus) String() string {
	return [...]string{"OPEN", "CLOSED"}[d]
}

type Account struct {
	RefId  int           `json:"refId"`
	Limit  float64       `json:"limit"`
	Kind   AccountKind   `json:"kind"`
	Status AccountStatus `json:"status"`
}

type AccountRepository interface {
	SaveAccount(account Account) error
	FindAllAccounts() ([]*Account, error)
	Debit()
}

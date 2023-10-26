package internal_types

import "github.com/shopspring/decimal"

type Client struct {
	Id                 int             `json:"id"`
	Email              string          `json:"email" db:"email"`
	Phone              string          `json:"phone" db:"phone"`
	Address            string          `json:"address" db:"address"`
	BankAccount        string          `json:"bankAccount" db:"bank_account"`
	InvestmentAmount   decimal.Decimal `json:"investmentAmount" db:"investment_amount"`
	InvestmentStrategy string          `json:"investmentStrategy" db:"investment_strategy"`
}
type Clients []Client

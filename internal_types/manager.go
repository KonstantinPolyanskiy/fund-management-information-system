package internal_types

import "github.com/shopspring/decimal"

type Manager struct {
	Email            string          `json:"email" db:"email"`
	Phone            string          `json:"phone" db:"phone"`
	Address          string          `json:"address" db:"address"`
	BankAccount      string          `json:"bankAccount" db:"bank_account"`
	CapitalManagment decimal.Decimal `json:"capitalManagment" db:"capital_managment"`
	ProfitPercentDay float64         `json:"profitPercentDay" db:"profit_percent_day"`
}

type Managers []Manager

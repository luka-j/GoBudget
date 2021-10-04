package Entities

import (
	"GoBudget/Database"
	"database/sql"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	Database.Model
	payee string
	note  sql.NullString
}

type TransactionEnvelope struct {
	Database.Model
	Transaction Transaction
	Envelope    Envelope
	Amount      decimal.Decimal `gorm:"type:decimal"`
}

package Entities

import (
	"GoBudget/Database"
	"database/sql"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	Database.Model
	Payee         string
	BankAccountID uuid.UUID
	Note          sql.NullString
}

type TransactionEnvelope struct {
	Database.Model
	Transaction Transaction
	Envelope    Envelope
	Amount      decimal.Decimal `gorm:"type:decimal"`
}

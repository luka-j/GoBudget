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
	Envelopes     []TransactionEnvelope
}

type TransactionEnvelope struct {
	Database.Model
	Transaction Transaction
	Envelope    Envelope
	Amount      decimal.Decimal `gorm:"type:decimal"`
}

func CreateTransaction(payee string, bankAccountId string, note string, userId string) { //todo amount

}

func DeleteTransaction(id string, userId string) {

}

//todo get transactions with paging, different retrieval queries

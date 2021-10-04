package Entities

import (
	"GoBudget/Database"
	"github.com/shopspring/decimal"
)

type EnvelopeType string

const (
	Monthly EnvelopeType = "monthly"
	Yearly               = "yearly"
	Goal                 = "goal"
)

type Envelope struct {
	Database.Model
	BankAccount    BankAccount
	name           string
	EnvelopeType   EnvelopeType    `gorm:"type:envelope_type"`
	BudgetedAmount decimal.Decimal `gorm:"type:decimal"`
	CurrentAmount  decimal.Decimal `gorm:"type:decimal"`
}

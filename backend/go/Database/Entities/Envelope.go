package Entities

import (
	"GoBudget/Database"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Envelope struct {
	Database.Model
	UserID           uuid.UUID
	Name             string
	EnvelopeInterval int16           `gorm:"type:smallint"`
	BudgetedAmount   decimal.Decimal `gorm:"type:decimal"`
	CurrentAmount    decimal.Decimal `gorm:"type:decimal"`
}

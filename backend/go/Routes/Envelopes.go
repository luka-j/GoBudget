package Routes

import "github.com/shopspring/decimal"

type CreateEnvelopeRequest struct {
	Name     string          `json:"name" binding:"required"`
	Interval int16           `json:"type" binding:"required"`
	Budget   decimal.Decimal `json:"budget" binding:"required"`
}

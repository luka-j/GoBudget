package Routes

import "github.com/shopspring/decimal"

type EnvelopeAmount struct {
	EnvelopeId string          `json:"envelopeId"`
	Amount     decimal.Decimal `json:"amount"`
}

type CreateTransactionRequest struct {
	Payee         string           `json:"payee"`
	BankAccountId string           `json:"bankAccountId"`
	Note          string           `json:"note"`
	Amount        []EnvelopeAmount `json:"amount"`
}

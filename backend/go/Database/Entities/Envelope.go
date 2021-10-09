package Entities

import (
	"GoBudget/Database"
	"errors"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"strconv"
)

var ERROR_NO_SUCH_FIELD = errors.New("unknown field when attempting to patch model")

type Envelope struct {
	Database.Model
	UserID           uuid.UUID
	Name             string
	EnvelopeInterval int16           `gorm:"type:smallint"`
	BudgetedAmount   decimal.Decimal `gorm:"type:decimal"`
	CurrentAmount    decimal.Decimal `gorm:"type:decimal"`
}

func CreateEnvelope(name string, interval int16, budget decimal.Decimal, userId string) (Envelope, error) {
	user, err := GetUser(userId)
	if err != nil {
		return Envelope{}, err
	}
	envelope := Envelope{Name: name, EnvelopeInterval: interval, BudgetedAmount: budget, UserID: user.ID}
	createResult := Database.DB.Create(&envelope)
	if createResult.Error != nil {
		return Envelope{}, createResult.Error
	}
	return envelope, nil
}

func GetEnvelopesByUser(userId string) ([]Envelope, error) {
	var envelopes []Envelope
	getResult := Database.DB.Where("user_id=?", userId).Find(&envelopes)
	if getResult.Error != nil {
		return []Envelope{}, getResult.Error
	}
	return envelopes, nil
}

func EditEnvelope(id string, userId string, field string, value string) (Envelope, error) {
	var envelope Envelope
	getResult := Database.DB.Where("id=? AND user_id=?", id, userId).Take(&envelope)
	if getResult.Error != nil {
		return Envelope{}, getResult.Error
	}

	switch field {
	case "name":
		envelope.Name = value
	case "interval":
		interval, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			return Envelope{}, err
		}
		envelope.EnvelopeInterval = int16(interval)
	case "budget":
		budget, err := decimal.NewFromString(value)
		if err != nil {
			return Envelope{}, err
		}
		envelope.BudgetedAmount = budget
	default:
		return Envelope{}, ERROR_NO_SUCH_FIELD
	}

	return envelope, nil
}

func DeleteEnvelope(id string, userId string) error {
	var envelope Envelope
	getResult := Database.DB.Where("id=? AND user_id=?", id, userId).Take(&envelope)
	if getResult.Error != nil {
		return getResult.Error
	}
	result := Database.DB.Delete(&envelope)
	return result.Error
}

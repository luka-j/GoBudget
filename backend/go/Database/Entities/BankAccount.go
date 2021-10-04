package Entities

import (
	"GoBudget/Database"
	"github.com/google/uuid"
)

type BankAccount struct {
	Database.Model
	Name   string
	UserID uuid.UUID `gorm:"references:ID;foreignKey:user_id"`
}

func CreateAccount(name string, userId string) (BankAccount, error) {
	user, err := GetUser(userId)
	if err != nil {
		return BankAccount{}, err
	}
	account := BankAccount{Name: name, UserID: user.ID}
	createResult := Database.DB.Create(&account)
	if createResult.Error != nil {
		return BankAccount{}, createResult.Error
	}
	return account, nil
}

func GetAccountsByUser(userId string) ([]BankAccount, error) {
	var accounts []BankAccount
	getResult := Database.DB.Where("user_id=?", userId).Find(&accounts)
	if getResult.Error != nil {
		return []BankAccount{}, getResult.Error
	}
	return accounts, nil
}

func RenameAccount(id string, userId string, newName string) (BankAccount, error) {
	var account BankAccount
	getResult := Database.DB.Where("id=? AND user_id=?", id, userId).Take(&account)
	if getResult.Error != nil {
		return BankAccount{}, getResult.Error
	}
	account.Name = newName
	saveResult := Database.DB.Save(&account)
	if saveResult.Error != nil {
		return BankAccount{}, saveResult.Error
	}
	return account, nil
}

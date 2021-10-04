package Entities

import "GoBudget/Database"

type BankAccount struct {
	Database.Model
	Name string
	User User
}

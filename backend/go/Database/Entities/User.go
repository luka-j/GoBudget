package Entities

import (
	"GoBudget/Database"
	"gorm.io/gorm"

	"github.com/matthewhartstonge/argon2"
)

type User struct {
	Database.Model
	Username string
	Email    string `gorm:"unique"`
	Password string
}

func CreateUser(username, email, password string) {
	argon := argon2.DefaultConfig()
	hashedPassword, err := argon.HashEncoded([]byte(password))
	if err != nil {
		panic(err)
	}

	user := User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	createResult := Database.DB.Create(&user)
	if createResult.Error != nil {
		panic(createResult.Error)
	}
}

func LoginUser(email, password string) (User, error) {
	var user User
	takeResult := Database.DB.Where("email = ?", email).Take(&user)
	if takeResult.Error == gorm.ErrRecordNotFound {
		return User{}, takeResult.Error
	}
	if takeResult.Error != nil {
		panic(takeResult.Error)
	}

	ok, err := argon2.VerifyEncoded([]byte(password), []byte(user.Password))
	if err != nil {
		panic(err)
	}
	if ok {
		return user, nil
	} else {
		return User{}, gorm.ErrRecordNotFound
	}

}

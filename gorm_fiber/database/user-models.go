package database

import (
	"github.com/learn-go-projects/gorm_fiber/lib"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

func (db *Db) CreateUser(user *User) error {
	hash_pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash_pass)
	result := db.Query.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *Db) LoginUser(user *User) (string, error) {
	select_user := new(User)
	result := db.Query.Where("email = ?", user.Email).First(&select_user)
	if result.Error != nil {
		return "", result.Error
	}

	if err := bcrypt.CompareHashAndPassword([]byte(select_user.Password), []byte(user.Password)); err != nil {
		return "", err
	}
	return lib.CreateJwt(user.ID)
}

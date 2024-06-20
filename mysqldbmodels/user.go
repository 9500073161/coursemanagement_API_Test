// mysqldbmodels/user.go
package mysqldbmodels

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string `gorm:"password"`
}

func (db *DBClient) CreateUser(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := User{Email: email, Password: string(hashedPassword)}
	result := db.Conn.Create(&user)
	return result.Error
}

func (db *DBClient) GetUserByEmail(email string) (*User, error) {
	var user User
	result := db.Conn.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, result.Error
}

func (db *DBClient) AuthenticateUser(email, password string) (*User, error) {
	user, err := db.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

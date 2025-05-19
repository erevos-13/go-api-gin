package models

import (
	"errors"

	"example.com/gin-api/db"
	"example.com/gin-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	hashPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stm.Exec(u.Email, hashPass)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId
	return err
}

func (u *User) SignIn() error {
	query := `select id,password from users where email=?`
	row := db.DB.QueryRow(query, u.Email)
	var retrievePassword string
	err := row.Scan(&u.ID, &retrievePassword)
	if err != nil {
		return errors.New("invalid Credentials in sign in")
	}
	isValid := utils.VerifyPassword(u.Password, retrievePassword)
	if !isValid {
		return errors.New("invalid Credentials")
	}
	return nil
}

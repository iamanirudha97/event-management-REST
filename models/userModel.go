package models

import (
	"errors"
	"fmt"

	"example.com/eventbrite/db"
	"example.com/eventbrite/utils"
)

type User struct {
	Id       int64
	Email    string `binding: "required`
	Password string `binding: "required`
}

func (u User) SaveUser() (*int64, error) {
	query := "INSERT INTO users(email, password) VALUES ($1, $2) RETURNING id"

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		fmt.Println("Bcrypt Hash Error : ", err)
		return nil, err
	}

	var userId int64
	err = db.DB.QueryRow(query, u.Email, hashedPassword).Scan(&userId)
	if err != nil {
		fmt.Println("Error Creating new user")
		fmt.Println(err)
		return nil, err
	}

	return &userId, nil
}

func (u User) ValidateCredentials() error {
	query := "SELECT id,password FROM users WHERE email = $1 "

	var retrievedPassword string
	err := db.DB.QueryRow(query, u.Email).Scan(&u.Id, &retrievedPassword)
	if err != nil {
		fmt.Println("Can not find User")
		fmt.Println(err)
		return err
	}
	isValid := utils.IsPasswordValid(u.Password, retrievedPassword)

	if !isValid {
		return errors.New("invalid credentials")
	}
	return nil
}

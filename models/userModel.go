package models

import (
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

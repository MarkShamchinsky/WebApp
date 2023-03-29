package main

import (
	"fmt"
)

type User struct {
	ID        int    `json:"ID,-"`
	Login     string `json:"Login"`
	Password  string `json:"Password"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Role      string `json:"Role"`
}

func (u User) Create() error {
	row := connection.QueryRow(`INSERT INTO "User" 
    ("Login","Password","FirstName","LastName","Role") 
	VALUES ($1, $2, $3, $4, 'manager') RETURNING "ID"`,
		u.Login, u.Password, u.FirstName, u.LastName)

	e := row.Scan(&u.ID)
	fmt.Println(u.ID)
	if e != nil {
		return e
	}
	fmt.Println("Create new user with id", u.ID)
	return nil
}

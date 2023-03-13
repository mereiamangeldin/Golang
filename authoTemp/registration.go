package main

import (
	"Golang/assignment-1/models"
	"fmt"
)

type User struct {
	user_id      int
	name         string
	phone_number string
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := models.GetDB()
	if err != nil {
		fmt.Println("invalid name or password")
		panic(err)
	}
	fmt.Printf("Welcome %s\n", models.Name)
	defer db.Close()
	rows, err := db.Query("select user_id, name, phone_number from phone_numbers")
	CheckErr(err)
	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.user_id, &user.name, &user.phone_number)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, user)
	}
	for _, u := range users {
		fmt.Println(u.user_id, u.name, u.phone_number)
	}
}

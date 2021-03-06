package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mf-sakura/golang_study/db/sqlx/infrastructure"
	"github.com/mf-sakura/golang_study/db/sqlx/interfaces/controllers"
)

func main() {
	sqlHandler := infrastructure.NewSQLHandler()
	defer sqlHandler.Conn.Close()

	userController := controllers.NewUserController(sqlHandler.Conn)
	option := flag.String("a", "", "action")
	id := flag.String("i", "", "user id")
	firstName := flag.String("f", "Alan", "first name")
	lastName := flag.String("l", "Turing", "last name")
	flag.Parse()

	switch *option {
	case "index":
		users, err := userController.Index()
		if err != nil {
			log.Fatal(err)
		}
		if len(users) == 0 {
			fmt.Println("No users found.")
		}
		for _, user := range users {
			fmt.Printf("ID: %v, FirstName: %v, LastName: %v\n", user.ID, user.FirstName, user.LastName)
		}
		return
	case "show":
		if *id == "" {
			log.Fatal("You need a user.id.")
		}
		user, err := userController.Show(*id)
		if err != nil {
			log.Fatal(err)
		}
		if user == nil {
			fmt.Printf("No user found. ID: %v\n", *id)
		} else {
			fmt.Printf("ID: %v, FirstName: %v, LastName: %v\n", user.ID, user.FirstName, user.LastName)
		}
		return
	case "create":
		if *firstName == "" {
			log.Fatal("You need a first name.")
		}
		if *lastName == "" {
			log.Fatal("You need a last name.")
		}
		user, err := userController.Create(*firstName, *lastName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %v, FirstName: %v, LastName: %v\n", user.ID, user.FirstName, user.LastName)
		return
	default:
		log.Fatal("Unrecognized option.")
	}
}

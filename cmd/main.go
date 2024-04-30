package main

import (
	"database/sql"
	"fmt"
	"log"

	add "github.com/Abduazim0811/dictionary/integral/Add"
	enguzb "github.com/Abduazim0811/dictionary/integral/english-uzbek"
	uzbeng "github.com/Abduazim0811/dictionary/integral/uzbek-english"
	models "github.com/Abduazim0811/dictionary/integral/models"
	_ "github.com/lib/pq"
)

func main() {
	var choice int

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "localhost", "Abdu0811", "dictionary")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("[1] Uzb-eng\t[2] Eng-uzb\t[3] Add\t[4] Exit")
	fmt.Scanln(&choice)

	models.Models(db)

	switch choice {
	case 1:
		uzbeng.Uzeng(db)
	case 2:
		enguzb.Enguz(db)
	case 3:
		add.Add(db)
	case 4:
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid choice")
	}
}

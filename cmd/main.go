package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	var (
		son int
	)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "localhost", "Abdu0811", "dictionary")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Print("[1]Uzb-eng\t[2]Eng-uzb\t[3]Add")
	fmt.Scanln(&son)
	switch son{
	case 1:
		
	}

}

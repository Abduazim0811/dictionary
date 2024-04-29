package main

import (
	"database/sql"
	"fmt"
	"log"

	eng_uzb "github.com/Abduazim0811/dictionary/integral/english-uzbek"
	uzb_eng "github.com/Abduazim0811/dictionary/integral/uzbek-english"
	add "github.com/Abduazim0811/dictionary/integral/Add"
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
	fmt.Print("[1]Uzb-eng\t[2]Eng-uzb\t[3]Add\t[4]Exit")
	fmt.Scanln(&son)
	switch son {
	case 1:
		uzb_eng.Uzeng(db)
	case 2:
		eng_uzb.Enguz(db)

	case 3:
		add.Add(db)
	}

}

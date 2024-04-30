package add

import (
	"database/sql"
	"fmt"
	"log"
)

func Add(db *sql.DB) {
	var (
		english, uzbek string
	)
	i := 0
	for i > 0 {
		son:=0
		fmt.Print("english: ")
		fmt.Scanln(&english)
		fmt.Print("uzbek: ")
		fmt.Scanln(&uzbek)
		query := "Insert into dict(english,uzbek) values($1,$2)"
		_, err := db.Exec(query, english, uzbek)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Malumot qoshildi")
		fmt.Println("Yana malumot qushishni hohlaysizmi?\n[1]HA\t[2]Yoq")
		fmt.Scanln(&son)
		if son==2{
			break
		}
	}

}

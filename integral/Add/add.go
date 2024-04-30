package add

import (
	"database/sql"
	"fmt"
	"log"
)

func Add(db *sql.DB) {
	var (
		english,uzbek string
	)
	fmt.Print("english: ")
	fmt.Scanln(&english)
	fmt.Print("uzbek: ")
	fmt.Scanln(&uzbek)
	
	query:="Insert into dict(english,uzbek) values($1,$2)"
	_,err:=db.Exec(query,english,uzbek)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Malumot qoshildi")
}

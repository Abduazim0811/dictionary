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
	_,err:=db.Exec(`
	Create table if not exists dict(
		id serial primary key,
		english VARCHAR(100) unique not null,
		uzbek VARCHAR(100) unique not null
	)`)
	if err!=nil{
		log.Fatal(err)
	}

	query:="Insert into dict(english,uzbek) values($1,$2)"
	_,err=db.Exec(query,english,uzbek)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Malumot qoshildi")
}

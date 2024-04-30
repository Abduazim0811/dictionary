package models

import (
	"database/sql"
	"log"
)

var MP map[string]string

func Models(db *sql.DB) {
	MP = make(map[string]string)

	rows, err := db.Query("SELECT english, uzbek FROM dict;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var english, uzbek string

		if err := rows.Scan(&english, &uzbek); err != nil {
			log.Fatal(err)
		}
		MP[english] = uzbek
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

}

package main

import (
    "database/sql"
    "fmt"
    "log"

    add "github.com/Abduazim0811/dictionary/integral/Add"
    enguzb "github.com/Abduazim0811/dictionary/integral/english-uzbek"
    models "github.com/Abduazim0811/dictionary/integral/models"
    uzbeng "github.com/Abduazim0811/dictionary/integral/uzbek-english"
    _ "github.com/lib/pq"
)

func main() {
    var choice int
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
                       "localhost", 5432, "postgres", "Abdu0811", "dictionary")
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }
    defer db.Close()

    _, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS dict(
        id SERIAL PRIMARY KEY,
        english VARCHAR(100) UNIQUE NOT NULL,
        uzbek VARCHAR(100) UNIQUE NOT NULL);
    `)
    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }

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
    default:
        fmt.Println("Invalid choice")
    }
}


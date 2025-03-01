package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	// sqlite-tietokanta luodaan, jos sellaista ei ole
	dsn := "database.sqlite"
	DB, err = sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalf("Virhe avatessa tietokantaa: %v", err)
	}

	// Varmistetaan, että yhteys toimii
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Virhe yhdistäessä tietokantaan: %v", err)
	}

	fmt.Println("Yhteys tietokantaan muodostettu")
}

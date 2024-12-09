package config

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

/*
OpenConnectionPostgresSQL

	fungsi untuk membuka open koneksi ke database postgresql.
*/
func OpenConnectionPostgresSQL() (*sql.DB, error) {

	// deklarasi variabel yang dibutuhkan untuk koneksi database
	host := os.Getenv("localhost")
	port := os.Getenv("5432")
	user := os.Getenv("postgres")
	password := os.Getenv("postgres")
	dbname := os.Getenv("user")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if password == "" {
		password = "postgres"
	}
	if dbname == "" {
		dbname = "postgres"
	}

	psqlMerge := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// buat koneksi ke database.
	dbConnection, err := sql.Open("postgres", psqlMerge)
	if err != nil {
		fmt.Println("Error connecting to database")
		return nil, err
	}

	// ping ke database
	err = dbConnection.Ping()
	if err != nil {
		fmt.Println("Error pinging database")
		return nil, err
	}

	return dbConnection, nil
}

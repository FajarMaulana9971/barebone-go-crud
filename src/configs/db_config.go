package configs

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type configs struct {
	DB   *sql.DB
	PORT string
	DSN  string
}

func LoadConfig() (*configs, error) {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found")
	}

	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Println("DSN not configure")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Port not configure")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// err ini menampung error dikarenakan bukan dari namanya, tetapi dari function db.ping nya itu return error. sama seperti line 19
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// bisa dijadikan seperti ini tapi bikin bingung njing
	// if err = db.Ping() ; err != nil {
	// 	return nil, err
	// }

	return &configs{
		DB:   db,
		PORT: port,
		DSN:  dsn,
	}, nil

}

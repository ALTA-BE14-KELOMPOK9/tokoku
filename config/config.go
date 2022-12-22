package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string
	DBNAME string
}

func ReadConfig() *Config {
	// Load .env ke local pc
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Gagal membaca .env", err.Error())
	}

	config := Config{}
	config.DBUSER = os.Getenv("DBUSER")
	config.DBPASS = os.Getenv("DBPASS")
	config.DBHOST = os.Getenv("DBHOST")
	config.DBPORT = os.Getenv("DBPORT")
	config.DBNAME = os.Getenv("DBNAME")

	return &config
}

func OpenConnection(c Config) *sql.DB {
	// format source username:password@tcp(host:port)/databaseName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.DBUSER, c.DBPASS, c.DBHOST, c.DBPORT, c.DBNAME)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Terjadi error", err.Error())
	}

	return db
}

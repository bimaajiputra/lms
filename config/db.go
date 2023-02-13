package config

import (
	"database/sql"
	"fmt"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConPg() *sql.DB {
	var err error
	godotenv.Load(".env")
	connString := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=disable", os.Getenv("DB_NAME_PG"), os.Getenv("DB_USER_PG"), os.Getenv("DB_PASSWORD_PG"), os.Getenv("DB_HOST_PG"))
	Db, err := sql.Open(os.Getenv("DB_DRIVER_PG"), connString)

	if err != nil {
		panic(err)
	}
	//fmt.Println("Sukses Konek ke Db!XX")
	return Db
}

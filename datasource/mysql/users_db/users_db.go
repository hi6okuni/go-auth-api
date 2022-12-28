package users_db

import (
	"database/sql"
	"log"

	// "os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
	// username = os.Getenv("MYSQL_USER")
	// password = os.Getenv("MYSQL_PASSWORD")
	// host     = os.Getenv("MYSQL_HOST")
	// schema   = os.Getenv("MYSQL_DATABASE")
)

func init() {
	// username:password@tcp(host)/user_schema
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", "docker:password@tcp(godockerDB)/godocker?charset=utf8")
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}

package db

import (
	"database/sql"
	"fmt"
	"log"
	"reporting-utility/internal/utils/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	
	var err error
	dsn := config.AppConfig.Database.DSN
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("DB open error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	fmt.Println("✅ MySQL Connected")
}

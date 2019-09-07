package db

import (
	"database/sql"

	"github.com/sqsinformatique/backend/cfg"
	"github.com/sqsinformatique/backend/utils"

	_ "github.com/jackc/pgx/stdlib"
)

// The "db" package level variable will hold the reference to our database instance
var db *sql.DB

func InitDB() (err error) {
	if db != nil {
		return
	}
	utils.Infoln("Connect to database:", cfg.DefaultCfg.DSN)
	db, err = sql.Open("pgx", cfg.DefaultCfg.DSN)
	if err != nil {
		return
	}
	err = db.Ping()
	return
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		utils.Fatal(err)
	}
}

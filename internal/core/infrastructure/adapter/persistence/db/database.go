package db

import (
	"database/sql"
	_ "embed"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jonathanch7/go-modern-api/internal/shared/config"
	"log"
)

//go:embed migrations/tables.sql
var migrationsTables string

type Database struct {
	DB *sql.DB
}

func ConnectToMysql(config config.ConfigDB) *Database {
	db := openConnection("mysql", config.User+":"+config.Password+"@tcp("+config.Host+")/"+config.Schema+"?parseTime=true")
	database := &Database{db}
	database.execMigrations()
	return database
}

func (d *Database) execMigrations() {
	log.Println("executing migrations in database")
	if _, err := d.DB.Exec(migrationsTables); err != nil {
		log.Println("error executing migrations ", err.Error())
	}
}

func openConnection(dbDriver, url string) *sql.DB {
	log.Println("connecting with database ", dbDriver)
	db, err := sql.Open(dbDriver, url)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

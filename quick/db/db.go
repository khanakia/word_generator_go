package db

import (
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

type DB struct {
	DB      *sql.DB
	Dialect string
}

func New(rootdir string) *DB {
	driver := viper.GetString("database.driver")
	if driver == "postgres" {
		return NewPostgres()
	}

	return NewSqlite(rootdir)
}

func NewPostgres() *DB {
	username := viper.GetString("database.user")
	password := viper.GetString("database.password")
	databaseName := viper.GetString("database.name")
	databaseHost := viper.GetString("database.host")
	databasePort := viper.GetString("database.port")
	sslmode := viper.GetString("database.sslmode")

	dbDSN := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s port=%s", databaseHost, username, databaseName, sslmode, password, databasePort)

	config, err := pgx.ParseConfig(dbDSN)
	if err != nil {
		panic(err)
	}

	db := stdlib.OpenDB(*config)

	// db, err := sql.Open("postgres", dbDSN)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return &DB{
		DB:      db,
		Dialect: dialect.Postgres,
	}
}

func NewSqlite(rootdir string) *DB {
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/db.sqlite3?cache=shared&_fk=1&_journal_mode=WAL", rootdir))
	if err != nil {
		panic(fmt.Errorf("error initializing db: %w", err))
	}

	return &DB{
		DB:      db,
		Dialect: dialect.SQLite,
	}
}

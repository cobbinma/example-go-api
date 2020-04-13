package postgres

import (
	"database/sql"
	"fmt"
	"github.com/cobbinma/example-go-api/config"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DBClient interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	GetPet(query string, args ...interface{}) (*dbPet, error)
	GetPets(query string, args ...interface{}) (dbPets, error)
	Ping() error
	DB() *sql.DB
}

type dbClient struct {
	db *sqlx.DB
}

func NewDBClient() DBClient {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		config.DBHost,
		config.DBName,
		config.DBUser,
		config.DBPassword)

	driver := "postgres"

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		logrus.Fatalln("Could not open database: ", err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	return &dbClient{db: db}
}

func (dbc *dbClient) Exec(query string, args ...interface{}) (sql.Result, error) {
	return dbc.db.Exec(query, args...)
}

func (dbc *dbClient) GetPets(query string, args ...interface{}) (dbPets, error) {
	pets := dbPets{}
	if err := dbc.db.Select(&pets, query, args...); err != nil {
		return nil, err
	}
	return pets, nil
}

func (dbc *dbClient) GetPet(query string, args ...interface{}) (*dbPet, error) {
	pet := dbPet{}
	if err := dbc.db.Get(&pet, query, args...); err != nil {
		return nil, err
	}
	return &pet, nil
}

func (dbc *dbClient) Ping() error {
	return dbc.db.Ping()
}

func (dbc *dbClient) DB() *sql.DB {
	return dbc.db.DB
}

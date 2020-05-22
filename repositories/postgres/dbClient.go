package postgres

import (
	"database/sql"
	"fmt"
	"github.com/cobbinma/example-go-api/config"
	"github.com/jmoiron/sqlx"
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

func NewDBClient() (*dbClient, func() error, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s",
		config.DBHost,
		config.DBName,
		config.DBUser,
		config.DBPassword,
		config.DBSSLMode)

	driver := "postgres"

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open database : %w", err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	dbc := &dbClient{db: db}

	return dbc, dbc.Close, nil
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

func (dbc *dbClient) Close() error {
	return dbc.DB().Close()
}

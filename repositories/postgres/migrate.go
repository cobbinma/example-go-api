package postgres

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	post "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

func (p *postgres) Migrate() error {
	driver, err := post.WithInstance(p.dbClient.DB(), &post.Config{})
	if err != nil {
		return fmt.Errorf("could not create database driver : %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://files/migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("error instantiating migrate : %w", err)
	}

	version, dirty, _ := m.Version()
	logrus.Infof("Database version %d, dirty %t", version, dirty)

	logrus.Infof("Starting migration")

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("an error occurred while syncing the database.. %w", err)
	}

	logrus.Infoln("Migration complete")
	return nil
}

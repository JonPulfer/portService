package repository

import (
	"database/sql"
	"fmt"

	"github.com/JonPulfer/portsService/internal/portService"
)

type postgresPortDB struct {
	db *sql.DB
}

type PostgresConfig struct {
	DBURL string
}

func NewPostgresPortRepos(pgconfig PostgresConfig) (PortRepository, error) {
	db, err := sql.Open("postgres", pgconfig.DBURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open the db: %w", err)
	}

	return &postgresPortDB{db: db}, nil
}

func (p postgresPortDB) Store(port portService.Port) error {
	//TODO implement me
	panic("implement me")
}

func (p postgresPortDB) Fetch(portID string) (portService.Port, error) {
	//TODO implement me
	panic("implement me")
}

func (p postgresPortDB) Update(port portService.Port) error {
	//TODO implement me
	panic("implement me")
}

package repository

import (
	"database/sql"
	"log"

	"github.com/lucitez/later/pkg/repository/util"

	"github.com/lucitez/later/pkg/model"

	// Postgres driver
	_ "github.com/lib/pq"
)

// Hostname ...
type Hostname struct {
	DB *sql.DB
}

// NewHostname for wire generation
func NewHostname(db *sql.DB) Hostname {
	return Hostname{db}
}

var hostnameSelectStatement = util.GenerateSelectStatement(model.Hostname{}, "hostnames")

// Insert inserts a new hostname
func (repository *Hostname) Insert(hostname model.Hostname) error {

	statement := util.GenerateInsertStatement(hostname, "hostnames")

	_, err := repository.DB.Exec(
		statement,
		hostname.ID,
		hostname.Hostname,
		hostname.ContentType,
		hostname.CreatedAt,
		hostname.UpdatedAt,
		hostname.DeletedAt,
	)

	return err
}

// ByHostname gets a hostname by the hostname name
func (repository *Hostname) ByHostname(hn string) *model.Hostname {
	var hostname model.Hostname

	statement := hostnameSelectStatement + `
	WHERE hostname = $1
	AND deleted_at IS NULL;
	`

	row := repository.DB.QueryRow(statement, hn)

	return hostname.ScanRow(row)
}

// All returns all hostnames
func (repository *Hostname) All(limit int) []model.Hostname {
	statement := hostnameSelectStatement + `
	WHERE deleted_at IS NULL
	LIMIT $1;
	`
	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		log.Fatal(err)
	}

	return repository.scanRows(rows)
}

func (repository *Hostname) scanRows(rows *sql.Rows) []model.Hostname {
	hostnames := []model.Hostname{}

	defer rows.Close()

	for rows.Next() {
		var hostname model.Hostname
		hostname.ScanRows(rows)
		hostnames = append(hostnames, hostname)
	}

	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return hostnames
}

package repository

import (
	"database/sql"
	"later/pkg/repository/util"

	"later/pkg/model"

	// Postgres driver
	_ "github.com/lib/pq"
)

// Domain ...
type Domain struct {
	DB *sql.DB
}

// NewDomain for wire generation
func NewDomain(db *sql.DB) Domain {
	return Domain{db}
}

var domainSelectStatement = util.GenerateSelectStatement(model.Domain{}, "domains")

// Insert inserts a new domain
func (repository *Domain) Insert(domain *model.Domain) (*model.Domain, error) {

	statement := util.GenerateInsertStatement(*domain, "domains")

	_, err := repository.DB.Exec(
		statement,
		domain.ID,
		domain.Domain,
		domain.ContentType,
		domain.CreatedAt,
		domain.UpdatedAt,
		domain.DeletedAt)

	if err != nil {
		return nil, err
	}

	return domain, nil
}

// ByDomain gets a domain by the domain name
func (repository *Domain) ByDomain(domainName string) (*model.Domain, error) {
	var domain model.Domain

	statement := domainSelectStatement + `
	WHERE domain = $1
	AND deleted_at IS NULL;
	`

	row := repository.DB.QueryRow(statement, domainName)

	err := domain.ScanRow(row)

	return &domain, err
}

// All returns all domains
func (repository *Domain) All(limit int) ([]model.Domain, error) {
	statement := domainSelectStatement + `
	WHERE deleted_at IS NULL
	LIMIT $1;
	`
	rows, err := repository.DB.Query(statement, limit)

	if err != nil {
		return nil, err
	}

	return repository.scanRows(rows)
}

func (repository *Domain) scanRows(rows *sql.Rows) ([]model.Domain, error) {
	domains := []model.Domain{}

	defer rows.Close()

	for rows.Next() {
		var domain model.Domain
		err := domain.ScanRows(rows)

		if err != nil {
			return nil, err
		}
		domains = append(domains, domain)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}

	return domains, nil
}

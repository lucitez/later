package repository

import (
	"database/sql"

	// Postgres driver
	_ "github.com/lib/pq"
	"later/pkg/model"
)

// DomainRepository ...
type DomainRepository struct {
	DB *sql.DB
}

// NewDomainRepository for wire generation
func NewDomainRepository(db *sql.DB) DomainRepository {
	return DomainRepository{db}
}

// Insert inserts a new domain
func (repository *DomainRepository) Insert(domain *model.Domain) (*model.Domain, error) {

	statement := `
	INSERT INTO domains (id, domain, content_type)
	VALUES (
		$1,
		$2,
		$3
	)
	`

	_, err := repository.DB.Exec(
		statement,
		domain.ID,
		domain.Domain,
		domain.ContentType)

	if err != nil {
		return nil, err
	}

	return domain, nil
}

// ByDomain gets a domain by the domain name
func (repository *DomainRepository) ByDomain(domainName string) (*model.Domain, error) {
	var domain model.Domain

	statement := `
	SELECT * FROM domains 
	WHERE domain = $1
	`

	row := repository.DB.QueryRow(statement, domainName)

	err := row.Scan(
		&domain.ID,
		&domain.Domain,
		&domain.ContentType,
		&domain.CreatedAt,
		&domain.UpdatedAt,
		&domain.DeletedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &domain, nil
}

// All returns all domains
func (repository *DomainRepository) All(limit int) ([]model.Domain, error) {
	domains := []model.Domain{}

	rows, err := repository.DB.Query(`SELECT * FROM domains LIMIT $1`, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var domain model.Domain
		err := domain.ScanRows(rows)

		if err != nil {
			return nil, err
		}
		domains = append(domains, domain)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return domains, nil
}

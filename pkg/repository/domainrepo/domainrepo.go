package domainrepo

import (
	"database/sql"

	// Postgres driver
	_ "github.com/lib/pq"
	"later.co/pkg/later/domain"
)

// DB is this repository's database connection
var DB *sql.DB

// Insert inserts a new domain
func Insert(domain *domain.Domain) (*domain.Domain, error) {

	statement := `
	INSERT INTO domains (id, domain, content_type)
	VALUES (
		$1,
		$2,
		$3
	)
	`

	_, err := DB.Exec(
		statement,
		domain.ID,
		domain.Domain,
		domain.ContentType)

	if err != nil {
		return nil, err
	}

	return domain, nil
}

// ByDomain gets a domain by id
func ByDomain(domainName string) (*domain.Domain, error) {
	var domain domain.Domain

	statement := `
	SELECT * FROM domains 
	WHERE domain = $1
	`

	row := DB.QueryRow(statement, domainName)

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
func All(limit int) ([]domain.Domain, error) {
	domains := []domain.Domain{}

	rows, err := DB.Query(`SELECT * FROM domains LIMIT $1`, limit)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var domain domain.Domain
		err := rows.Scan(
			&domain.ID,
			&domain.Domain,
			&domain.ContentType,
			&domain.CreatedAt,
			&domain.UpdatedAt,
			&domain.DeletedAt)

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

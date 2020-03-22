package entity

import "database/sql"

type entity interface {
	ScanRows(rows *sql.Rows) error
	ScanRow(row *sql.Row) error
}

package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// DbGetMultipleRow : To fetch multiple records.
func (dbService DBHandler) DbGetMultipleRow(query string, args ...interface{}) (pgx.Rows, error) {
	var err error
	DB, err = dbService.InitDbReader()
	if err == nil {
		rows, exeError := DB.Query(context.Background(), query, args...)
		return rows, exeError
	}
	return nil, err
}

// DbGetSingleRow : To get a single record.
func (dbService DBHandler) DbGetSingleRow(query string) (pgx.Row, error) {
	DB, err := dbService.InitDbReader()
	if err == nil {
		rows := DB.QueryRow(context.Background(), query)
		return rows, err
	}
	return nil, err
}

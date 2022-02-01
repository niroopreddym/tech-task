package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// DbExecuteScalar : To insert/update records.
func (dbService DBHandler) DbExecuteScalar(query string, args ...interface{}) (pgx.Rows, error) {
	var err error
	DB, err = dbService.InitDbWriter()
	if err == nil {
		rows, err := DB.Query(context.Background(), query, args...)
		if err != nil {
			return nil, err
		}
		return rows, nil
	}
	return nil, err
}

// DbExecuteConflict : To check foreign key violations...
func (dbService DBHandler) DbExecuteConflict(query string, args ...interface{}) error {
	var err error
	DB, err = dbService.InitDbWriter()
	if err == nil {
		_, exeError := DB.Exec(context.Background(), query, args...)
		if exeError != nil {
			return exeError
		}
	}
	return nil
}

// DbWriter : To insert/update records in the database.
func (dbService DBHandler) DbWriter(query string) error {
	var err error
	DB, err = dbService.InitDbWriter()
	if err == nil {
		_, exeError := DB.Query(context.Background(), query)
		return exeError
	}
	return err
}

// DbExecuteQuery : To execute queries.
func (dbService DBHandler) DbExecuteQuery(query string) (int64, error) {
	var err error
	DB, err = dbService.InitDbWriter()
	if err == nil {
		rows, execErr := DB.Exec(context.Background(), query)
		if execErr == nil {
			var rowAffected int64
			rowAffected = rows.RowsAffected()
			return rowAffected, nil
		}
		return 0, execErr
	}
	return 0, err
}

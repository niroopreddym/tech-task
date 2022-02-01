package database

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

//DbIface exposes db interconnection methods
type DbIface interface {
	TxBegin() (pgx.Tx, error)
	InitDbReader() (*pgx.Conn, error)
	InitDbWriter() (*pgx.Conn, error)
	CreateConnection(connstring string) (*pgx.Conn, error)
	TxQuery(tx pgx.Tx, query string) (pgx.Rows, error)
	DbGetMultipleRow(query string, args ...interface{}) (pgx.Rows, error)
	DbExecuteQuery(query string) (int64, error)
	DbExecuteScalar(query string, args ...interface{}) (pgx.Rows, error)
	DbExecuteConflict(query string, args ...interface{}) error
	DbWriter(query string) error
	TxExecuteStmt(tx pgx.Tx, query string, args ...interface{}) (pgconn.CommandTag, error)
	TxComplete(tx pgx.Tx) error
	DbClose()
}

package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

var (
	// DB : for the database at the global space
	DB *pgx.Conn
	// ReaderConnectionString ...
	ReaderConnectionString string
	// WriterConnectionString ...
	WriterConnectionString string
)

//DBHandler provides the class implementation for DbIface interface
type DBHandler struct {
	DatabaseService DbIface
}

// DBNewHandler ...
func DBNewHandler() *DBHandler {
	SetConnectionStrings()
	return &DBHandler{
		DatabaseService: nil,
	}
}

// SetConnectionStrings ...
func SetConnectionStrings() {
	// "server=192.168.99.100 port=5432 user=postgres password=postgres sslmode=disable dbname=testdb"
	ReaderConnectionString = "postgres://postgres:postgres@127.0.0.1:5432/postgres"
	WriterConnectionString = "postgres://postgres:postgres@127.0.0.1:5432/postgres"
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// InitDbReader : for the database at the global space
func (dbService DBHandler) InitDbReader() (*pgx.Conn, error) {
	var err error
	DB, err = dbService.CreateConnection(ReaderConnectionString)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

// InitDbWriter : for the database at the global space
func (dbService DBHandler) InitDbWriter() (*pgx.Conn, error) {
	DB, err := dbService.CreateConnection(WriterConnectionString)
	if err != nil {
		return nil, err
	}

	return DB, nil
}

// CreateConnection : Creates the Connection
func (dbService DBHandler) CreateConnection(connectionString string) (*pgx.Conn, error) {
	var err error
	DB, err = pgx.Connect(context.Background(), connectionString)
	if err != nil {
		panic(err)
	}
	return DB, nil
}

//DbClose : Close the DB connectivity.
func (dbService DBHandler) DbClose() {
	err := DB.Close(context.Background())
	if err != nil {
		log.Println(err)
	}
}

//TxQuery : To execute a query and fetch rows. This will typically perform an insert & select (or) a plain select.
func (dbService DBHandler) TxQuery(tx pgx.Tx, query string) (pgx.Rows, error) {
	rows, err := tx.Query(context.Background(), query)
	if err != nil {
		if rberror := tx.Rollback(context.Background()); rberror != nil {
			return nil, rberror
		}
		return nil, err
	}

	return rows, nil
}

//TxBegin : To begin transaction.
func (dbService DBHandler) TxBegin() (pgx.Tx, error) {
	var err error
	DB, err = dbService.CreateConnection(WriterConnectionString)
	if err != nil {
		return nil, err
	}

	tx, err := DB.Begin(context.Background())
	return tx, err
}

//TxExecuteStmt : Executes the Query. Usually an INSERT/UPDATE.
func (dbService DBHandler) TxExecuteStmt(tx pgx.Tx, query string, args ...interface{}) (pgconn.CommandTag, error) {
	res, err := tx.Exec(context.Background(), query, args...)
	if err != nil {
		if rberror := tx.Rollback(context.Background()); rberror != nil {
			return nil, rberror
		}
		return nil, err
	}
	return res, nil
}

// TxComplete : Save Changes to the Database.
func (dbService DBHandler) TxComplete(tx pgx.Tx) error {
	if err := tx.Commit(context.Background()); err != nil {
		return err
	}
	return nil
}

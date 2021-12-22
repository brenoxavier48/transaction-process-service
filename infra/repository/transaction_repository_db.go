package repository

import (
	"database/sql"
	"log"
	"time"
)

type TransactionRepositoryDB struct {
	db *sql.DB
}

func NewTransactionRepositoryDB(db *sql.DB) *TransactionRepositoryDB {
	return &TransactionRepositoryDB{
		db: db,
	}
}

func (t *TransactionRepositoryDB) Insert(id string, accountId string, amount float64, status string, errorMessage string) error {
	stmt, err := t.db.Prepare(`
		insert into transactions (id, account_id, amount, status, error_message, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7)
	`)

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	_, err = stmt.Exec(id, accountId, amount, status, errorMessage, time.Now(), time.Now())

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}

package db

import (
	"context"
	"database/sql"
	"fmt"
)

// * Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// * NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// * execTx executes a function within a db transactions
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("Tx err: %v, rb err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

// * TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json: "from_account_id"`
	ToAccountID   int64 `json: "to_account_id"`
	Amount        int64 `json: "amount"`
}

// * TransferTxResult is the result of transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json: "transfer"`
	FromAccount int64    `json: "from_account"`
	ToAccount   int64    `json: "to_account"`
	FromEntry   Entry    `json: "from_entry"`
	ToEntry     Entry    `json: "to_entry"`
}

// * TransferTx performs a money transfer from one account to other.
// * It creates a transfer record, add account entries,
// * and update accounts' balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.Cre

		return nil
	})
}
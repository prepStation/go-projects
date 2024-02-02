package db

import (
	"context"
	"fmt"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

//Db transaction

// Defines all the functionalities needed to execute db transaction
// and Queries
// Because 	the queries struct only define functions
// that can retrieve or insert data on a single table
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error)
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
}

// Defines all the functionalities needed to execute db transaction
// and Queries
// Because 	the queries struct only define functions
// that can retrieve or insert data on a single table
type SQlStore struct {
	*Queries
	connPool *pgxpool.Pool
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQlStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

// execTx executes a function within a db transaction
func (store *SQlStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx error %v, rbErr %v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}

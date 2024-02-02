package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

// TransferTxParams defines the input for Transfer transaction
type VerifyEmailTxParams struct {
	EmailID    int64
	SecretCode string
}

// VerifyEmailTxResult defines the result of VerifyEmail transaction
type VerifyEmailTxResult struct {
	User        User
	VerifyEmail VerifyEmail
}

func (store *SQlStore) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error
		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			ID:         arg.EmailID,
			SecretCode: arg.SecretCode,
		})
		if err != nil {
			return err
		}

		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			IsEmailVerified: pgtype.Bool{
				Bool:  true,
				Valid: true,
			},
			Username: result.VerifyEmail.Username,
		})
		return err
	})

	return result, err

}

package gapi

import (
	"context"
	"fmt"
	"testing"
	"time"

	db "github.com/prepStation/simple_bank/db/sqlc"
	"github.com/prepStation/simple_bank/token"
	"github.com/prepStation/simple_bank/utils"
	"github.com/prepStation/simple_bank/worker"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)
	return server
}
func newContextWithBearerToken(t *testing.T, userName string, duration time.Duration, tokenMaker token.Maker) context.Context {
	accessToken, _, err := tokenMaker.CreateToken(userName, duration)
	require.NoError(t, err)
	bearerToken := fmt.Sprintf("%s %s", authorizationBearer, accessToken)
	md := metadata.MD{
		authorizationHeader: []string{
			bearerToken,
		},
	}
	return metadata.NewIncomingContext(context.Background(), md)
}

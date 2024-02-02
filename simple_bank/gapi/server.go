package gapi

import (
	"fmt"

	db "github.com/prepStation/simple_bank/db/sqlc"
	"github.com/prepStation/simple_bank/pb"
	"github.com/prepStation/simple_bank/token"
	"github.com/prepStation/simple_bank/utils"
	"github.com/prepStation/simple_bank/worker"
)

// Server serves grpc requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	store           db.Store
	tokenMaker      token.Maker
	config          utils.Config
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new grpc server and setup routing
func NewServer(config utils.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &Server{
		config:          config,
		tokenMaker:      tokenMaker,
		store:           store,
		taskDistributor: taskDistributor,
	}

	return server, nil
}

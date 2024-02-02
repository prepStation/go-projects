package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prepStation/simple_bank/utils"
)

var testStore Store

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load env config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatalf("Cannot connect to database %v\n", err)
	}
	testStore = NewStore(connPool)
	os.Exit(m.Run())
}

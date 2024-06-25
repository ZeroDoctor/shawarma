package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/zerodoctor/shawarma/internal/api"
	"github.com/zerodoctor/shawarma/internal/db/sqlite"
	"github.com/zerodoctor/shawarma/internal/logger"
)

var log *logrus.Logger = logger.Log

func main() {
	ctx := context.Background()

	db, err := sqlite.NewConnection(ctx, "shawarma.db")
	if err != nil {
		log.Fatalf("failed to init db [error=%s]", err.Error())
	}

	api := api.NewAPI(db)
	if err := api.Run(ctx, ":4000"); err != nil {
		log.Fatalf("failed to run api [error=%s]", err.Error())
	}
}

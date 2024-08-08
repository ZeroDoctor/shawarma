package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/zerodoctor/shawarma/internal/api"
	"github.com/zerodoctor/shawarma/internal/db/sqlite"
	"github.com/zerodoctor/shawarma/internal/logger"
	zdutil "github.com/zerodoctor/zdgo-util"
)

var log *logrus.Logger = logger.Log

func loadEnv() {
	env := os.Getenv("ENV")
	if env != "prod" {
		env = "dev"
		log.SetLevel(logrus.DebugLevel)
	}
	os.Setenv("ENV", env)
	wd, _ := zdutil.GetExecPath()
	log.Infof("loading api [env=%s] [commit=%s]...", env, os.Getenv("GIT_COMMIT"))
	godotenv.Load(wd + "/server/" + "." + env + ".env")
}

func main() {
	loadEnv()
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

package main

import (
	"context"

	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/pkg/pipeline"
)

func main() {
	_, err := db.NewSqliteDB()
	if err != nil {
		panic(err)
	}

	if err := pipeline.Run(context.Background(), "./example/example.pkl"); err != nil {
		panic(err)
	}
}

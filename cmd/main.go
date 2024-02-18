package main

import (
	"context"

	"github.com/zerodoctor/shawarma/pkg/pipeline"
)

func main() {
	err := pipeline.Run(context.Background())
	if err != nil {
		panic(err)
	}
}

package pipeline

import (
	"context"
	"fmt"

	"github.com/zerodoctor/shawarma/pkg/config"
)

func Run(ctx context.Context) error {
	conf, err := config.LoadFromPath(ctx, "./example/example.pkl")
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", conf.Pipeline)

	return nil
}

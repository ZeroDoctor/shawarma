package pipeline

import (
	"context"
	"fmt"

	"github.com/zerodoctor/shawarma/pkg/config"
)

func Run(ctx context.Context, path string) error {
	conf, err := config.LoadFromPath(ctx, path)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", conf.Pipeline)

	return nil
}

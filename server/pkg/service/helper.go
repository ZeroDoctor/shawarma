package service

import (
	"fmt"
	"strings"
)

func combindErr(errs []error) error {
	if errs == nil || len(errs) <= 0 {
		return nil
	}

	var builder strings.Builder
	for range errs {
		builder.WriteString("[error=%w]")
	}

	return fmt.Errorf(builder.String(), errs)
}

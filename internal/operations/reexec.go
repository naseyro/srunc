package operations

import (
	"fmt"

	"github.com/naseyro/srunc/internal/container"
)

type ReexecOpts struct {
	ID string
}

func Reexec(opts *ReexecOpts) error {
	cntr, err := container.Load(opts.ID)
	if err != nil {
		return fmt.Errorf("error loading container: %w", err)
	}

	if err := cntr.Reexec(); err != nil {
		return fmt.Errorf("error re executing container: %w", err)
	}

	return nil
}

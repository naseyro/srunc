package operations

import (
	"fmt"

	"github.com/naseyro/srunc/internal/container"
)

type DeleteOpts struct {
	ID    string
	Force bool
}

func Delete(opts *DeleteOpts) error {
	ctr, err := container.Load(opts.ID)
	if err != nil {
		return fmt.Errorf("error loading container %w", err)
	}
	if err = ctr.Delete(opts.Force); err != nil {
		return fmt.Errorf("error deleting container %w", err)
	}
	return nil
}

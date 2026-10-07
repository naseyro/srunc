package operations

import (
	"fmt"
	"strconv"

	"github.com/naseyro/srunc/internal/container"
	"golang.org/x/sys/unix"
)

type SignalOpts struct {
	ID     string
	Signal string
}

func Signal(opts *SignalOpts) error {
	cntr, err := container.Load(opts.ID)
	if err != nil {
		return fmt.Errorf("error loading container %d", opts.ID)
	}
	sgnl, err := strconv.Atoi(opts.Signal)
	if err != nil {
		return fmt.Errorf("error converting signal from stdin input")
	}
	if err := cntr.Signal(unix.Signal(sgnl)); err != nil {
		return fmt.Errorf("error sending signal %s to process %s", opts.Signal, opts.ID)
	}
	if err := cntr.Save(); err != nil {
		return fmt.Errorf("error saving container state")
	}
	return nil
}

package operations

import (
	"encoding/json"
	"fmt"

	"github.com/naseyro/srunc/internal/container"
)

type StateOpts struct {
	ID string
}

func State(opts *StateOpts) (string, error) {
	ctr, err := container.Load(opts.ID)
	if err != nil {
		return "", err
	}
	state, err := json.Marshal(ctr.State)
	if err != nil {
		return "", fmt.Errorf("error marshaling state: %w", err)
	}
	return string(state), nil
}

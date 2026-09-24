package operations

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/naseyro/srunc/internal/container"
	"github.com/opencontainers/runtime-spec/specs-go"
)

type CreateOpts struct {
	ID     string
	Bundle string
}

func Create(opts *CreateOpts) error {
	bundle, err := filepath.Abs(opts.Bundle)
	if err != nil {
		return fmt.Errorf("error retrieving absolute path from bundle: %w", err)
	}

	config, err := os.ReadFile(filepath.Join(bundle, "config.json"))
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	var spec *specs.Spec
	if err := json.Unmarshal(config, &spec); err != nil {
		return fmt.Errorf("error unmarshalling config: %w", err)
	}

	ctr, err := container.New(&container.NewContainerOpts{
		ID:     opts.ID,
		Bundle: bundle,
		Spec:   spec,
	})
	if err != nil {
		return fmt.Errorf("error creating container: %w", err)
	}

	if err := ctr.Save(); err != nil {
		return fmt.Errorf("error saving container state: %w", err)
	}
	return nil
}

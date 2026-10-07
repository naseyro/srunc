package cli

import (
	"github.com/naseyro/srunc/internal/operations"
	"github.com/spf13/cobra"
)

func signalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "signal [flags] CONTAINER_ID SIGNAL",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			containerID := args[0]
			signal := args[1]

			return operations.Signal(&operations.SignalOpts{
				ID:     containerID,
				Signal: signal,
			})
		},
	}

	return cmd
}

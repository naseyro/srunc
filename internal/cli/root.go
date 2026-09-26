package cli

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "srunc",
		SilenceUsage:  true,
		SilenceErrors: false,
	}
	cmd.AddCommand(
		stateCmd(),
		createCmd(),
		startCmd(),
		deleteCmd(),
		killCmd(),
		reexecCmd(),
	)
	return cmd
}

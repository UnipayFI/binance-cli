package cmd

import "github.com/spf13/cobra"

var (
	UnifiedCmd = &cobra.Command{
		Use:   "unified",
		Short: "unified",
	}
)

func init() {
	RootCmd.AddCommand(UnifiedCmd)
}

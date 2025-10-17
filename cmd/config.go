package cmd

import "github.com/spf13/cobra"

var configCmd = &cobra.Command{
	Use: "config",
}

func init() {
    configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configShowCmd)
}

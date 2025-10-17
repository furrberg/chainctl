package cmd

import (
    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "os"
    "text/tabwriter"
)

var configShowCmd = &cobra.Command{
	Use: "show",
	Short: "Shows current config",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch cfg.Output {
		case "table":
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "KEY\tVALUE")
			fmt.Fprintf(w, "provider\t%s\n", cfg.Provider)
			fmt.Fprintf(w, "rpc-url\t%s\n", cfg.RPCURL)
			fmt.Fprintf(w, "timeout\t%d\n", cfg.Timeout)
			fmt.Fprintf(w, "output\t%s\n", cfg.Output)
			w.Flush()
        case "json":
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			_ = enc.Encode(cfg)
		}
		return nil
    },
}
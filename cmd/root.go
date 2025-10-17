package cmd

import (
    "fmt"
    "github.com/furrberg/chainctl/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
    "os"
)

var cfg config.Config

var rootCmd = &cobra.Command{
	Use: "chainctl",
	Short: "The root command of the tool.",
	Long: "Use flags --provider to specify backend provider (mock|evm), --rpc-url to provide RPC endpoint URL (required if provider=evm), --output to set output format (table|json)",

}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func init() {
	rootCmd.AddCommand(configCmd)

	// global configuration flags with default values
	rootCmd.PersistentFlags().String("provider", "mock", "backend provider (mock|evm)")
	rootCmd.PersistentFlags().String("rpc-url", "", "RPC endpoint URL (required if provider=evm)")
	rootCmd.PersistentFlags().Int("timeout", 5, "timeout in seconds")
	rootCmd.PersistentFlags().String("output", "table", "output format (table|json)")

	viper.BindPFlag("provider", rootCmd.PersistentFlags().Lookup("provider"))
	viper.BindPFlag("rpc-url", rootCmd.PersistentFlags().Lookup("rpc-url"))
	viper.BindPFlag("timeout", rootCmd.PersistentFlags().Lookup("timeout"))
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))

	rootCmd.PersistentPreRunE = func (cmd *cobra.Command, args []string) error {
		viper.ReadInConfig()
		viper.Unmarshal(&cfg)
		return config.Validate(cfg)
	}
}
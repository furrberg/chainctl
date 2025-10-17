package config

import (
	"fmt"
)

type Config struct {
    Provider    string `mapstructure:"provider" yaml:"provider" json:"provider"`
    RPCURL      string `mapstructure:"rpc-url" yaml:"rpc-url" json:"rpc-url"`
    Timeout		int    `mapstructure:"timeout" yaml:"timeout" json:"timeout"`
    Output      string `mapstructure:"output" yaml:"output" json:"output"`
}


func Validate(c Config) error {
	 if c.Provider != "evm" && c.Provider != "mock" {
        return fmt.Errorf("unsupported provider: %s", c.Provider)
    }

	if c.Provider == "evm" && c.RPCURL == "" {
		return fmt.Errorf("No link given!")
	}

	if c.Timeout > 5 {
		return fmt.Errorf("Timeout value set too high")
	}

	if c.Output != "table" && c.Output != "json" {
		return fmt.Errorf("Unsupported output format!")
	}

	return nil
}

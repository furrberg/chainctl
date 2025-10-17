package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "gopkg.in/yaml.v3"
    "os"
)

var configInitCmd = &cobra.Command{
	Use: "init",
	Short: "Creates a config.yaml file with the copy of current config",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, errHomeDir := os.UserConfigDir()
        if errHomeDir != nil {
            return fmt.Errorf("the $HOME directory couldn't be located")
        }
        *&dir = dir + "/chainctl/"
        errChdir := os.Chdir(dir)
        if errChdir != nil {
            return fmt.Errorf("Wrong PATH")
        }

        file, errFile := os.Create("config.yaml")
        if errFile != nil {
            return fmt.Errorf("Couldn't resolve PATH to write the file to")
        }

        encoder := yaml.NewEncoder(file)
        encoder.Encode(cfg)

        return nil
    },
}
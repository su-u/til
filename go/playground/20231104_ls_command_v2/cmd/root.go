package cmd

import (
	"github.com/spf13/cobra"
	"github.com/su-u/ls_command_v2/ls"
)

var rootCmd = &cobra.Command{
	Use:  "ls",
	Long: "",
	Args: argsValid(),
	RunE: runE(),
}

func Execute() error {
	err := rootCmd.Execute()
	return err
}

func init() {
	rootCmd.Flags().StringP("", "d", ".", "ディレクトリ")
}

func argsValid() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		return nil
	}
}

func runE() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		dirname, err := cmd.Flags().GetString("")
		ls.WithDir(dirname)
		if err != nil {
			return err
		}
		return nil
	}
}

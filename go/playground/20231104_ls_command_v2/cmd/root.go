package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "example",
	Long: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := cmd.Flags().GetBool("toggle")
		if err != nil {
			return err
		}

		fmt.Printf("toggle: %t", t)
		return nil
	},
}

func Execute() error {
	err := rootCmd.Execute()
	return err
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "フラグの説明")
}

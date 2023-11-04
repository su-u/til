package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/su-u/ls_command_v2/ls"
)

var lsCmd = &cobra.Command{
	Use:  "ls",
	Long: "",
	Args: argsValid(),
	RunE: runE(),
}

func Execute() error {
	err := lsCmd.Execute()
	return err
}

func init() {
}

func argsValid() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) >= 2 {
			return errors.New("2つ以上の引数は指定できません")
		}
		return nil
	}
}

func runE() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		//fmt.Printf("%s\n", strings.Join(args, ""))

		var dirname = "."
		if len(args) >= 1 {
			dirname = args[0]
		}
		err := ls.WithDir(dirname)
		if err != nil {
			return err
		}
		return nil
	}
}

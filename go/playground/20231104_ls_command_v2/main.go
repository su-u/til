package main

import (
	"fmt"
	"os"

	"github.com/su-u/ls_command_v2/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

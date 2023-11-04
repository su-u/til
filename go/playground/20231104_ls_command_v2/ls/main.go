package ls

import (
	"fmt"
	"os"
)

func WithDir(dirname string) {
	if entries, err := os.ReadDir(dirname); err != nil {
		err = fmt.Errorf("func ReadDir error: %v", err)
		fmt.Println(err)
		os.Exit(1)
	} else {
		output(entries)
	}
}

func output(entries []os.DirEntry) {
	for _, entry := range entries {
		if entry.Name()[0] == '.' {
			continue
		}

		fmt.Printf("%s ", entry.Name())
	}
	fmt.Println()
}

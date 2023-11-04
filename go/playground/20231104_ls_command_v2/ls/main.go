package ls

import (
	"fmt"
	"os"
)

func WithDir(dirname string) error {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return err
	}

	output(entries)
	return nil
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

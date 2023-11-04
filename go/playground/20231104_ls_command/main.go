package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		ls()
	} else {
		lsWithDir(os.Args[1])
	}
}

func ls() {
	lsWithDir(".")
}

func lsWithDir(dirname string) {
	dirEntries, err := os.ReadDir(dirname)
	if err != nil {
		panic(err)
	}

	output(dirEntries)
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

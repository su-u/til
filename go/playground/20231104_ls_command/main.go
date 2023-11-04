package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		ls()
	} else if len(os.Args) == 2 {
		lsWithDir(os.Args[1])
	} else {
		fmt.Println("Usage: ls [dirname]")
		os.Exit(1)
	}
}

func ls() {
	lsWithDir(".")
}

func lsWithDir(dirname string) {
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

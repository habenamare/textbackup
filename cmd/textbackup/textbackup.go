package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/habenamare/textbackup/pkg/dirtree"
)

func main() {
	backupDirectory := flag.String("directory", ".", "Directory to backup")
	// exportFormat := flag.String("format", "json", "Follow symlinks")
	flag.Parse()

	// If the 'directory' flag was not set, make the value for that flag the
	// current working directory.
	if *backupDirectory == "." {
		wDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		*backupDirectory = wDir
	}

	rootFileNode, err := dirtree.RootFileNode(*backupDirectory)
	if err != nil {
		panic(err)
	}

	// Print a JSON representation of the root FileNode struct.
	fmt.Println(rootFileNode.JSON())
}

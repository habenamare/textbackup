// Package dirtree contains a routine to get a tree representation of a given
// directory.
package dirtree

import (
	"io"
	"os"
	"path/filepath"
)

func isEmptyDir(dirPath string) (bool, error) {
	dirF, err := os.Open(dirPath)
	if err != nil {
		return false, err
	}

	// Try to read at most 1 file from the dirPath directory.
	_, readErr := dirF.Readdir(1)
	dirF.Close()

	// If the end of the directory is reached while trying to read a single file
	// return true, otherwise false.
	if readErr == io.EOF {
		return true, nil
	}

	return false, readErr
}

// RootFileNode walks a directory specified by rootPath and returns a FileNode
// representation of the directory. This FileNode representation contains the
// files in the directory and it's sub-directories as FileNode structs.
func RootFileNode(rootPath string) (*FileNode, error) {
	var rootFileNode *FileNode

	currentParentPath := rootPath
	var currParentNode *FileNode

	var dirStack *DirStack

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if path == rootPath {
			rootFileNode = &FileNode{Name: info.Name()}

			// Initialize 'dirStack' with the root file node.
			dirStack = NewDirStack(rootFileNode)

			currParentNode = rootFileNode

			return nil
		}

		// Change 'currParentNode' and 'currentParentPath', if the parent path of
		// this path is different from the last 'currentParentPath'.
		parentPath := filepath.Dir(path)
		if parentPath != currentParentPath {
			currentParentPath = parentPath
			currParentNode = dirStack.Pop()
		}

		// Create a FileNode struct for this path and add it as a child to the
		// 'currentParentNode' FileNode struct.
		newNode := &FileNode{Name: info.Name()}
		currParentNode.AddChild(newNode)

		// If this path is a directory and if it is not empty, push it into 'dirStack'.
		if info.IsDir() {
			emptyDir, _ := isEmptyDir(path)
			if !emptyDir {
				dirStack.Push(newNode)
			}
		}

		return nil
	})

	return rootFileNode, err
}

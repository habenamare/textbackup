package dirtree

import "encoding/json"

// A FileNode is used to represent a file node in a directory tree.
// It can also be used to represent the directory tree itself, where
// the children of the directory can be traversed to using the
// Children FileNode array.
type FileNode struct {
	Name     string
	Children []*FileNode `json:",omitempty"`
}

// AddChild adds the f FileNode as a child to a FileNode.
func (pf *FileNode) AddChild(f *FileNode) {
	pf.Children = append(pf.Children, f)
}

// JSON returns a JSON representation of a FileNode as a string.
func (pf *FileNode) JSON() string {
	jsonB, err := json.Marshal(pf)
	if err != nil {
		return ""
	}

	return string(jsonB)
}

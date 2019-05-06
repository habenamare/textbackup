package dirtree

import (
	"fmt"
	"strings"
)

// A StackItem represents an item on a DirStack.
type StackItem struct {
	value *FileNode
	next  *StackItem
}

// A DirStack is a stack that should be initialized with a FileNode struct.
// After initialization, additional FileNode structs can be pushed and poped.
// When poping, if the size of the stack is 1, it will just be returned, not
// removed. This means that the FileNode struct used for initialization or the
// first FileNode struct pushed will never be removed.
type DirStack struct {
	top *StackItem
	len int
}

// NewDirStack returns a DirStack that is initialized with the f FileNode struct.
func NewDirStack(f *FileNode) *DirStack {
	return new(DirStack).init(f)
}
func (dS *DirStack) init(rootF *FileNode) *DirStack {
	dS.top = &StackItem{value: rootF}
	dS.len = 1

	return dS
}

// Push pushes the f FileNode struct into a DirStack.
func (dS *DirStack) Push(f *FileNode) {
	newTop := &StackItem{value: f, next: dS.top}

	dS.top = newTop
	dS.len++
}

// Pop returns and removes the top FileNode struct from a DirStack, which is
// the one inserted last. But it only removes the FileNode struct, if the
// length of the stack is > 1.
func (dS *DirStack) Pop() *FileNode {
	oldTop := dS.top

	if dS.len > 1 {
		newTop := oldTop.next
		dS.top = newTop
		dS.len--
	}

	return oldTop.value
}

// String returns a string representation of a DirStack.
func (dS *DirStack) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Length: %d\n", dS.len))
	b.WriteString("Items [bottom -> top]: ")

	var dSCopy DirStack
	dSCopy = *dS

	for i := dS.len; i > 0; i-- {
		topS := dSCopy.top.value.Name

		b.WriteString(fmt.Sprintf("%s -> ", topS))
		dSCopy.Pop()
	}

	return strings.TrimSuffix(b.String(), " -> ")
}

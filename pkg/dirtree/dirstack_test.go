package dirtree

import "testing"

func TestNewDirStack(t *testing.T) {
	rootFileNode := &FileNode{Name: "root"}
	dirStack := NewDirStack(rootFileNode)

	if dirStack.top.value != rootFileNode ||
		dirStack.top.next != nil ||
		dirStack.len != 1 {
		t.Error("DirStack was not initialized properly.")
	}
}

func TestPush(t *testing.T) {
	rootFileNode := &FileNode{Name: "root"}
	dirStack := NewDirStack(rootFileNode)

	aFileNode := &FileNode{Name: "Some FileNode struct"}
	dirStack.Push(aFileNode)

	if dirStack.top.value != aFileNode ||
		dirStack.top.next.value != rootFileNode ||
		dirStack.len != 2 {
		t.Error("[1] FileNode was not pushed to the DirStack properly.")
	}

	anotherFileNode := &FileNode{Name: "Another FileNode"}
	dirStack.Push(anotherFileNode)

	if dirStack.top.value != anotherFileNode ||
		dirStack.top.next.value != aFileNode ||
		dirStack.len != 3 {
		t.Error("[2] FileNode was not pushed to the DirStack properly.")
	}
}

func TestPop(t *testing.T) {
	rootFileNode := &FileNode{Name: "root"}
	dirStack := NewDirStack(rootFileNode)

	aFileNode := &FileNode{Name: "Some FileNode struct"}
	anotherFileNode := &FileNode{Name: "Another FileNode"}
	dirStack.Push(aFileNode)
	dirStack.Push(anotherFileNode)

	popedFileNode := dirStack.Pop()
	if popedFileNode != anotherFileNode ||
		dirStack.top.value != aFileNode ||
		dirStack.len != 2 {
		t.Error("[1] DirStack did not pop a FileNode properly.")
	}

	popedFileNode2 := dirStack.Pop()
	if popedFileNode2 != aFileNode ||
		dirStack.top.value != rootFileNode ||
		dirStack.len != 1 {
		t.Error("[2] DirStack did not pop a FileNode properly.")
	}

	popedFileNode3 := dirStack.Pop()
	if popedFileNode3 != rootFileNode ||
		dirStack.top.value != rootFileNode ||
		dirStack.len != 1 {
		t.Error("[1] DirStack did not perform the Pop operation properly.")
	}

	dirStack.Pop()
	dirStack.Pop()
	dirStack.Pop()
	popedFileNode4 := dirStack.Pop()
	if popedFileNode4 != rootFileNode ||
		dirStack.len != 1 {
		t.Error("[2] DirStack did not perform the Pop operation properly.")
	}

}

func TestString(t *testing.T) {
	rootFileNode := &FileNode{Name: "root"}
	dirStack := NewDirStack(rootFileNode)

	aFileNode := &FileNode{Name: "Some FileNode struct"}
	anotherFileNode := &FileNode{Name: "Another FileNode"}
	dirStack.Push(aFileNode)
	dirStack.Push(anotherFileNode)

	testString := "Length: 3\nItems [bottom -> top]: Another FileNode -> Some FileNode struct -> root"
	if dirStack.String() != testString {
		t.Error("DirStack string representation is incorrect.")
	}
}

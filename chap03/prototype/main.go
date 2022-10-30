package main

import "fmt"

type (
	// INode is the prototype interface
	INode interface {
		print(string)
		clone() INode
	}

	// File is a concrete prototype
	File struct {
		name string
	}

	// Folder is a concrete prototype
	Folder struct {
		name     string
		children []INode
	}
)

// File's collection of methods
func (s *File) print(indentation string) {
	fmt.Println(indentation + s.name)
}

func (s *File) clone() INode {
	return &File{name: s.name + "_clone"}
}

// Folder's collection of methods
func (s *Folder) print(indentation string) {
	fmt.Println(indentation + s.name)
	for _, child := range s.children {
		child.print(indentation + indentation)
	}
}

func (s *Folder) clone() INode {
	cloneFolder := &Folder{name: s.name + "_clone"}
	var tempChildren []INode
	for _, child := range s.children {
		tempChildren = append(tempChildren, child.clone())
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

// main function
func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		children: []INode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []INode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("Printing hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}

package main

import "fmt"

type (
	// Component is leaf
	Component interface {
		search(keyword string)
	}

	// File is component interface
	File struct {
		name string
	}

	// Folder is composite
	Folder struct {
		components []Component
		name       string
	}
)

// File's collection of methods
func (s *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, s.name)
}

func (s *File) getName() string {
	return s.name
}

// Folder's collection of methods
func (s *Folder) search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, s.name)
	for _, composite := range s.components {
		composite.search(keyword)
	}
}

func (s *Folder) add(c Component) {
	s.components = append(s.components, c)
}

// main function
func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{name: "Folder1"}
	folder1.add(file1)

	folder2 := &Folder{name: "Folder2"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)
	folder2.search("rose")
}

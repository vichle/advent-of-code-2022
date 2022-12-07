package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/vichle/advent-of-code-2022/shared"
)

type INodeType int

const (
	Directory INodeType = iota
	File
)

type INode struct {
	name     string
	_type    INodeType
	children []*INode
	parent   *INode
	size     int
}

func (node *INode) String() string {
	return node.recursiveToString("- ")
}

func (node *INode) recursiveToString(indent string) string {
	if node._type == File {
		return indent + node.name + " (file, size: " + strconv.Itoa(node.size) + ")\n"
	}
	childString := ""
	for _, child := range node.children {
		childString += child.recursiveToString(indent + " ")
	}
	return indent + node.name + " (dir) \n" + childString
}

func newDirectory(name string, parent *INode) INode {
	return INode{
		name:     name,
		_type:    Directory,
		children: make([]*INode, 0),
		parent:   parent,
		size:     -1,
	}
}

func newFile(name string, parent *INode, size int) INode {
	return INode{
		name:     name,
		_type:    File,
		children: make([]*INode, 0),
		parent:   parent,
		size:     size,
	}
}

func parseInput(lines []string) *INode {
	root := newDirectory("/", nil)
	current := &root
	for _, line := range lines {
		if line[0] == '$' {
			command := line[2:4]
			switch command {
			case "cd":
				rawParams := line[5:]
				if rawParams == ".." {
					current = current.parent
				} else if rawParams == "/" {
					current = &root
				} else {
					var found *INode
					for _, child := range current.children {
						if child.name == rawParams && child._type == Directory {
							found = child
							break
						}
					}
					if found == nil {
						panic("No child with name '" + rawParams + "' found")
					}
					current = found
				}
				break
			case "ls":
				// Do nothing, no other commands are followed by output
				break
			default:
				panic("Unknown command '" + line + "'")
			}
		} else {
			// The line contains output
			l := strings.Split(line, " ")
			name := l[1]
			var newInode INode
			if strings.HasPrefix(line, "dir") {
				newInode = newDirectory(name, current)
			} else {
				size, _ := strconv.Atoi(l[0])
				newInode = newFile(name, current, size)
			}
			current.children = append(current.children, &newInode)
		}
	}
	return &root
}

type DiskUsage struct {
	path string
	size int
}

func (du *DiskUsage) String() string {
	return du.path + " " + strconv.Itoa(du.size)
}

func (node *INode) du() []DiskUsage {
	return node.recursiveDiskUsage("")
}

func (node *INode) recursiveDiskUsage(path string) []DiskUsage {
	dus := make([]DiskUsage, 0)
	totals := 0
	newPath := path + node.name + "/"
	for _, child := range node.children {
		if child._type == File {
			fmt.Println("Counting " + child.name + " (" + strconv.Itoa(child.size) + ")")
			totals += child.size
		} else {
			childDus := child.recursiveDiskUsage(newPath)
			dus = append(dus, childDus...)
			totals += dus[len(dus)-1].size // Last element contains the total size of the child
		}
	}
	dus = append(dus, DiskUsage{
		path: newPath,
		size: totals,
	})

	return dus
}

func main() {
	contents := strings.TrimSpace(shared.ReadFileContents("input.txt"))
	root := parseInput(strings.Split(contents, "\n"))

	sizes := root.du()
	totalSpace := 70000000
	spaceNeededForUpdate := 30000000
	spaceFree := totalSpace - sizes[len(sizes)-1].size // "/" always ends up last
	needToDelete := spaceNeededForUpdate - spaceFree
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i].size < sizes[j].size
	})

	totalSizeUnder100k := 0
	var directoryToDelete DiskUsage
	found := false
	for _, size := range sizes {
		if !found && size.size > needToDelete {
			found = true
			directoryToDelete = size
		}
		if size.size < 100000 {
			totalSizeUnder100k += size.size
		}
	}

	fmt.Printf("Root: \n%v\n", root)
	fmt.Printf("Disk Usage: \n%v\n", sizes)
	fmt.Printf("Total < 100k: %v\n", totalSizeUnder100k)
	fmt.Printf("Directory to delete: %v (%v)\n", directoryToDelete.path, directoryToDelete.size)
}

package projects

import "fmt"

type DefaultVisitor struct {
}

func (d *DefaultVisitor) VisitFile(parent *ProjectItem, item *ProjectItem) {
	fmt.Println(fmt.Sprintf("Item name: %s , type: %s ", item.Name, "file"))
}

func (d *DefaultVisitor) VisitDir(parent *ProjectItem, item *ProjectItem) {
	fmt.Println(fmt.Sprintf("Item name: %s , type: %s ", item.Name, "dir"))
}

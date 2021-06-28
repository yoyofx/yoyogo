package projects

import "fmt"

type DefaultVisitor struct {
	current string
	p       *ProjectItem
}

func NewDefaultVisitor() *DefaultVisitor {
	return &DefaultVisitor{
		current: "",
		p:       nil,
	}
}

func (d *DefaultVisitor) getCurrent(parent *ProjectItem) string {
	if d.p != parent {
		d.current = d.current + "/" + parent.Name
		d.p = parent
	}
	return d.current
}

func (d *DefaultVisitor) VisitFile(parent *ProjectItem, item *ProjectItem) {
	current := d.getCurrent(parent)
	fmt.Println(fmt.Sprintf("Item name: %s/%s , type: %s ", current, item.Name, "file"))
}

func (d *DefaultVisitor) VisitDir(parent *ProjectItem, item *ProjectItem) {
	current := d.getCurrent(parent)
	fmt.Println(fmt.Sprintf("Item name: %s/%s , type: %s ", current, item.Name, "file"))
}

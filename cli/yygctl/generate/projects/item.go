package projects

import "path"

const (
	ProjectItemFile = iota
	ProjectItemDir
)

type ProjectItem struct {
	Name    string
	Path    string
	Type    int
	Content string
	Dom     []*ProjectItem
}

func NewProjectFile(fileName string) *ProjectItem {
	subItem := &ProjectItem{
		Name: fileName,
		Type: ProjectItemFile,
	}
	return subItem
}

func NewProjectDir(name string) *ProjectItem {
	subItem := &ProjectItem{
		Name: name,
		Type: ProjectItemDir,
		Path: "/",
	}
	return subItem
}

func (item *ProjectItem) AddFile(fileName string) {
	subItem := NewProjectFile(fileName)
	item.Dom = append(item.Dom, subItem)
}

func (item *ProjectItem) AddFileWithContent(fileName string, content string) {
	subItem := NewProjectFile(fileName)
	subItem.Content = content
	subItem.Path = path.Join(item.Path, fileName)
	item.Dom = append(item.Dom, subItem)
}

func (item *ProjectItem) AddDir(name string) *ProjectItem {
	subItem := NewProjectDir(name)
	subItem.Path = path.Join(item.Path, name)
	item.Dom = append(item.Dom, subItem)
	return subItem
}

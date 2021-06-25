package projects

type Project struct {
	Name    string
	Memo    string
	Path    string
	Dom     *ProjectItem
	visitor Visitor
}

func NewEmptyProject(name string, memo string) *Project {
	return &Project{
		Name:    name,
		Memo:    memo,
		Path:    "",
		Dom:     NewProjectDir(name),
		visitor: NewDefaultVisitor(),
	}
}

func (project *Project) SetVisitor(visitor Visitor) {
	project.visitor = visitor
}

func (project *Project) With(genFunc func(root *ProjectItem)) *Project {
	genFunc(project.Dom)
	return project
}

func (project *Project) List() {
	project.lookupItems(nil, project.Dom.Dom)
}

func (project *Project) lookupItems(parent *ProjectItem, items []*ProjectItem) {
	for _, item := range items {
		if project.visitor != nil {
			if item.Type == ProjectItemFile {
				project.visitor.VisitFile(parent, item)
			} else {
				project.visitor.VisitDir(parent, item)
			}
		}
		if item.Dom != nil {
			project.lookupItems(item, item.Dom)
		}
	}
}

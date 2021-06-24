package projects

type Visitor interface {
	VisitFile(parent *ProjectItem, file *ProjectItem)
	VisitDir(parent *ProjectItem, dir *ProjectItem)
}

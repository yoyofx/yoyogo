package projects

import (
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions/platform/consolecolors"
	"io/fs"
	"os"
	"path"
	"text/template"
	"time"
)

const logo = `
__    __  _____  __    __  _____   _____   _____        _____   _       _  
\ \  / / /  _  \ \ \  / / /  _  \ /  ___| /  _  \      /  ___| | |     | | 
 \ \/ /  | | | |  \ \/ /  | | | | | |     | | | |      | |     | |     | | 
  \  /   | | | |   \  /   | | | | | |  _  | | | |      | |     | |     | | 
  / /    | |_| |   / /    | |_| | | |_| | | |_| |      | |___  | |___  | | 
 /_/     \_____/  /_/     \_____/ \_____/ \_____/      \_____| |_____| |_| 
`

type Project struct {
	Name    string
	Memo    string
	Path    string
	Dom     *ProjectItem
	visitor Visitor
}

func (project *Project) CreateProject(data interface{}) {
	//根据项目名称创建文件夹
	fmt.Println("==========欢迎使用yoyogo cli  ==============================================")
	fmt.Println(consolecolors.Blue(logo))
	fmt.Println("===========================================================================")
	project.Path = path.Join(project.Path, project.Name)
	fmt.Println("项目目录" + project.Path)
	fmt.Println("正在生产项目模板......")
	time.Sleep(time.Second)
	//创建文件和文件夹
	/*	if project.Path=="" {
		project.Path=project.Name
	}*/
	CreateDirAndFile(project.Dom, project.Path, data)
}

func CreateDirAndFile(pI *ProjectItem, parentPath string, data interface{}) {
	currentPath := path.Join(parentPath, pI.Path)
	fmt.Println("currentPath" + currentPath)
	if pI.Type == ProjectItemDir {
		pI.Path = path.Join(currentPath, pI.Name)
		err := os.MkdirAll(pI.Path, fs.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		if len(pI.Dom) > 0 {
			for _, x := range pI.Dom {
				if x.Type == ProjectItemDir {
					CreateDirAndFile(x, pI.Path, data)
				} else {
					fmt.Println("create file: " + path.Join(currentPath, x.Name))
					file, err := os.Create(path.Join(currentPath, x.Name))
					if err != nil {
						fmt.Println(err)
					}
					tel, _ := template.New("console").Parse(x.Content)
					tel.Execute(file, data)
				}
			}
		}
	} else {
		fmt.Println("create file: " + path.Join(currentPath, pI.Name))
		file, err := os.Create(path.Join(currentPath, pI.Name))
		if err != nil {
			fmt.Println(err)
		}
		tel, _ := template.New("console").Parse(pI.Content)
		tel.Execute(file, data)
	}
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

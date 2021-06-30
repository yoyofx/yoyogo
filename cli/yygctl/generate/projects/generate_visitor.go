package projects

import (
	"fmt"
	"github.com/yoyofx/yoyogo"
	"html/template"
	"io/fs"
	"os"
	"path"
)

type Generator struct {
	ProjectName string            // 项目名称
	TargetDir   string            // 目标目录
	vars        map[string]string // 模板变量
}

func NewGenerator(projectName string, target string, vars map[string]string) *Generator {
	generator := &Generator{
		ProjectName: projectName,
		TargetDir:   target,
	}
	if vars != nil {
		generator.vars = vars
	} else {
		generator.vars = make(map[string]string)
	}
	generator.vars["ModelName"] = projectName
	generator.vars["Version"] = yoyogo.Version
	return generator
}

func (g *Generator) VisitFile(parent *ProjectItem, item *ProjectItem) {
	filepath := path.Join(g.TargetDir, item.Path)
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
	}
	if parent.Path == "/" {
		g.vars["CurrentModelName"] = "main"
	} else {
		g.vars["CurrentModelName"] = parent.Name
	}

	tel, _ := template.New("console").Parse(item.Content)
	err = tel.Execute(file, g.vars)
	_ = file.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("create file: " + filepath)
}

func (g *Generator) VisitDir(parent *ProjectItem, item *ProjectItem) {
	dirPath := path.Join(g.TargetDir, item.Path)
	err := os.MkdirAll(dirPath, fs.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("created dir: " + dirPath)
}

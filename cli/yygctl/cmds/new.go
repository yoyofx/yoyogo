package cmds

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoyofx/yoyogo/cli/yygctl/generate/projects"
	"github.com/yoyofx/yoyogo/cli/yygctl/generate/templates"
	"github.com/yoyofx/yoyogo/cli/yygctl/telplate"
	"strings"
)

var l bool
var projectName string
var dirPath string
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "create new yoyogo demo by template",
	Long:  "create new yoyogo demo by template",
	Args: func(cmd *cobra.Command, args []string) error {
		if !l && len(args) == 0 {
			return errors.New(" requires at least 1 arg(s), only received 0")
		}
		if l {
			fmt.Println(strings.Join(templates.GetProjectList(), ","))
			fmt.Println("console / webapi / mvc / grpc / xxl-job")
			return nil
		}
		telMap := [5]string{
			"console", "webapi", "mvc", "grpc", "xxl-job",
		}
		isHave := false
		for _, x := range telMap {
			if x == args[0] {
				isHave = true
			}
		}
		if !isHave {
			return errors.New("telpalte name dont have; console / webapi / mvc / grpc / xxl-job")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if !l {
			createProject()
		}
	},
}

func init() {
	NewCmd.Flags().BoolVarP(&l, "templates", "l", false, "list of template")
	NewCmd.Flags().StringVarP(&dirPath, "path", "p", "", "dir path")
	NewCmd.Flags().StringVarP(&projectName, "projectName", "n", "demo", "the name of project")
}

func createProject() {
	project := projects.Project{
		Name: projectName,
		Path: dirPath,
		Dom: &projects.ProjectItem{
			Name: projectName,
			Type: projects.ProjectItemDir,
		},
	}
	project.With(func(root *projects.ProjectItem) {
		root.AddFileWithContent("main.go", telplate.ConsoleMainTel)
		root.AddFileWithContent("startup.go", telplate.ConsoleStartUpTel)
		root.AddFileWithContent("go.mod", telplate.ConsoleGoModTel)
		root.AddFileWithContent("config.yml", telplate.ConsoleConfigTel)
		root.AddFileWithContent("service.go", telplate.ConsoleServiceTel)
	})
	data := map[string]interface{}{
		"ModelName": projectName,
	}
	project.CreateProject(data)
	/*var tPath string
	if dirPath == "" {
		tPath = projectName

	} else {
		tPath = dirPath + projectName
	}
	data := map[string]interface{}{
		"ModelName": projectName,
	}
	os.Mkdir(dirPath, fs.ModeDir)
	for _, x := range project.Dom.Dom {
		fileName:=path.Join(tPath,x.Name)
		_,createErr:= os.Create(fileName)
		if createErr!=nil {
			fmt.Println(createErr)
		}
		tel, _ := template.New("console").Parse(x.Content)
		file, err := os.OpenFile(dirPath+"/"+x.Name, os.O_CREATE|os.O_WRONLY, 0755)
		if err!=nil {
			fmt.Println(err)
		}
		tel.Execute(file, data)
	}*/
}

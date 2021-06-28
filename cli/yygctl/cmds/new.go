package cmds

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoyofx/yoyogo/cli/yygctl/generate/projects"
	"github.com/yoyofx/yoyogo/cli/yygctl/generate/templates"
	"github.com/yoyofx/yoyogo/cli/yygctl/telplate"
	"html/template"
	"io/fs"
	"os"
	"strings"
)

var l bool
var projectName string
var path string
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
		fmt.Println(projectName)
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
			fmt.Println("project:" + projectName)
			createProject()
		}

	},
}

func init() {
	NewCmd.Flags().BoolVarP(&l, "templates", "l", false, "list of template")
	NewCmd.Flags().StringVarP(&path, "path", "p", "", "dir path")
	NewCmd.Flags().StringVarP(&projectName, "projectName", "n", "demo", "the name of project")
}

func createProject() {
	project := projects.Project{
		Name: projectName,
		Path: path,
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
	var dirPath string
	if path == "" {
		dirPath = projectName

	} else {
		dirPath = path + projectName
	}
	data := map[string]interface{}{
		"ModelName": projectName,
	}
	os.Mkdir(dirPath, fs.ModeDir)
	for _, x := range project.Dom.Dom {
		tel, _ := template.New("demo").Parse(x.Content)
		file, err := os.OpenFile(dirPath+"/"+x.Name, os.O_CREATE|os.O_WRONLY, 0755)
		fmt.Println(err)
		tel.Execute(file, data)
	}
}

func createProjectMain() {

	data := map[string]interface{}{
		"ModelName": projectName,
	}
	tel, _ := template.New("demo").Parse(telplate.ConsoleMainTel)
	tel.Execute(os.Stdout, data)

}

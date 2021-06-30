package cmds

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoyofx/yoyogo"
	"github.com/yoyofx/yoyogo/abstractions/platform/consolecolors"
	"github.com/yoyofx/yoyogo/cli/yygctl/generate/templates"
	"strings"
	"time"
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
		if l { // l 表示什么呢？
			fmt.Println(strings.Join(templates.GetProjectList(), " / "))
			return nil
		}
		// 在注册里取不会没有。
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if !l {
			createProject(args[0])
		}
	},
}

func init() {
	NewCmd.Flags().BoolVarP(&l, "templates", "l", false, "list of template")
	NewCmd.Flags().StringVarP(&dirPath, "path", "p", "", "dir path")
	NewCmd.Flags().StringVarP(&projectName, "projectName", "n", "demo", "the name of project")
}

func createProject(template string) {
	// 所有模板定义： template/init.go
	// 模板目录 /template/console   定义： project_define.go
	project := templates.GetProjectByName(template)

	logo := yoyogo.Logo
	fmt.Println(consolecolors.Blue(string(logo)))
	fmt.Println(" ")
	fmt.Printf("%s   (version:  %s)", consolecolors.Green(":: YoyoGo ::"), consolecolors.Blue(yoyogo.Version))
	fmt.Print(consolecolors.Blue(`
light and fast , dependency injection based micro-service framework written in Go.
`))

	fmt.Println("正在生成项目模板......")
	time.Sleep(500 * time.Millisecond)

	project.Generate(dirPath, projectName)
	fmt.Println("生成项目模板完毕.")

	//---------------------------------------------------------------------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------------------------------------------------------------------
	//project := projects.Project{
	//	Name: projectName,
	//	Path: dirPath,
	//	Dom: &projects.ProjectItem{
	//		Name: projectName,
	//		Path: "/",
	//		Type: projects.ProjectItemDir,
	//	},
	//}
	//project.With(func(root *projects.ProjectItem) {
	//	root.AddFileWithContent("main.go", telplate.ConsoleMainTel)
	//	root.AddFileWithContent("startup.go", telplate.ConsoleStartUpTel)
	//	root.AddFileWithContent("go.mod", telplate.ConsoleGoModTel)
	//	root.AddFileWithContent("config.yml", telplate.ConsoleConfigTel)
	//	root.AddFileWithContent("service.go", telplate.ConsoleServiceTel)
	//})
	//data := map[string]interface{}{
	//	"ModelName": projectName,
	//}
	//project.CreateProject(data)
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

package cmds

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoyofx/yoyogo/cli/yygctl/telplate"
	"html/template"
	"os"
)

var l bool
var projectName string
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "create new yoyogo demo by template",
	Long:  "create new yoyogo demo by template",
	Args: func(cmd *cobra.Command, args []string) error {
		if !l && len(args) == 0 {
			return errors.New(" requires at least 1 arg(s), only received 0")
		}
		if l {
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
			createProjectMain()
		}

	},
}

func init() {
	NewCmd.Flags().BoolVarP(&l, "templates", "l", false, "list of template")
	NewCmd.Flags().StringVarP(&projectName, "projectName", "n", "demo", "the name of project")
}

func createProjectMain() {

	data := map[string]interface{}{
		"ModelName": projectName,
	}
	tel, _ := template.New("demo").Parse(telplate.ConsoleMainTel)
	tel.Execute(os.Stdout, data)

}

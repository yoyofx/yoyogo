package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoyofx/yoyogo/utils"
	"runtime"
	"strings"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "run application of yoyogo fx",
	Long:  `run application of yoyogo fx`,
	Run: func(cmd *cobra.Command, args []string) {
		runProject()
	},
}

func runProject() {
	fmt.Println("running project...")

	if runtime.GOOS == "windows" {
		runProjectWithWindows()
	} else {
		runProjectWithLinux()
	}
}

// linux下编译打包
func runProjectWithLinux() {
	pwd, _ := utils.ExecShell("pwd")
	pwd = strings.Replace(pwd, " ", "", -1)
	pwd = strings.Replace(pwd, "\r\n", "", -1)
	pwdArr := utils.Explode("/", pwd)
	if len(pwdArr) == 0 {
		return
	}
	projectName := pwdArr[len(pwdArr)-1]
	cmd := fmt.Sprintf("go run %s", pwd)

	fmt.Printf("project: %s, pwd: %s", projectName, pwd)
	fmt.Println(cmd)
	out, _ := utils.ExecShell(fmt.Sprintf("go run %s", pwd))
	fmt.Println(out)
}

// windows下编译打包
func runProjectWithWindows() {
	pwd, _ := utils.ExecShell("cd")
	pwd = strings.Replace(pwd, " ", "", -1)
	pwd = strings.Replace(pwd, "\r\n", "", -1)
	pwdArr := utils.Explode("\\", pwd)
	if len(pwdArr) == 0 {
		return
	}
	projectName := pwdArr[len(pwdArr)-1]
	cmd := fmt.Sprintf("go run %s", pwd)

	fmt.Printf("project: %s, pwd: %s", projectName, pwd)
	fmt.Println(cmd)
	out, _ := utils.ExecShell(fmt.Sprintf("go run %s", pwd))
	fmt.Println(out)
}

package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
	"strings"
	"yygctl/utils"
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "build Project application of yoyogo fx",
	Long:  `build Project application of yoyogo fx`,
	Run: func(cmd *cobra.Command, args []string) {
		buildProject()
	},
}

func buildProject() {
	fmt.Println("build all")

	if runtime.GOOS == "windows" {
		buildProjectWithWindows()
	} else {
		buildProjectWithLinux()
	}

	//移动静态文件
	utils.CopyPath("config"+utils.DirDot(), "build"+utils.DirDot()+"config")
	fmt.Println("build success")
}

// linux下编译打包
func buildProjectWithLinux() {
	pwd, _ := utils.ExecShell("pwd", "")
	pwd = strings.Replace(pwd, " ", "", -1)
	pwd = strings.Replace(pwd, "\r\n", "", -1)
	pwdArr := utils.Explode("/", pwd)
	if len(pwdArr) == 0 {
		return
	}
	projectName := pwdArr[len(pwdArr)-1]
	utils.ExecShell(fmt.Sprintf("go build -o build/%s", projectName), "")
}

// windows下编译打包
func buildProjectWithWindows() {
	pwd, _ := utils.ExecShell("cd", "")
	pwd = strings.Replace(pwd, " ", "", -1)
	pwd = strings.Replace(pwd, "\r\n", "", -1)
	pwdArr := utils.Explode("\\", pwd)
	if len(pwdArr) == 0 {
		return
	}
	projectName := pwdArr[len(pwdArr)-1]

	utils.ExecShell(fmt.Sprintf("go build -o build/%s.exe", projectName), "")
}

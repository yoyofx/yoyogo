package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

// ExecShell 执行shell命令
func ExecShell(shell string, dir string) (stdout string, stderr string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", shell)
	} else {
		cmd = exec.Command("/bin/bash", "-c", shell)
	}

	var out bytes.Buffer
	var stderrBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderrBuf
	if dir != "" {
		cmd.Dir = dir
	}
	err := cmd.Run()
	if err != nil {
		fmt.Println("shell exec have an error", "err", err)
	}

	return out.String(), stderrBuf.String()
}

//
//if err := cmd.Start(); err != nil {
//log.Fatal(err)
//}
//fmt.Println(cmd.Wait())
//}
//
//type outstream struct{}
//
//func (out outstream) Write(p []byte) (int, error) {
//	fmt.Println(string(p))
//	return len(p), nil
//}

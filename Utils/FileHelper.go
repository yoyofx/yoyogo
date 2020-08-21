package Utils

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CopyFile 拷贝文件
func CopyFile(src, dst string) bool {
	if len(src) == 0 || len(dst) == 0 {
		return false
	}
	srcFile, e := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
	if e != nil {
		panic(e)
	}
	defer srcFile.Close()

	dst = strings.Replace(dst, "\\", "/", -1)
	dstPathArr := strings.Split(dst, "/")
	dstPathArr = dstPathArr[0 : len(dstPathArr)-1]
	dstPath := strings.Join(dstPathArr, "/")

	dstFileInfo := GetFileInfo(dstPath)
	if dstFileInfo == nil {
		if e := os.MkdirAll(dstPath, os.ModePerm); e != nil {
			panic(e)
		}
	}

	dstFile, e := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if e != nil {
		panic(e)
	}
	defer dstFile.Close()

	if _, e := io.Copy(dstFile, srcFile); e != nil {
		panic(e)
	} else {
		return true
	}

}

// CopyPath 拷贝目录
func CopyPath(src, dst string) bool {
	srcFileInfo := GetFileInfo(src)
	if srcFileInfo == nil || !srcFileInfo.IsDir() {
		return false
	}
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		relationPath := strings.Replace(path, src, DirDot(), -1)
		dstPath := strings.TrimRight(strings.TrimRight(dst, "/"), "\\") + relationPath
		if !info.IsDir() {
			if CopyFile(path, dstPath) {
				return nil
			} else {
				return errors.New(path + " copy fail")
			}
		} else {
			if _, err := os.Stat(dstPath); err != nil {
				if os.IsNotExist(err) {
					if err := os.MkdirAll(dstPath, os.ModePerm); err != nil {
						panic(err)
						return err
					} else {
						return nil
					}
				} else {
					panic(err)
					return err
				}
			} else {
				return nil
			}
		}
	})

	if err != nil {
		panic(err)
	}
	return true

}

// GetFileInfo 获取文件信息
func GetFileInfo(src string) os.FileInfo {
	if fileInfo, e := os.Stat(src); e != nil {
		if os.IsNotExist(e) {
			return nil
		}
		return nil
	} else {
		return fileInfo
	}
}

func WriteFile(filePath string, content string) error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.WriteString(content)
	return err
}

func RemoveFile(filePath string) error {
	return os.Remove(filePath)
}

func DirDot() string {
	if runtime.GOOS == "windows" {
		return "\\"
	} else {
		return "/"
	}
}

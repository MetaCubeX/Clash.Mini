package common

import (
	"os"
	path "path/filepath"
	"runtime"
)

var (
	pwd		 		string
	executable 		string
	executableDir 	string
	osWindows		bool
)

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	executable, err = os.Executable()
	if err != nil {
		panic(err)
	}
	executableDir = path.Dir(executable)

	osWindows = runtime.GOOS == "windows"
}


// GetPwdPath 获取工作目录下的路径
func GetPwdPath(p... string) string {
	if len(p) == 0 {
		return pwd
	}
	return path.Join(append([]string{pwd}, p...)...)
}

// GetExecutable 获取应用路径
func GetExecutable() string {
	return executable
}

// GetExecutablePath 获取应用目录下的路径
func GetExecutablePath(p... string) string {
	if len(p) == 0 {
		return executableDir
	}
	return path.Join(append([]string{executableDir}, p...)...)
}

func IsWindows() bool {
	return osWindows
}

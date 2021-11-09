package uac

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"strings"
	"syscall"

	"github.com/Clash-Mini/Clash.Mini/constant"
)

const (
	CallFlagArg	= "uac-call"

	runasVerb   = `runas`
)

type Arg struct {
	ArgName 	string
	EqualValue 	string
	TackValue 	string
}

var (
	AmAdmin		= AmAdminNow()

	done		bool
	isUacCall	bool

	doFuncs		[]*func(maybeArgMap *map[string]*Arg, args []string) (done bool)
	doFuncMap	= make(map[string]*func(arg *Arg, args []string) (done bool))
)

func RunWhenAdmin() {
	err := BindFuncWithArg(CallFlagArg, func(arg *Arg, args []string) (done bool) {
		isUacCall = true
		return
	})
	if err != nil {
		// TODO:
		return
	}
	if AmAdmin {
		DoAllFuncWithArgs()
	}
	if isUacCall && done {
		os.Exit(0)
	}
	return
}

func DoAllFuncWithArgs() {
	args := os.Args
	argsCount := len(args)
	maybeArgMap := make(map[string]*Arg)
	for i, arg := range args {
		equalIndex := strings.Index(arg, "=")
		if equalIndex > -1 {
			argName := arg[:equalIndex]
			equalValue := arg[equalIndex:]
			argObj := &Arg{arg, equalValue, ""}
			if f, ok := doFuncMap[argName]; ok {
				if (*f)(argObj, args) {
					done = true
				}
				continue
			} else {
				maybeArgMap[argName] = argObj
			}
		} else {
			var tackValue string
			if i < argsCount - 1 {
				tackValue = args[i + 1]
			}
			argObj := &Arg{arg, "", tackValue}
			if f, ok := doFuncMap[arg]; ok {
				if (*f)(argObj, args) {
					done = true
				}
				continue
			} else {
				maybeArgMap[arg] = argObj
			}
		}
	}
	for _, f := range doFuncs {
		if (*f)(&maybeArgMap, args) {
			done = true
		}
	}
}

func AppendFunc(doFunc func(maybeArgMap *map[string]*Arg, args []string) (done bool)) {
	doFuncs = append(doFuncs, &doFunc)
}

func BindFuncWithArg(arg string, doFunc func(arg *Arg, args []string) (done bool)) error {
	if _, ok := doFuncMap[arg]; ok {
		return fmt.Errorf("[uac] a doFunc has been bound with the same arg [%s]", arg)
	}
	doFuncMap[arg] = &doFunc
	return nil
}

func CheckAndRunElevated(exe, args string) (err error) {
	if !AmAdmin {
		err = RunElevated(exe, args)
	} else {
		err = Run(exe, args)
	}
	return
}

func RunMeWithArg(argName, argValue string) error {
	var arg string
	if len(argValue) > 0 {
		arg = fmt.Sprintf("%s %s=%s", CallFlagArg, argName, argValue)
	} else {
		arg = fmt.Sprintf("%s %s", CallFlagArg, argName)
	}
	return RunElevated(constant.Executable, arg)
}

func RunElevated(exe, args string) (err error) {
	verbPtr, _ := syscall.UTF16PtrFromString(runasVerb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	err = windows.ShellExecute(0, verbPtr, exePtr, argPtr, nil, 0)
	if err != nil {
		return
	}
	return
}

func Run(exe, args string) (err error) {
	verbPtr, _ := syscall.UTF16PtrFromString("")
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	err = windows.ShellExecute(0, verbPtr, exePtr, argPtr, nil, 0)
	if err != nil {
		return
	}
	return
}

func AmAdminNow() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

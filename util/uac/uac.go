package uac

import (
	"fmt"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/MetaCubeX/Clash.Mini/constant"
	"github.com/MetaCubeX/Clash.Mini/log"
	stringUtils "github.com/MetaCubeX/Clash.Mini/util/string"
)

const (
	CallFlagArg = "call"

	runasVerb = `runas`
	//adminUser   = `Administrator`
)

type Arg struct {
	ArgName    string
	EqualValue string
	TackValue  string
}

type CallTime uint8

const (
	AnyTime CallTime = iota
	OnlyGeneral
	OnlyUac
)

var (
	AmAdmin = AmElevated()

	done bool
	//isUacCall			bool

	uacDoFuncs       []*func(maybeArgMap *map[string]*Arg, args []string) (done bool)
	generalDoFuncs   []*func(maybeArgMap *map[string]*Arg, args []string) (done bool)
	uacDoFuncMap     = make(map[string]*func(arg *Arg, args []string) (done bool))
	generalDoFuncMap = make(map[string]*func(arg *Arg, args []string) (done bool))
)

func RunWhenAdmin() {
	//err := BindFuncWithArg(CallFlagArg, AnyTime, func(arg *Arg, args []string) (done bool) {
	//	isUacCall = arg.EqualValue == "uac" && AmAdmin
	//	return
	//})
	//if err != nil {
	//	// TODO:
	//	return
	//}
	DoAllFuncWithArgs()
	if done {
		os.Exit(0)
	}
	return
}

func DoAllFuncWithArgs() {
	args := os.Args[1:]
	argsCount := len(args)
	maybeArgMap := make(map[string]*Arg)
	for i, arg := range args {
		equalIndex := strings.Index(arg, "=")
		if equalIndex > -1 {
			argName := arg[:equalIndex]
			equalValue := arg[equalIndex+1:]
			equalValue = unescapeArg(equalValue)
			argObj := &Arg{argName, equalValue, ""}

			if AmAdmin {
				f, exists := uacDoFuncMap[argName]
				if exists {
					if (*f)(argObj, args) {
						done = true
					}
					continue
				} else {
					maybeArgMap[argName] = argObj
				}
			}
			if f, exists := generalDoFuncMap[argName]; exists {
				if (*f)(argObj, args) {
					done = true
				}
				continue
			} else {
				maybeArgMap[argName] = argObj
			}
		} else {
			var tackValue string
			if i < argsCount-1 {
				tackValue = args[i+1]
				tmpArg := tackValue
				// TODO: refactor
				equalIndex := strings.Index(tmpArg, "=")
				if equalIndex > -1 {
					tmpArg = tmpArg[:equalIndex]
				}
				if _, exists := uacDoFuncMap[tmpArg]; exists {
					continue
				}
			}
			tackValue = unescapeArg(tackValue)
			argObj := &Arg{arg, "", tackValue}

			if AmAdmin {
				if f, ok := uacDoFuncMap[arg]; ok {
					if (*f)(argObj, args) {
						done = true
					}
					continue
				} else {
					maybeArgMap[arg] = argObj
				}
			}
			if f, ok := generalDoFuncMap[arg]; ok {
				if (*f)(argObj, args) {
					done = true
				}
				continue
			} else {
				maybeArgMap[arg] = argObj
			}
		}
	}
	if AmAdmin {
		for _, f := range uacDoFuncs {
			if (*f)(&maybeArgMap, args) {
				done = true
			}
		}
	}
	for _, f := range generalDoFuncs {
		if (*f)(&maybeArgMap, args) {
			done = true
		}
	}
}

func unescapeArg(s string) string {
	s = strings.TrimSpace(s)
	return stringUtils.UnescapeArgQuote(s)
}

func AppendFunc(callTime CallTime, doFunc func(maybeArgMap *map[string]*Arg, args []string) (done bool)) {
	if callTime <= OnlyUac {
		uacDoFuncs = append(uacDoFuncs, &doFunc)
	}
	if callTime <= OnlyGeneral {
		generalDoFuncs = append(generalDoFuncs, &doFunc)
	}
}

func BindFuncWithArg(arg string, callTime CallTime, doFunc func(arg *Arg, args []string) (done bool)) error {
	if callTime <= OnlyUac {
		_, exists := uacDoFuncMap[arg]
		if exists {
			return fmt.Errorf("[uac] a doFunc has been bound in uac funcs with the same arg [%s]", arg)
		} else {
			uacDoFuncMap[arg] = &doFunc
		}
	}
	if callTime <= OnlyGeneral {
		_, exists := generalDoFuncMap[arg]
		if exists {
			return fmt.Errorf("[uac] a doFunc has been bound in general funcs with the same arg [%s]", arg)
		} else {
			generalDoFuncMap[arg] = &doFunc
		}
	}
	return nil
}

func GetCallArg(arg string) string {
	return fmt.Sprintf("%s %s", CallFlagArg, arg)
}

func CheckAndRunAsElevated(exe, args string) (err error) {
	if !AmAdmin {
		err = RunAsElevate(exe, args)
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
	return RunAsElevate(constant.Executable, arg)
}

func RunAsElevate(exe, args string) (err error) {
	// TODO: get exit code
	verbPtr, _ := syscall.UTF16PtrFromString(runasVerb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	err = windows.ShellExecute(0, verbPtr, exePtr, argPtr, nil, 0)
	return
}

func Run(exe, args string) (err error) {
	command := exec.Command(exe, args)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := command.Output()
	log.Infoln("[uac] ran: %s", string(output))
	return err
	//verbPtr, _ := syscall.UTF16PtrFromString("")
	//exePtr, _ := syscall.UTF16PtrFromString(exe)
	//argPtr, _ := syscall.UTF16PtrFromString(args)
	//
	//err = windows.ShellExecute(0, verbPtr, exePtr, argPtr, nil, 0)
	//if err != nil {
	//	return
	//}
	//return
}

func AmElevated() bool {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, "", registry.ALL_ACCESS)
	defer k.Close()
	return err == nil
}

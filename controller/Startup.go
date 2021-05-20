package controller

import (
	"fmt"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
	"strings"
	"syscall"
)

var ExE = "Clash.Mini"
var Path64 = `SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Run`
var Path32 = `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`
var LM = registry.LOCAL_MACHINE

func CmdMain() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "enable":
			Startup()
			break
		case "disable":
			DeleteStartup()
			break
		}
	}
}

func Startup() {
	if windows.GetCurrentThreadEffectiveToken().IsElevated() {
		cmdEnable()
	} else {
		runMeElevated("enable")
	}
}

func DeleteStartup() {
	if windows.GetCurrentThreadEffectiveToken().IsElevated() {
		cmdDisable()
	} else {
		runMeElevated("disable")
	}
}

func cmdDisable() {
	k, err := registry.OpenKey(LM, Path64, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		k2, _ := registry.OpenKey(LM, Path32, registry.QUERY_VALUE|registry.SET_VALUE)
		err = k2.DeleteValue(ExE)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	err = k.DeleteValue(ExE)
	if err != nil {
		log.Fatal(err)
	}
}

func cmdEnable() {
	strEXEName := os.Args[0]
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, Path64, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		k2, err2 := registry.OpenKey(registry.LOCAL_MACHINE, Path32, registry.QUERY_VALUE|registry.SET_VALUE)
		if err2 != nil {
			log.Fatal(err2)
			return
		}
		err = k2.SetStringValue(ExE, strEXEName)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = k.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	err = k.SetStringValue(ExE, strEXEName)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = k.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func runMeElevated(args ...string) {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	var argsArr string
	if len(args) == 0 {
		args = os.Args[1:]
	}
	argsArr = strings.Join(args, " ")
	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(argsArr)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

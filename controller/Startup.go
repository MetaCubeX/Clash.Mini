package controller

import (
	"fmt"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

var (
	regArg     []string
	exeName    = "Clash.Mini"
	exePath, _ = os.Executable()
	Path64     = `SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Run`
	Path32     = `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`
	LM         = registry.LOCAL_MACHINE
	Verb       = "runas"
)

func Command(args string) {
	if windows.GetCurrentThreadEffectiveToken().IsElevated() {
		err := cmdReg(args, Path64)
		if err != nil {
			err := cmdReg(args, Path32)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		return
	} else {
		err := runMeElevated(args, Path64)
		if err != nil {
			err2 := runMeElevated(args, Path32)
			if err2 != nil {
				log.Fatal(err2)
			}
			return
		}
		return
	}
}

func cmdReg(args, Path string) error {

	k, err := registry.OpenKey(LM, Path, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
		return err
	}
	switch {
	case args == "add":
		err = k.SetStringValue(exeName, exePath)
		if err != nil {
			log.Fatal(err)
			return err
		}
	case args == "delete":
		err = k.DeleteValue(exeName)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	err = k.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func runMeElevated(args, path string) error {

	Reg := `reg.exe`
	RunLM := `HKEY_LOCAL_MACHINE`
	RunasPath := filepath.Join(RunLM, path)

	switch args {
	case `add`:
		regArg = []string{`add`, RunasPath, "/v", exeName, "/t", "REG_SZ", "/d", exePath, "/f"}
	case `delete`:
		regArg = []string{`delete`, RunasPath, "/v", exeName, "/f"}
	}
	regArgs := strings.Join(regArg, " ")

	verbPtr, _ := syscall.UTF16PtrFromString(Verb)
	exePtr, _ := syscall.UTF16PtrFromString(Reg)
	argPtr, _ := syscall.UTF16PtrFromString(regArgs)

	var showCmd int32 = 0 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, nil, showCmd)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func RegCompare() (b bool) {
	k, err := registry.OpenKey(LM, Path64, registry.QUERY_VALUE)
	if err != nil {
		k, err = registry.OpenKey(LM, Path64, registry.QUERY_VALUE)
		if err != nil {
			return false
		}
	}
	defer k.Close()

	_, _, err = k.GetStringValue(exeName)
	if err != nil {
		return false
	}
	return true

}

package controller

import (
	"fmt"
	"github.com/beevik/etree"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

var (
	Taskexe     = `schtasks`
	Verb        = "runas"
	taskName    = "Clash.Mini"
	taskTree    = `SOFTWARE\Clash.Mini`
	regArg      []string
	taskFile    = filepath.Join(".", "task.xml")
	taskPath, _ = os.Getwd()
	Filepath    = filepath.Join(taskPath, taskFile)
)

func RegCompare(cmd string) (b bool) {
	key, err := registry.OpenKey(registry.CURRENT_USER, taskTree, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer key.Close()
	s, _, _ := key.GetStringValue(cmd)
	if s == "ON" || s == "Lite" {
		return true
	} else {
		return false
	}
}

func Regcmd(cmd, value string) {
	key, exists, err := registry.CreateKey(registry.CURRENT_USER, taskTree, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()
	if exists {
		fmt.Println("键已存在")
	} else {
		fmt.Println("新建注册表键")
	}
	switch cmd {
	case "Task":
		if value == "ON" {
			key.SetStringValue("Task", "ON")
		} else {
			key.SetStringValue("Task", "OFF")
		}
	case "Sys":
		if value == "ON" {
			key.SetStringValue("Sys", "ON")
		} else {
			key.SetStringValue("Sys", "OFF")
		}
	case "MMBD":
		if value == "Lite" {
			key.SetStringValue("MMBD", "Lite")
		} else {
			key.SetStringValue("MMBD", "Max")
		}
	}
}

func TaskCommand(args string) error {
	switch args {
	case `create`:
		xml := TaskBuild()
		cache := ioutil.WriteFile(taskFile, xml, 0644)
		if cache != nil {
			return cache
		}
		regArg = []string{`/create`, `/tn`, taskName, `/XML`, Filepath}
		Regcmd("Task", "ON")
	case `delete`:
		regArg = []string{`/delete`, `/tn`, taskName, `/f`}
		Regcmd("Task", "OFF")
	}

	regArgs := strings.Join(regArg, " ")
	verbPtr, _ := syscall.UTF16PtrFromString(Verb)
	exePtr, _ := syscall.UTF16PtrFromString(Taskexe)
	argPtr, _ := syscall.UTF16PtrFromString(regArgs)

	var showCmd int32 = 0 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, nil, showCmd)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func TaskBuild() (xml []byte) {

	taskComName := os.Args[0]
	taskWorkingPath, _ := os.Getwd()
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="GBK"`)
	Task := doc.CreateElement(`Task`)
	Task.CreateAttr("version", "1.2")
	Task.CreateAttr("xmlns", `http://schemas.microsoft.com/windows/2004/02/mit/task`)
	RegistrationInfo := Task.CreateElement("RegistrationInfo")
	Date := RegistrationInfo.CreateElement("Date")
	Date.CreateText("2021-05-25T19:59:49.3146822")
	Author := RegistrationInfo.CreateElement("Author")
	Author.CreateText("Clash.Mini")
	Triggers := Task.CreateElement("Triggers")
	LogonTrigger := Triggers.CreateElement("LogonTrigger")
	LEnabled := LogonTrigger.CreateElement("Enabled")
	LEnabled.CreateText("true")
	Principals := Task.CreateElement("Principals")
	Principal := Principals.CreateElement("Principal")
	Principal.CreateAttr("id", "Author")
	GroupId := Principal.CreateElement("GroupId")
	GroupId.CreateText("S-1-5-32-544")
	RunLevel := Principal.CreateElement("RunLevel")
	RunLevel.CreateText("HighestAvailable")
	Settings := Task.CreateElement("Settings")
	MultipleInstancesPolicy := Settings.CreateElement("MultipleInstancesPolicy")
	MultipleInstancesPolicy.CreateText("IgnoreNew")
	DisallowStartIfOnBatteries := Settings.CreateElement("DisallowStartIfOnBatteries")
	DisallowStartIfOnBatteries.CreateText("true")
	StopIfGoingOnBatteries := Settings.CreateElement("StopIfGoingOnBatteries")
	StopIfGoingOnBatteries.CreateText("true")
	AllowHardTerminate := Settings.CreateElement("AllowHardTerminate")
	AllowHardTerminate.CreateText("true")
	StartWhenAvailable := Settings.CreateElement("StartWhenAvailable")
	StartWhenAvailable.CreateText("false")
	RunOnlyIfNetworkAvailable := Settings.CreateElement("RunOnlyIfNetworkAvailable")
	RunOnlyIfNetworkAvailable.CreateText("false")
	IdleSettings := Settings.CreateElement("IdleSettings")
	StopOnIdleEnd := IdleSettings.CreateElement("StopOnIdleEnd")
	StopOnIdleEnd.CreateText("true")
	RestartOnIdle := IdleSettings.CreateElement("RestartOnIdle")
	RestartOnIdle.CreateText("false")
	AllowStartOnDemand := Settings.CreateElement("AllowStartOnDemand")
	AllowStartOnDemand.CreateText("true")
	Enabled := Settings.CreateElement("Enabled")
	Enabled.CreateText("true")
	Hidden := Settings.CreateElement("Hidden")
	Hidden.CreateText("false")
	RunOnlyIfIdle := Settings.CreateElement("RunOnlyIfIdle")
	RunOnlyIfIdle.CreateText("false")
	WakeToRun := Settings.CreateElement("WakeToRun")
	WakeToRun.CreateText("false")
	ExecutionTimeLimit := Settings.CreateElement("ExecutionTimeLimit")
	ExecutionTimeLimit.CreateText("PT72H")
	Priority := Settings.CreateElement("Priority")
	Priority.CreateText("7")
	Actions := Task.CreateElement("Actions")
	Actions.CreateAttr("Context", "Author")
	Exec := Actions.CreateElement("Exec")
	TaskCom := Exec.CreateElement("Command")
	TaskCom.CreateText(taskComName)
	WorkingDirectory := Exec.CreateElement("WorkingDirectory")
	WorkingDirectory.CreateText(taskWorkingPath)
	doc.Indent(2)
	xml, _ = doc.WriteToBytes()
	return xml
}

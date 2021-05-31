package controller

import (
	"fmt"
	"github.com/Clash-Mini/Clash.Mini/constant"
	"io/ioutil"
	"os"
	path "path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/Clash-Mini/Clash.Mini/cmd"
	"github.com/Clash-Mini/Clash.Mini/cmd/cron"
	"github.com/Clash-Mini/Clash.Mini/cmd/mmdb"
	"github.com/Clash-Mini/Clash.Mini/cmd/sys"
	"github.com/Clash-Mini/Clash.Mini/cmd/task"
	"github.com/Clash-Mini/Clash.Mini/util"

	"github.com/Dreamacro/clash/log"
	"github.com/beevik/etree"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

const (
	taskExe     = `schtasks`
	runasVerb   = "runas"
	regTaskTree = `SOFTWARE\Clash.Mini`
	taskName    = "ClashMini"
)

var (
	taskFile = path.Join(".", "task.xml")
)

func RegCompare(command cmd.CommandType) (b bool) {
	key, err := registry.OpenKey(registry.CURRENT_USER, regTaskTree, registry.QUERY_VALUE)
	if err != nil {
		if os.IsNotExist(err) {
			err := RegCmd(cmd.Invalid)
			if err != nil {
				log.Errorln("RegCompare RegCmd error: %v", err)
				return false
			}
		} else {
			log.Errorln("RegCompare OpenKey error: %v", err)
		}
		return false
	}
	defer func(key registry.Key) {
		err := key.Close()
		if err != nil {
			log.Errorln("RegCompare Close error: %v", err)
			b = false
		}
	}(key)
	value, _, err := key.GetStringValue(command.GetName())
	if err != nil {
		log.Errorln("RegCompare GetStringValue error: %v", err)
		return false
	}
	// TODO: waiting for confirm cases
	switch command {
	case cmd.Task:
		return task.ParseType(value).IsON()
	case cmd.Sys:
		return sys.ParseType(value).IsON()
	case cmd.MMDB:
		return mmdb.ParseType(value).IsON()
	case cmd.Cron:
		return cron.ParseType(value).IsON()
	default:
		fmt.Printf("command \"%s\" is not support\n", command)
		return false
	}
}

// RegCmd 注册命令
func RegCmd(value cmd.GeneralType) error {
	key, exists, err := registry.CreateKey(registry.CURRENT_USER, regTaskTree, registry.ALL_ACCESS)
	if err != nil {
		log.Fatalln("RegCmd CreateKey failed: %v", err)
	}
	defer func(key registry.Key) {
		err := key.Close()
		if err != nil {
			log.Errorln("RegCmd Close error: %v", err)
			return
		}
	}(key)

	if exists {
		log.Warnln("注册表键已存在")
	} else {
		log.Infoln("新建注册表键值")
	}
	command := value.GetCommandType()
	if value == cmd.Invalid {
		log.Infoln("被动新建注册表键值")
		value = value.GetDefault()
	}
	if !command.IsValid(value) {
		return fmt.Errorf("command \"%s\" is not supported type \"%s\"", command.GetName(), value.String())
	}
	if err := key.SetStringValue(command.GetName(), value.String()); err != nil {
		return err
	}
	return nil
}

// getTaskRegArgs 拼接任务计划参数
func getTaskRegArgs(opera string, args ...string) string {
	regArg := append([]string{"/" + opera, `/tn`, taskName}, args...)
	return strings.Join(regArg, " ")
}

// TaskCommand 执行任务计划
func TaskCommand(taskType task.Type) (err error) {
	var taskArgs string
	switch taskType {
	case task.ON:
		xml, err := TaskBuild()
		if err != nil {
			return err
		}
		if err = ioutil.WriteFile(taskFile, xml, 0644); err != nil {
			return err
		}
		taskArgs = getTaskRegArgs("create", "/XML", path.Join(constant.PWD, taskFile))
		break
	case task.OFF:
		taskArgs = getTaskRegArgs("delete", "/f")
	}
	err = RegCmd(taskType)
	if err != nil {
		return err
	}

	verbPtr, _ := syscall.UTF16PtrFromString(runasVerb)
	exePtr, _ := syscall.UTF16PtrFromString(taskExe)
	argPtr, _ := syscall.UTF16PtrFromString(taskArgs)

	err = windows.ShellExecute(0, verbPtr, exePtr, argPtr, nil, 0)
	if err != nil {
		return err
	}
	return err
}

// TaskBuild 生成任务计划XML
func TaskBuild() (xml []byte, err error) {
	selfExeName := os.Args[0]
	selfWorkingPath, _ := os.Getwd()
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-16"`)
	tTask := doc.CreateElement(`Task`)
	tTask.CreateAttr("version", "1.2")
	tTask.CreateAttr("xmlns", `http://schemas.microsoft.com/windows/2004/02/mit/task`)

	tRegistrationInfo := tTask.CreateElement("RegistrationInfo")
	tDescription := tRegistrationInfo.CreateElement("Description")
	tDescription.CreateText("此任务将在用户登录后自动运行Clash.Mini，如果停用此任务将无法保持Clash.Mini自动运行。")
	tAuthor := tRegistrationInfo.CreateElement("Author")
	tAuthor.CreateText(util.AppTitle)
	tDate := tRegistrationInfo.CreateElement("Date")
	tDate.CreateText(time.Now().Format("2006-01-02T15:04:05"))

	tTriggers := tTask.CreateElement("Triggers")
	tLogonTrigger := tTriggers.CreateElement("LogonTrigger")
	tLogonTriggerEnabled := tLogonTrigger.CreateElement("Enabled")
	tLogonTriggerEnabled.CreateText("true")

	tPrincipals := tTask.CreateElement("Principals")
	tPrincipal := tPrincipals.CreateElement("Principal")
	tPrincipal.CreateAttr("id", "Author")
	tGroupId := tPrincipal.CreateElement("GroupId")
	tGroupId.CreateText("S-1-5-32-544")
	tRunLevel := tPrincipal.CreateElement("RunLevel")
	tRunLevel.CreateText("HighestAvailable")

	tSettings := tTask.CreateElement("Settings")
	tMultipleInstancesPolicy := tSettings.CreateElement("MultipleInstancesPolicy")
	tMultipleInstancesPolicy.CreateText("IgnoreNew")
	tDisallowStartIfOnBatteries := tSettings.CreateElement("DisallowStartIfOnBatteries")
	tDisallowStartIfOnBatteries.CreateText("false")
	tStopIfGoingOnBatteries := tSettings.CreateElement("StopIfGoingOnBatteries")
	tStopIfGoingOnBatteries.CreateText("false")
	tAllowHardTerminate := tSettings.CreateElement("AllowHardTerminate")
	tAllowHardTerminate.CreateText("true")
	tStartWhenAvailable := tSettings.CreateElement("StartWhenAvailable")
	tStartWhenAvailable.CreateText("false")
	tRunOnlyIfNetworkAvailable := tSettings.CreateElement("RunOnlyIfNetworkAvailable")
	tRunOnlyIfNetworkAvailable.CreateText("true")
	tIdleSettings := tSettings.CreateElement("IdleSettings")
	tStopOnIdleEnd := tIdleSettings.CreateElement("StopOnIdleEnd")
	tStopOnIdleEnd.CreateText("true")
	tRestartOnIdle := tIdleSettings.CreateElement("RestartOnIdle")
	tRestartOnIdle.CreateText("false")
	tAllowStartOnDemand := tSettings.CreateElement("AllowStartOnDemand")
	tAllowStartOnDemand.CreateText("true")
	tEnabled := tSettings.CreateElement("Enabled")
	tEnabled.CreateText("true")
	tHidden := tSettings.CreateElement("Hidden")
	tHidden.CreateText("false")
	tRunOnlyIfIdle := tSettings.CreateElement("RunOnlyIfIdle")
	tRunOnlyIfIdle.CreateText("false")
	tWakeToRun := tSettings.CreateElement("WakeToRun")
	tWakeToRun.CreateText("false")
	tExecutionTimeLimit := tSettings.CreateElement("ExecutionTimeLimit")
	tExecutionTimeLimit.CreateText("PT72H")
	tPriority := tSettings.CreateElement("Priority")
	tPriority.CreateText("7")

	tActions := tTask.CreateElement("Actions")
	tActions.CreateAttr("Context", "Author")
	tExec := tActions.CreateElement("Exec")
	tTaskCom := tExec.CreateElement("Command")
	tTaskCom.CreateText(selfExeName)
	tWorkingDirectory := tExec.CreateElement("WorkingDirectory")
	tWorkingDirectory.CreateText(selfWorkingPath)

	doc.Indent(2)
	xml, err = doc.WriteToBytes()
	if err != nil {
		return nil, err
	}
	xml, _, err = transform.Bytes(unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM).NewEncoder(), xml)
	if err != nil {
		return nil, err
	}
	return xml, err
}

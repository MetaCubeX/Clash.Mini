package controller

import (
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
	"time"

	"github.com/Clash-Mini/Clash.Mini/app"
	"github.com/Clash-Mini/Clash.Mini/cmd/task"
	"github.com/Clash-Mini/Clash.Mini/config"
	"github.com/Clash-Mini/Clash.Mini/constant"
	uacUtils "github.com/Clash-Mini/Clash.Mini/util/uac"

	"github.com/beevik/etree"
)

const (
	taskExe     = `schtasks`
	taskName    = `Clash.Mini`
)

// getTaskRegArgs 拼接任务计划参数
func getTaskRegArgs(opera string, args ...string) string {
	regArg := append([]string{"/" + opera, `/tn`, taskName}, args...)
	return strings.Join(regArg, " ")
}

// TaskCommandCall 执行任务计划并回调
func TaskCommandCall(taskType task.Type, callback func(err error)) {
	callback(TaskCommand(taskType))
}

func TaskCommand(taskType task.Type) (err error) {
	var taskArgs string
	switch taskType {
	case task.ON:
		xml, err := TaskBuild()
		if err != nil {
			return err
		}
		if err = ioutil.WriteFile(constant.TaskFile, xml, 0644); err != nil {
			return err
		}
		taskArgs = getTaskRegArgs("create", "/XML", constant.TaskFile)
		break
	case task.OFF:
		taskArgs = getTaskRegArgs("delete", "/f")
	}
	err = config.SetCmd(taskType)
	if err != nil {
		return
	}
	err = uacUtils.CheckAndRunMeElevated(taskExe, taskArgs)
	return err
}

// TaskBuild 生成任务计划XML
func TaskBuild() (xml []byte, err error) {
	selfExeName := constant.Executable
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-16"`)
	tTask := doc.CreateElement(`Task`)
	tTask.CreateAttr("version", "1.2")
	tTask.CreateAttr("xmlns", `http://schemas.microsoft.com/windows/2004/02/mit/task`)

	tRegistrationInfo := tTask.CreateElement("RegistrationInfo")
	tDescription := tRegistrationInfo.CreateElement("Description")
	tDescription.CreateText("此任务将在用户登录后自动运行Clash.Mini，如果停用此任务将无法保持Clash.Mini自动运行。")
	tAuthor := tRegistrationInfo.CreateElement("Author")
	tAuthor.CreateText(app.Name)
	tDate := tRegistrationInfo.CreateElement("Date")
	tDate.CreateText(time.Now().Format("2006-01-02T15:04:05"))

	tTriggers := tTask.CreateElement("Triggers")
	tLogonTrigger := tTriggers.CreateElement("LogonTrigger")
	tLogonTriggerEnabled := tLogonTrigger.CreateElement("Enabled")
	tLogonTriggerEnabled.CreateText("true")
	tLogonTriggerDelay := tLogonTrigger.CreateElement("Delay")
	tLogonTriggerDelay.CreateText("PT5S")

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
	tWorkingDirectory.CreateText(constant.ExecutableDir)

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

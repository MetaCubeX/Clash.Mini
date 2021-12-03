package task

import (
	"github.com/Clash-Mini/Clash.Mini/cmd"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Clash-Mini/Clash.Mini/app"
	"github.com/Clash-Mini/Clash.Mini/constant"
	uacUtils "github.com/Clash-Mini/Clash.Mini/util/uac"

	"github.com/beevik/etree"
)

const (
	taskExe  = `schtasks`
	taskName = `Clash.Mini`
)

type Type string

const (
	ON  Type = "on"
	OFF Type = "off"

	Invalid Type = ""
)

var (
	typeMap = map[string]Type{
		ON.String():  ON,
		OFF.String(): OFF,
	}
)

// String implements cmd.GeneralType
func (t Type) String() string {
	return string(t)
}

// GetCommandType implements cmd.GeneralType
func (t Type) GetCommandType() cmd.CommandType {
	return cmd.Task
}

// GetDefault implements cmd.GeneralType
func (t Type) GetDefault() cmd.GeneralType {
	return OFF
}

func ParseType(s string) Type {
	typeEnum, ok := typeMap[s]
	if !ok {
		return Invalid
	}
	return typeEnum
}

func ParseTypeWeak(s string) Type {
	s = strings.ToLower(s)
	return ParseType(s)
}

func (t Type) IsValid() bool {
	return t != Invalid && string(t) != ""
}

func IsValid(s string) bool {
	return ParseType(s).IsValid()
}

// IsPositive implements cmd.GeneralType
func (t Type) IsPositive() bool {
	return t == ON
}

// getTaskRegArgs 拼接任务计划参数
func getTaskRegArgs(opera string, args ...string) string {
	regArg := append([]string{"/" + opera, `/tn`, taskName}, args...)
	return strings.Join(regArg, " ")
}

func DoCommand(taskType Type) (err error) {
	var taskArgs string
	switch taskType {
	case ON:
		xml, err := buildSchtaskFile()
		if err != nil {
			return err
		}
		if err = ioutil.WriteFile(constant.TaskFile, xml, 0644); err != nil {
			return err
		}
		taskArgs = getTaskRegArgs("create", "/XML", strconv.Quote(constant.TaskFile))
		break
	case OFF:
		taskArgs = getTaskRegArgs("delete", "/f")
	}
	err = uacUtils.RunAsElevate(taskExe, taskArgs)

	if taskType == ON {
		time.Sleep(3 * time.Second)
		defer os.Remove(constant.TaskFile)
	}
	return err
}

// buildSchtaskFile 组装任务计划XML文件
func buildSchtaskFile() (xml []byte, err error) {
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

package app

import (
	"fmt"
	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/denisbrodbeck/machineid"
)

const (
	buggerLogHeader 	= "bugger"
)

var (
	machineId 	string
)

// InitBugsnag 初始化Bugsnag
func InitBugsnag() {
	writeMsg(buggerLogHeader, "initializing the bug reporter")
	var err error
	machineId, err = machineid.ProtectedID(fmt.Sprintf("%s-%s-%s", "clash.mini", "$MACHINE_ID_SECRET_VERSION$", "$MACHINE_ID_SECRET$"))
	if err != nil {
		writeMsg(buggerLogHeader, "cannot generate protected machine id: %v", err)
		machineId = "anonymous"
	} else {
		writeMsg(buggerLogHeader, "the machine id has been generated: %s", machineId)
	}
	bugsnag.OnBeforeNotify(func(event *bugsnag.Event, config *bugsnag.Configuration) error {
		event.User = &bugsnag.User{ Id: machineId }
		return nil
	})
	appVersion := fmt.Sprintf("%s-%s", Version, CommitId)
	stage := "$BRANCH$"
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:         "$BUGSNAG_KEY$",
		AppVersion: 	appVersion,
		ReleaseStage:   stage,
		ProjectPackages: []string{"main", "github.com/Clash-Mini/Clash.Mini/*"},
		AutoCaptureSessions: true,
		//Logger: log.Logger,
	})
	config := fmt.Sprintf("appVersion: [%s], stage: [%s], machineId: [%s]", appVersion, stage, machineId)
	writeMsg(buggerLogHeader, "initialized the bug reporter: %s", config)
}

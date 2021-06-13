package app

//import (
//	"fmt"
//	"github.com/bugsnag/bugsnag-go/v2"
//	"github.com/denisbrodbeck/machineid"
//)

//var (
//	machineId string
//)

func BuggerInit() {
//	id, err := machineid.ProtectedID("$MACHINE_ID_SECRET$")
//	if err != nil {
//		fmt.Println(fmt.Errorf("cannot get unique machine id: %v", err))
//	}
//	machineId = id
//	bugsnag.Configure(bugsnag.Configuration{
//		APIKey:         "$BUGSNAG_KEY$",
//		AppVersion: 	Version,
//		ReleaseStage:   "$BRANCH$",
//		ProjectPackages: []string{"github.com/Clash-Mini/Clash.Mini/*"},
//		AutoCaptureSessions: true,
//	})
//	bugsnag.OnBeforeNotify(func(event *bugsnag.Event, config *bugsnag.Configuration) error {
//		event.User = &bugsnag.User{ Id: machineId }
//		return nil
//	})
}

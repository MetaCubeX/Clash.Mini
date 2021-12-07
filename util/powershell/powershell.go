package powershell

import (
	"github.com/Clash-Mini/Clash.Mini/log"
	"os"
	"os/exec"
)

type PowerShell struct {
	powerShell string
}

func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}
func (p *PowerShell) Execute(args ...string) error {
	cmd := exec.Command(p.powerShell, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}
func ShowCmd() error {
	posh := New()
	err := posh.Execute("get-content", "-encoding utf8", log.RotateWriter.CurrentFileName(), "-Wait", "-Tail", "30")
	return err
}

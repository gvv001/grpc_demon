//go:build windows

package lib

import (
	"os/exec"
)

func ExecShellCommand(params string) (string, error) {
	cmd := exec.Command("powershell", params)
	out, err := cmd.Output()
	if err!=nil{
		return "", err
	}
	return string(out), nil
}

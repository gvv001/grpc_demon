//go:build linux

package lib

import (
	"fmt"
	"os/exec"
)

func ExecShellCommand(file, params string) (string, error) {
	cmd := exec.Command(file, params)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("file=%v, params %v: %v", file, params, err.Error())
	}
	return string(out), nil
}

package tooling

import (
	"os/exec"
	"strings"
)

func RunCommand(cmd string) string {
	_cmd := exec.Command("/bin/sh", "-c", cmd)
	output, err := _cmd.Output()

	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

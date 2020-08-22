package utils

import (
	"fmt"
	"os/exec"
	"runtime"
)

var cmds = map[string]string{
	"windows": "explorer.exe",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func Open(uri string) error {
	run, ok := cmds[runtime.GOOS]
	if !ok {
		return fmt.Errorf("platform not support: %v", runtime.GOOS)
	}
	cmd := exec.Command(run, uri)
	return cmd.Start()
}

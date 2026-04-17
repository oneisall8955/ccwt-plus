//go:build windows

package service

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/UserExistsError/conpty"
)

func quoteWindowsArg(arg string) string {
	if arg == "" {
		return `""`
	}
	if strings.ContainsAny(arg, " \t\"") {
		escaped := strings.ReplaceAll(arg, `"`, `\"`)
		return `"` + escaped + `"`
	}
	return arg
}

func buildWindowsCommandLine(shell string) string {
	name := strings.ToLower(filepath.Base(shell))
	switch name {
	case "pwsh.exe", "pwsh":
		return quoteWindowsArg(shell) + " -NoLogo -NoProfile"
	case "powershell.exe", "powershell":
		return quoteWindowsArg(shell) + " -NoLogo -NoProfile"
	default:
		return quoteWindowsArg(shell)
	}
}

func startWindowsPTY(shell, workdir string, env []string, rows, cols uint16) (io.ReadWriteCloser, func(rows, cols uint16) error, func(), error) {
	if !conpty.IsConPtyAvailable() {
		return nil, nil, nil, fmt.Errorf("windows ConPTY unavailable")
	}

	commandLine := buildWindowsCommandLine(shell)
	cpty, err := conpty.Start(
		commandLine,
		conpty.ConPtyDimensions(int(cols), int(rows)),
		conpty.ConPtyWorkDir(workdir),
		conpty.ConPtyEnv(env),
	)
	if err != nil {
		return nil, nil, nil, err
	}

	resize := func(rows, cols uint16) error {
		return cpty.Resize(int(cols), int(rows))
	}
	cleanup := func() {
		_ = cpty.Close()
	}
	return cpty, resize, cleanup, nil
}

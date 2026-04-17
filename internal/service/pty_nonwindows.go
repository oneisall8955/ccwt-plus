//go:build !windows

package service

import (
	"fmt"
	"io"
)

func startWindowsPTY(shell, workdir string, env []string, rows, cols uint16) (io.ReadWriteCloser, func(rows, cols uint16) error, func(), error) {
	return nil, nil, nil, fmt.Errorf("windows ConPTY unsupported on this platform")
}

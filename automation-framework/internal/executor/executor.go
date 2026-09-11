package executor

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"time"

	"github.com/icey148/qa-automation/automation-framework/internal/result"
)

type Executor struct{}

func New() *Executor { return &Executor{} }

func (e *Executor) Run(ctx context.Context, command string) result.ExecutionResult {
	started := time.Now()
	var stdout, stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	exitCode := 0
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			exitCode = ee.ExitCode()
		} else {
			exitCode = 1
		}
	}

	return result.ExecutionResult{
		Command: command, ExitCode: exitCode, StartedAt: started,
		Duration: time.Since(started), Stdout: strings.TrimSpace(stdout.String()),
		Stderr: strings.TrimSpace(stderr.String()), Successful: exitCode == 0,
	}
}

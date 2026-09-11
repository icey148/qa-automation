package result

import "time"

type ExecutionResult struct {
	Command    string        `json:"command"`
	ExitCode   int           `json:"exitCode"`
	StartedAt  time.Time     `json:"startedAt"`
	Duration   time.Duration `json:"duration"`
	Stdout     string        `json:"stdout"`
	Stderr     string        `json:"stderr"`
	Successful bool          `json:"successful"`
}

package report

import (
	"fmt"
	"io"

	"github.com/icey148/qa-automation/automation-framework/internal/result"
)

func Markdown(w io.Writer, r result.ExecutionResult) {
	status := "PASS"
	if !r.Successful {
		status = "FAIL"
	}
	fmt.Fprintln(w, "# Test Execution Report")
	fmt.Fprintln(w)
	fmt.Fprintf(w, "- Status: **%s**\n", status)
	fmt.Fprintf(w, "- Command: `%s`\n", r.Command)
	fmt.Fprintf(w, "- Exit Code: `%d`\n", r.ExitCode)
	fmt.Fprintf(w, "- Duration: `%s`\n", r.Duration)
	fmt.Fprintln(w, "\n## Stdout")
	fmt.Fprintln(w, "```text")
	fmt.Fprintln(w, r.Stdout)
	fmt.Fprintln(w, "```")
	if r.Stderr != "" {
		fmt.Fprintln(w, "\n## Stderr")
		fmt.Fprintln(w, "```text")
		fmt.Fprintln(w, r.Stderr)
		fmt.Fprintln(w, "```")
	}
}

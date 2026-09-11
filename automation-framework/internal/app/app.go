package app

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/icey148/qa-automation/automation-framework/internal/auth"
	"github.com/icey148/qa-automation/automation-framework/internal/config"
	"github.com/icey148/qa-automation/automation-framework/internal/executor"
	"github.com/icey148/qa-automation/automation-framework/internal/report"
	"github.com/icey148/qa-automation/automation-framework/internal/result"
)

const version = "0.1.0"

func Run(args []string, stdout, stderr io.Writer) error {
	if len(args) == 0 {
		usage(stdout)
		return nil
	}
	switch args[0] {
	case "run":
		return runCmd(args[1:], stdout, stderr)
	case "auth":
		return authCmd(args[1:], stdout, stderr)
	case "report":
		return reportCmd(args[1:], stdout, stderr)
	case "version":
		fmt.Fprintln(stdout, version)
		return nil
	case "help", "-h", "--help":
		usage(stdout)
		return nil
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func usage(w io.Writer) {
	fmt.Fprintln(w, `qa - QA automation framework CLI

Usage:
  qa run     [--config path] [--command command] [--require-auth]
  qa auth    [--env dev] [--role user] [--provider sso] [--ttl 8h]
  qa report  --input result.json [--output report.md]
  qa version

The framework executes existing tests. It does not generate Test Cases or Test Scripts.`)
}

func runCmd(args []string, stdout, stderr io.Writer) error {
	fs := flag.NewFlagSet("run", flag.ContinueOnError)
	fs.SetOutput(stderr)
	configPath := fs.String("config", "../configs/dev.json", "config path")
	command := fs.String("command", "", "override test command")
	requireAuth := fs.Bool("require-auth", false, "require a valid auth session")
	resultPath := fs.String("result", "", "optional JSON result output path")
	timeout := fs.Duration("timeout", 15*time.Minute, "execution timeout")
	if err := fs.Parse(args); err != nil {
		return err
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		return err
	}
	if *command != "" {
		cfg.TestCommand = *command
	}
	if cfg.TestCommand == "" {
		return errors.New("test command is empty")
	}

	if *requireAuth {
		mgr := auth.NewManager(cfg.Auth.SessionDir)
		if _, err := mgr.Validate(cfg.Environment, cfg.Auth.SessionRole); err != nil {
			return err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	r := executor.New().Run(ctx, cfg.TestCommand)
	if r.Stdout != "" {
		fmt.Fprintln(stdout, r.Stdout)
	}
	if r.Stderr != "" {
		fmt.Fprintln(stderr, r.Stderr)
	}
	if *resultPath != "" {
		if err := writeJSON(*resultPath, r); err != nil {
			return err
		}
	}
	if !r.Successful {
		return fmt.Errorf("test command failed with exit code %d", r.ExitCode)
	}
	return nil
}

func authCmd(args []string, stdout, stderr io.Writer) error {
	fs := flag.NewFlagSet("auth", flag.ContinueOnError)
	fs.SetOutput(stderr)
	env := fs.String("env", "dev", "environment")
	role := fs.String("role", "user", "identity role")
	provider := fs.String("provider", "sso", "auth provider")
	root := fs.String("session-dir", ".auth", "session metadata directory")
	ttl := fs.Duration("ttl", 8*time.Hour, "session metadata TTL")
	if err := fs.Parse(args); err != nil {
		return err
	}

	s, err := auth.NewManager(*root).Bootstrap(*env, *role, *provider, *ttl)
	if err != nil {
		return err
	}
	fmt.Fprintf(stdout, "auth metadata initialized for %s/%s (%s), expires %s\n", *env, *role, s.Provider, s.ExpiresAt.Format(time.RFC3339))
	fmt.Fprintln(stdout, "note: real enterprise SSO/MFA login must be implemented by a provider adapter; secrets are not stored by this scaffold")
	return nil
}

func reportCmd(args []string, stdout, stderr io.Writer) error {
	fs := flag.NewFlagSet("report", flag.ContinueOnError)
	fs.SetOutput(stderr)
	input := fs.String("input", "", "ExecutionResult JSON path")
	output := fs.String("output", "", "optional markdown output path; stdout when omitted")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *input == "" {
		return errors.New("--input is required")
	}
	b, err := os.ReadFile(*input)
	if err != nil {
		return err
	}
	var r result.ExecutionResult
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}
	if *output == "" {
		report.Markdown(stdout, r)
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(*output), 0o755); err != nil {
		return err
	}
	f, err := os.Create(*output)
	if err != nil {
		return err
	}
	defer f.Close()
	report.Markdown(f, r)
	fmt.Fprintln(stdout, *output)
	return nil
}

func writeJSON(path string, value any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil && filepath.Dir(path) != "." {
		return err
	}
	b, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

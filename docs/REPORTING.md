# Reporting

Reports must be grounded in real execution output.

```text
Test Script
  -> Test Execution
  -> Raw Result / Logs / JUnit / Trace
  -> Deterministic Report
  -> Optional AI analysis/summary
```

AI may summarize or cluster failures, but pass/fail counts, durations, Test Case IDs and error logs must come from execution data.

## Runtime location

During GitHub Actions execution, reporting output is collected under:

```text
test-artifacts/
├── reports/
├── logs/
├── debug_info/
├── telemetry/
└── cases/
```

The whole directory is uploaded as a GitHub Actions Artifact after the test run. It is not committed to Git.

For UI runs, Playwright JUnit/HTML output plus traces/screenshots/videos are collected. For API runs, raw `go test -json` output is preserved.

## CLI

The Go CLI still supports deterministic report generation from a saved execution result:

```bash
qa run --result test-artifacts/reports/result.json
qa report --input test-artifacts/reports/result.json
qa report --input test-artifacts/reports/result.json --output test-artifacts/reports/report.md
```

Long-lived QA documents belong under `artifacts/`; runtime reports belong under `test-artifacts/` and are uploaded by CI.

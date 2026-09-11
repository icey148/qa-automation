# Reporting

Reports must be grounded in real execution output.

```text
Test Script
  -> Test Execution
  -> Raw Result (exit code/stdout/stderr/JUnit/JSON)
  -> Deterministic Report
  -> Optional AI analysis/summary
```

AI may summarize or cluster failures, but pass/fail counts, durations, Test Case IDs, and error logs must come from execution data.

The CLI supports:

```bash
qa run --result artifacts/result.json
qa report --input artifacts/result.json
qa report --input artifacts/result.json --output artifacts/report.md
```

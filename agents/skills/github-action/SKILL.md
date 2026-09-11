---
name: github-action
description: Generate GitHub Actions workflows that execute existing automated test scripts and publish deterministic test artifacts.
---

# GitHub Action

## Inputs
Required: existing Test Script(s) and repository/runtime context.

## Workflow
1. Detect test language and commands.
2. Resolve runtime versions and working directories.
3. Configure a safe authentication strategy for CI.
4. Prepare a temporary `test-artifacts/` workspace.
5. Run tests while preserving the real test outcome.
6. Collect logs, reports, debug information, telemetry and case-level evidence.
7. Upload `test-artifacts/` with `actions/upload-artifact` using `if: always()`.
8. Name uploaded artifacts with `${{ github.run_id }}` and `${{ github.run_attempt }}` so reruns remain distinct.
9. Restore the workflow/job failure status after upload when tests failed.
10. Never commit runtime test artifacts to the repository branch.
11. Never embed credentials, cookies, tokens or SSO session files.

## Runtime Artifact Layout

```text
test-artifacts/
├── metadata.json
├── debug_info/
├── logs/
├── reports/
├── telemetry/
├── test_generation_output/
└── cases/
```

The GitHub Actions run is the historical execution boundary. Runtime artifacts are uploaded to GitHub Actions Artifacts, not stored under `tests/` and not committed to `main`.

## Failure Evidence Rule
A failed test run must still upload available logs, screenshots, traces, videos, JUnit/JSON and debug information. Use `continue-on-error` only as a temporary collection mechanism; a later step must fail the job when the test step failed.

## SSO Rule
Interactive MFA/PingID is not appropriate for unattended CI. Prefer service identity, OIDC, test-only machine identity or another approved non-interactive mechanism.

# Codex Handoff

This repository is ready for continued implementation with Codex.

Codex should read the root `AGENTS.md` first. It defines repository-wide architecture, validation and safety rules.

## Current implementation state

Implemented:

- QA Skills for Strategy, Plan, Test Case, Test Script and GitHub Actions.
- Test Script hard dependency on Test Case.
- Reuse-first UI test generation rules.
- Go CLI automation framework skeleton.
- SSO/session abstraction skeleton.
- Playwright UI structure with `cases/`, `pages/` and `framework/`.
- GitHub Actions `api` and `ui` targets.
- Unified `test-artifacts/` runtime workspace.
- GitHub Actions upload using `run_id` + `run_attempt` and `if: always()`.
- Documentation for runtime artifacts and CI lifecycle.

## Recommended next work

1. Replace the simple Go CLI argument handling with Cobra while preserving command behavior.
2. Add a Go UI executor that invokes Playwright and normalizes its output into the common execution-result model.
3. Add deterministic parsers for Go JSON and Playwright JUnit.
4. Add per-Test-Case artifact mapping into `test-artifacts/cases/<case-id>/` where practical.
5. Add real provider adapters for approved SSO/non-interactive CI authentication.
6. Add more reusable UI components such as table, select, date picker and pagination.
7. Add generation validation metadata under `test-artifacts/test_generation_output/`.
8. Add report aggregation across API/UI results without relying on AI for pass/fail truth.

## Important invariants

- Do not generate Test Script without a Test Case.
- Do not put runtime results beside source test files.
- Do not commit `test-artifacts/` or `.auth/`.
- Do not duplicate existing reusable UI abstractions.
- Do not bypass MFA/PingID.
- Reports must derive objective status/count/duration data from execution output.

## GitHub Actions artifact contract

Each QA job builds this temporary directory:

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

The workflow uploads the whole directory as:

```text
qa-<target>-<github.run_id>-<github.run_attempt>
```

The upload happens even when tests fail, then the job is explicitly failed after artifact preservation.

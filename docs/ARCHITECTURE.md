# Architecture

## Responsibility boundaries

```text
agents/               AI behavior: what to generate and generation rules
automation-framework/ Go CLI runtime: execute tests, auth/session abstraction, results, reporting
tests/                Executable test code
artifacts/            Long-lived QA documents such as Strategy, Plan and Test Cases
test-artifacts/       Ephemeral runtime output; uploaded by GitHub Actions, never committed
configs/              Environment/runtime configuration
.auth/                Local authentication/session state; never commit secrets
.github/               CI execution entry points
```

## QA artifact dependency

```text
Requirement / Jira / Spec
   ├──> Test Strategy
   ├──> Test Plan
   └──> Test Case ──> Test Script ──> GitHub Actions / Execution
                                      │
                                      ▼
                               Runtime Artifacts
                                      │
                                      ▼
                                  Test Report
```

Only `Test Script -> Test Case` is a hard QA-artifact dependency.

If a user requests only Test Script, the orchestration layer must first create and validate Test Case, then generate Test Script.

## Generation vs execution

The Agent layer generates QA artifacts. The Go automation framework executes existing test code. The framework must not silently generate or change Test Cases/Test Scripts.

## UI test layering

```text
Test Case
  -> tests/ui/cases/
  -> tests/ui/pages/
  -> tests/ui/framework/
  -> Playwright
```

Generated UI tests must reuse existing fixtures, Page Objects and reusable components before creating new abstractions.

## Runtime artifact lifecycle

A GitHub Actions job creates one temporary `test-artifacts/` workspace. The workflow run and run attempt provide historical isolation.

```text
GitHub Actions Run
  -> test-artifacts/
     -> metadata.json
     -> debug_info/
     -> logs/
     -> reports/
     -> telemetry/
     -> test_generation_output/
     -> cases/
  -> actions/upload-artifact
```

Runtime output is append-only at the GitHub Actions history level and is not stored on the repository branch.

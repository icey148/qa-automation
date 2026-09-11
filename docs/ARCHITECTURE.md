# Architecture

## Responsibility boundaries

```text
agents/               AI behavior: what to generate and generation rules
automation-framework/ Go CLI runtime: execute tests, auth/session abstraction, results, reporting
tests/                Executable test code
artifacts/            Generated QA artifacts and run outputs
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
                               Execution Result
                                      │
                                      ▼
                                  Test Report
```

Only `Test Script -> Test Case` is a hard QA-artifact dependency.

If a user requests only Test Script, the orchestration layer must first create and validate Test Case, then generate Test Script.

## Generation vs execution

The Agent layer generates QA artifacts. The Go automation framework executes existing test code. The framework must not silently generate or change Test Cases/Test Scripts.

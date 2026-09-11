# AGENTS.md

This repository contains an AI-assisted QA artifact system plus an executable automation testing framework.

## Non-negotiable architecture rules

1. `agents/` owns generation behavior and Skills.
2. `automation-framework/` is the Go execution framework. It executes tests; it must not silently invent or modify Test Cases/Test Scripts.
3. `tests/` contains executable SUT tests.
4. `artifacts/` contains long-lived QA documents such as Strategy, Plan and Test Cases.
5. `test-artifacts/` is ephemeral runtime output. It must not be committed. GitHub Actions uploads it with `actions/upload-artifact`.
6. `.auth/` contains local/session state and must never contain committed credentials, cookies or tokens.

## Test Script hard dependency

A Test Script MUST have a corresponding generated and validated Test Case.

Required flow:

```text
Requirement / Evidence -> Test Case -> Test Script
```

If a requested Script exposes a missing scenario, update/add the Test Case first, validate it, then update/generate the Script.

## UI automation rules

Default browser automation stack: TypeScript + Playwright.

Use this layering:

```text
Test Case -> tests/ui/cases/ -> tests/ui/pages/ -> tests/ui/framework/ -> Playwright
```

Before creating helpers, Page Objects, fixtures or component wrappers, search existing reusable code first.

- `cases/`: Test Case-specific business flow and assertions.
- `pages/`: page-level business interactions and stable locators.
- `framework/`: cross-page fixtures, auth/session adapters, reusable components, helpers, assertions and test data.

Prefer Playwright role/label/test-id locators over brittle CSS/XPath selectors.
Do not duplicate SSO/login flows inside individual test cases.

## API automation rules

Default backend/API automation stack: Go.
Generated API tests should reuse the shared test framework/client/assertion/fixture abstractions rather than reimplementing HTTP/auth logic per case.

## Runtime artifact rules

GitHub Actions runtime output must be collected under:

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

Uploads must use `if: always()` so failed tests preserve evidence. Artifact names should include GitHub `run_id` and `run_attempt`.
If test execution uses `continue-on-error` to allow evidence collection, explicitly fail the job after the upload step when tests failed.

## Authentication rules

Do not bypass MFA/PingID. Local interactive authentication and CI authentication are separate concerns.
For CI prefer an approved non-interactive identity mechanism such as OIDC/service/test identity.
Never commit `.auth/` state or secrets.

## Validation

After relevant changes, run the applicable checks:

```bash
make test
make build
```

For UI changes:

```bash
cd tests/ui
npm install
npx playwright test --list
```

When an actual SUT/base URL is available, run the affected Playwright spec(s).

## Documentation

When changing architecture, generated artifact locations, CI behavior, authentication, reporting or Skill contracts, update the corresponding files under `docs/` and README in the same change.

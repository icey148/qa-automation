---
name: test-script
description: Generate executable automation scripts only after a corresponding Test Case has been generated and validated.
---

# Test Script

## Hard Dependency
A corresponding Test Case MUST exist before any Test Script is generated.

If the user requests Test Script and no Test Case exists:
1. Generate the Test Case first.
2. Validate the Test Case.
3. Use that Test Case as the authoritative source.
4. Generate the Test Script.

Required flow:

`Requirement / Evidence -> Test Case -> Test Script`

## Inputs
Required: Test Case.
Supporting: source code, OpenAPI, frontend source, automation framework, fixtures, environment configuration, repository conventions.

## Core Rule
Test Case defines WHAT to test. Supporting implementation evidence defines HOW to automate it.

The script MUST NOT add, remove, or change tested behavior, expected results, or material assertions from the Test Case.

## Execution-Type Routing
- API/backend: generate Go tests using the automation framework where practical.
- Web UI: generate Playwright TypeScript tests.

## Reuse-First Rule
Before generating any new helper, selector wrapper, Page Object, fixture, component abstraction, or test utility:
1. Search existing `tests/ui/framework/` reusable capabilities.
2. Search existing `tests/ui/pages/` Page Objects.
3. Search existing `tests/ui/cases/` for established conventions and reusable patterns.
4. Reuse an existing abstraction whenever it satisfies the Test Case.
5. Create a new Page Object or reusable framework component only when no suitable abstraction exists.
6. Never duplicate SSO/login logic inside individual UI test cases.

Recommended UI layering:

`Test Case -> cases/ -> pages/ -> framework/ -> Playwright`

Responsibilities:
- `cases/`: business scenario and assertions specific to the Test Case.
- `pages/`: page-level business interactions and stable locators.
- `framework/`: cross-page fixtures, authentication, components, helpers, assertions, and test data utilities.

## Frontend Evidence Resolution
For UI tests, inspect available frontend evidence before implementation:
- routes
- pages/components
- labels and accessible roles
- existing `data-testid` attributes
- existing Playwright tests
- existing Page Objects
- authentication fixtures
- project conventions

Prefer resilient user-facing Playwright locators such as role and label locators. Avoid brittle CSS/XPath selectors unless the application provides no stable semantic alternative.

## Validation
Before considering a generated UI script complete:
1. Ensure the canonical Test Case ID is preserved.
2. Ensure imports resolve against the existing UI framework structure.
3. Ensure the script does not duplicate reusable logic unnecessarily.
4. Ensure the script is discoverable by Playwright.
5. Execute or at minimum list/compile the test when the runtime is available.

## Output
- API/backend: Go `*_test.go` scripts.
- Web UI: Playwright TypeScript `*.spec.ts` scripts under `tests/ui/cases/`.
- Include a stable reference to the Test Case ID in every generated script.

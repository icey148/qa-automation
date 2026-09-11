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

## Output
- API/backend: prefer Go test scripts using the automation framework.
- Web UI: Playwright TypeScript may be used when browser coverage is required.
- Include a stable reference to the Test Case ID in the script.
